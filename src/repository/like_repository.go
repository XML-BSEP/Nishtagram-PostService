package repository

import (
	"context"
	"github.com/gocql/gocql"
	"post-service/domain"
	)

const (
	CreateLikeTable = "CREATE TABLE if not exists post_keyspace.Likes (post_id text, timestamp timestamp, profile_id text, " +
		"PRIMARY KEY (post_id, profile_id));"
	CreateDislikeTable = "CREATE TABLE if not exists post_keyspace.Dislikes (post_id text, timestamp timestamp, profile_id text, " +
		"PRIMARY KEY (post_id, profile_id));"
)

type LikeRepo interface {
	LikePost(postId string, profile *domain.Profile, ctx context.Context) error
	RemoveLike(postId string, profile *domain.Profile, ctx context.Context) error
	DislikePost(postId string, profile *domain.Profile, ctx context.Context) error
	RemoveDislike(postId string, profile *domain.Profile, ctx context.Context) error
	GetLikesForPost(postId string, ctx context.Context) ([]domain.Like, error)
	GetDislikesForPost(postId string, ctx context.Context) ([]domain.Dislike, error)
	GetNumOfLikesForPost(postId string, ctx context.Context) (uint64, error)
	GetNumOfDislikesForPost(postId string, ctx context.Context) (uint64, error)
}

type likeRepository struct {
	cassandraSession *gocql.Session
}

func (l likeRepository) GetLikesForPost(postId string, ctx context.Context) ([]domain.Like, error) {
	panic("implement me")
}

func (l likeRepository) GetDislikesForPost(postId string, ctx context.Context) ([]domain.Dislike, error) {
	panic("implement me")
}

func (l likeRepository) GetNumOfLikesForPost(postId string, ctx context.Context) (uint64, error) {
	panic("implement me")
}

func (l likeRepository) GetNumOfDislikesForPost(postId string, ctx context.Context) (uint64, error) {
	panic("implement me")
}

func (l likeRepository) LikePost(postId string, profile *domain.Profile, ctx context.Context) error {
	panic("implement me")
}

func (l likeRepository) DislikePost(postId string, profile *domain.Profile, ctx context.Context) error {
	panic("implement me")
}

func (l likeRepository) RemoveLike(postId string, profile *domain.Profile, ctx context.Context) error {
	panic("implement me")
}

func (l likeRepository) RemoveDislike(postId string, profile *domain.Profile, ctx context.Context) error {
	panic("implement me")
}

func NewLikeRepository(cassandraSession *gocql.Session) LikeRepo {
	var l =  &likeRepository{
		cassandraSession : cassandraSession,
	}
	err := l.cassandraSession.Query(CreateLikeTable).Exec()
	if err != nil {
		return nil
	}
	err = l.cassandraSession.Query(CreateDislikeTable).Exec()
	if err != nil {
		return nil
	}
	return l
}
