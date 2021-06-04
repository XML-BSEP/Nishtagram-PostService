package repository

import (
	"context"
	"github.com/gocql/gocql"
	"post-service/dto"
	"time"
)
const (
	CreateCommentTable = "CREATE TABLE if not exists post_keyspace.Comments (comment text, post_id text, comment_by text, timestamp timestamp, mentions list<text>, " +
		"PRIMARY KEY (post_id, comment_by));"
	InsertComment = "INSERT INTO post_keyspace.Comments (comment, post_id, comment_by, timestamp, mentions) VALUES (? , ?, ?, ?, ?);"
	DeleteComment = "DELETE FROM post_keyspace.Comments WHERE post_id = ? AND comment_by = ?;"
)
type CommentRepo interface {
	CommentPost(comment dto.CommentDTO, ctx context.Context) error
	DeleteComment(comment dto.CommentDTO, ctx context.Context) error
}

type commentRepository struct {
	cassandraSession *gocql.Session
}

func (c commentRepository) CommentPost(comment dto.CommentDTO, ctx context.Context) error {
	mentions := make([]string, len(comment.Mentions))

	for i, c := range comment.Mentions {
		mentions[i] = c.Id
	}

	err := c.cassandraSession.Query(InsertComment, comment.Comment, comment.PostId, comment.CommentBy.Id, time.Now(), mentions).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (c commentRepository) DeleteComment(comment dto.CommentDTO, ctx context.Context) error {
	err := c.cassandraSession.Query(DeletePost, comment.PostId, comment.CommentBy).Exec()
	if err != nil {
		return err
	}

	return nil
}

func NewCommentRepository(cassandraSession *gocql.Session) CommentRepo {
	var c = &commentRepository{
		cassandraSession: cassandraSession,
	}
	err := c.cassandraSession.Query(CreateCommentTable).Exec()
	if err != nil {
		return nil
	}
	return c
}
