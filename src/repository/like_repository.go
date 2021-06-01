package repository

import (
	"github.com/gocql/gocql"
	"post-service/domain"
	)

const (
	CreateLikeTable = "CREATE TABLE if not exists post_service.Likes (post_id, timestamp, profile_id, " +
		"PRIMARY KEY (post_id, profile_id));"
	CreateDislikeTable = "CREATE TABLE if not exists post_service.Dislikes (post_id, timestamp, profile_id, " +
		"PRIMARY KEY (post_id, profile_id));"
)

type LikeRepo interface {
	LikePost(postId uint, profile *domain.Profile) error
	RemoveLike(postId uint, profile *domain.Profile) error
	DislikePost(postId uint, profile *domain.Profile) error
	RemoveDislike(postId uint, profile *domain.Profile) error
	GetLikesForPost(postId uint) ([]domain.Like, error)
	GetDislikesForPost(postId uint) ([]domain.Dislike, error)
	GetNumOfLikesForPost(postId uint) (uint64, error)
	GetNumOfDislikesForPost(postId uint) (uint64, error)
}

type likeRepository struct {
	cassandraSession *gocql.Session
}

func (l likeRepository) GetLikesForPost(postId uint) ([]domain.Like, error) {
	panic("implement me")
}

func (l likeRepository) GetDislikesForPost(postId uint) ([]domain.Dislike, error) {
	panic("implement me")
}

func (l likeRepository) GetNumOfLikesForPost(postId uint) (uint64, error) {
	panic("implement me")
}

func (l likeRepository) GetNumOfDislikesForPost(postId uint) (uint64, error) {
	panic("implement me")
}

func (l likeRepository) LikePost(postId uint, profile *domain.Profile) error {
	panic("implement me")
}

func (l likeRepository) DislikePost(postId uint, profile *domain.Profile) error {
	panic("implement me")
}

func (l likeRepository) RemoveLike(postId uint, profile *domain.Profile) error {
	panic("implement me")
}

func (l likeRepository) RemoveDislike(postId uint, profile *domain.Profile) error {
	panic("implement me")
}

func NewLikeRepository(cassandraSession *gocql.Session) LikeRepo {
	var l =  &likeRepository{
		cassandraSession : cassandraSession,
	}
	l.cassandraSession.Query(CreateLikeTable).Exec()
	l.cassandraSession.Query(CreateDislikeTable).Exec()
	return l
}
