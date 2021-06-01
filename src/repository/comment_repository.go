package repository

import (
	"context"
	"github.com/gocql/gocql"
	"post-service/domain"
)
const (
	CreateCommentTable = "CREATE TABLE if not exists post_keyspace.Comments (id text, comment text, post_id text, comment_by text, timestamp timestamp, mentions list<text>, " +
		"PRIMARY KEY (post_id, comment_by));"
)
type CommentRepo interface {
	CommentPost(comment *domain.Comment, ctx context.Context) error
	DeleteComment(comment domain.Comment, ctx context.Context) error
}

type commentRepository struct {
	cassandraSession *gocql.Session
}

func (c commentRepository) CommentPost(comment *domain.Comment, ctx context.Context) error {
	panic("implement me")
}

func (c commentRepository) DeleteComment(comment domain.Comment, ctx context.Context) error {
	panic("implement me")
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
