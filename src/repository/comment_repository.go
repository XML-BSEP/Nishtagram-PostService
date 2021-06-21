package repository

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	logger "github.com/jelena-vlajkov/logger/logger"
	"post-service/domain"
	"post-service/dto"
	"post-service/gateway"
	"time"
)
const (
	CreateCommentTable = "CREATE TABLE if not exists post_keyspace.Comments (id text, comment text, post_id text, comment_by text, timestamp timestamp, mentions list<text>, " +
		"PRIMARY KEY (post_id, id));"
	InsertComment = "INSERT INTO post_keyspace.Comments (id, comment, post_id, comment_by, timestamp, mentions) VALUES (?, ? , ?, ?, ?, ?);"
	DeleteComment = "DELETE FROM post_keyspace.Comments WHERE post_id = ? AND comment_by = ?;"
	GetComments = "SELECT id, comment, post_id, comment_by, timestamp, mentions FROM post_keyspace.Comments WHERE post_id = ?;"
)
type CommentRepo interface {
	CommentPost(comment dto.CommentDTO, ctx context.Context) error
	DeleteComment(comment dto.CommentDTO, ctx context.Context) error
	GetComments(postId string, ctx context.Context) ([]dto.CommentDTO, error)
}

type commentRepository struct {
	cassandraSession *gocql.Session
	logger *logger.Logger
}

func (c commentRepository) GetComments(postId string, ctx context.Context) ([]dto.CommentDTO, error) {
	//id, comment, post_id, comment_by, timestamp, mentions
	var id, comment, post_id, comment_by string
	var timestamp time.Time
	var mentions []string
	var retVal []dto.CommentDTO
	iter := c.cassandraSession.Query(GetComments, postId).Iter().Scanner()

	for iter.Next() {
		iter.Scan(&id, &comment, &post_id, &comment_by, &timestamp, &mentions)
		dto := dto.CommentDTO{}
		profile, err := gateway.GetUser(context.Background(), comment_by)
		if err != nil {
			c.logger.Logger.Errorf("error while getting user info for %v, error: %v\n", comment_by, err)
		}
		dto.CommentBy = domain.Profile{Id: comment_by, ProfilePhoto: profile.ProfilePhoto, Username: profile.Username}
		dto.Comment = comment
		dto.PostId = postId
		retVal = append(retVal, dto)
	}

	return retVal, nil
}

func (c commentRepository) CommentPost(comment dto.CommentDTO, ctx context.Context) error {
	mentions := make([]string, 1)
	id := uuid.NewString()



	err := c.cassandraSession.Query(InsertComment, id, comment.Comment, comment.PostId, comment.CommentBy.Id, time.Now(), mentions).Exec()

	if err != nil {
		c.logger.Logger.Errorf("error while adding comment on post %v by user %v, error: %v\n", comment.PostId, comment.PostBy, err)
		return err
	}

	var numOfCom int

	c.cassandraSession.Query(GetNumOfCommentsForPost, comment.PostId, comment.PostBy).Iter().Scan(&numOfCom)
	numOfCom = numOfCom + 1
	c.cassandraSession.Query(AddCommentToPost, numOfCom, comment.PostId, comment.PostBy).Exec()
	return nil
}

func (c commentRepository) DeleteComment(comment dto.CommentDTO, ctx context.Context) error {
	err := c.cassandraSession.Query(DeletePost, comment.PostId, comment.CommentBy).Exec()
	if err != nil {
		c.logger.Logger.Errorf("error while deleting comment on post %v with id %v, error: %v\n", comment.PostId, comment.PostId, err)
		return err
	}

	return nil
}

func NewCommentRepository(cassandraSession *gocql.Session, logger *logger.Logger) CommentRepo {
	var c = &commentRepository{
		cassandraSession: cassandraSession,
		logger: logger,
	}
	err := c.cassandraSession.Query(CreateCommentTable).Exec()
	if err != nil {
		return nil
	}
	return c
}
