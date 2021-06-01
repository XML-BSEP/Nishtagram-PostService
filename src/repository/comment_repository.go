package repository

import (
	"github.com/gocql/gocql"
	"post-service/domain"
)
const (
	CreateCommentTable = "CREATE TABLE if not exists post_service.Comments (comment, post_id, comment_by,timestamp, mentions list<int>, " +
		"PRIMARY KEY (post_id, comment_by));"
)
type CommentRepo interface {
	CommentPost(comment *domain.Comment) error
	DeleteComment(comment domain.Comment) error
}

type commentRepository struct {
	cassandraSession *gocql.Session
}

func (c commentRepository) CommentPost(comment *domain.Comment) error {
	panic("implement me")
}

func (c commentRepository) DeleteComment(comment domain.Comment) error {
	panic("implement me")
}

func NewCommentRepository(cassandraSession *gocql.Session) CommentRepo {
	var c = &commentRepository{
		cassandraSession: cassandraSession,
	}
	c.cassandraSession.Query(CreateCommentTable).Exec()
	return c
}
