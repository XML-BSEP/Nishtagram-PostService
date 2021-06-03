package repository

import (
	"context"
	"github.com/gocql/gocql"
	"post-service/domain"
	"time"
)

const (
	CreateLikeTable = "CREATE TABLE if not exists post_keyspace.Likes (post_id text, timestamp timestamp, profile_id text, " +
		"PRIMARY KEY (post_id, profile_id));"
	CreateDislikeTable = "CREATE TABLE if not exists post_keyspace.Dislikes (post_id text, timestamp timestamp, profile_id text, " +
		"PRIMARY KEY (post_id, profile_id));"
	InsertLikeStatement = "INSERT INTO post_keyspace.Likes (post_id, timestamp, profile_id) VALUES (?, ?, ?) IF NOT EXISTS;"
	InsertDislikeStatement  = "INSERT INTO post_keyspace.Dislikes (post_id, timestamp, profile_id) VALUES (?, ?, ?) IF NOT EXISTS;"
	RemoveLike = "DELETE FROM post_keyspace.Likes WHERE post_id = ? AND profile_id = ?;"
	RemoveDislike = "DELETE FROM post_keyspace.Dislikes WHERE post_id = ? AND profile_id = ?;"
	GetAllLikesPerPost = "SELECT post_id, profiles_id, timestamp FROM post_keyspace.Likes WHERE post_id = ?;"
	GetAllDislikesPerPost = "SELECT post_id, profiles_id, timestamp FROM post_keyspace.Dislikes WHERE post_id = ?;"
	SeeIfLikeExists = "SELECT count(*) FROM post_keyspace.Likes WHERE post_id = ? AND profile_id = ?;"
	SeeIfDislikeExists = "SELECT count(*) FROM post_keyspace.Dislikes WHERE post_id = ? AND profile_id = ?;"
)

type LikeRepo interface {
	LikePost(postId string, postBy string, profile domain.Profile, ctx context.Context) error
	RemoveLike(postId string, postBy string, profile domain.Profile, ctx context.Context) error
	DislikePost(postId string, postBy string, profile domain.Profile, ctx context.Context) error
	RemoveDislike(postId string, postBy string, profile domain.Profile, ctx context.Context) error
	GetLikesForPost(postId string, ctx context.Context) ([]domain.Like, error)
	GetDislikesForPost(postId string, ctx context.Context) ([]domain.Dislike, error)
	SeeIfLikeExists(postId string, profileId string, ctx context.Context) bool
	SeeIfDislikeExists(postId string, profileId string, ctx context.Context) bool
}

type likeRepository struct {
	cassandraSession *gocql.Session
}

func (l likeRepository) SeeIfLikeExists(postId string, profileId string, ctx context.Context) bool {
	ifExists := 0
	l.cassandraSession.Query(SeeIfLikeExists, postId, profileId).Iter().Scan(&ifExists)
	return ifExists > 0
}

func (l likeRepository) SeeIfDislikeExists(postId string, profileId string, ctx context.Context) bool {
	ifExists := 0
	l.cassandraSession.Query(SeeIfDislikeExists, postId, profileId).Iter().Scan(&ifExists)
	return ifExists > 0
}

func (l likeRepository) GetLikesForPost(postId string, ctx context.Context) ([]domain.Like, error) {
	var profileId string
	var timestamp time.Time

	iter := l.cassandraSession.Query(GetAllLikesPerPost, postId).Iter().Scanner()
	var likes []domain.Like
	for iter.Next() {
		iter.Scan(&postId, &profileId, &timestamp)
		likes = append(likes, domain.NewLike(postId, profileId, timestamp))
	}

	return likes, nil
}

func (l likeRepository) GetDislikesForPost(postId string, ctx context.Context) ([]domain.Dislike, error) {
	var profileId string
	var timestamp time.Time

	iter := l.cassandraSession.Query(GetAllDislikesPerPost, postId).Iter().Scanner()
	var likes []domain.Dislike
	for iter.Next() {
		iter.Scan(&postId, &profileId, &timestamp)
		likes = append(likes, domain.NewDislike(postId, profileId, timestamp))
	}

	return likes, nil
}

func (l likeRepository) LikePost(postId string, postBy string, profile domain.Profile, ctx context.Context) error {

	err := l.cassandraSession.Query(InsertLikeStatement, postId, time.Now(), profile.Id).Exec()
	if err != nil {
		return err
	}

	var numOfLikes int
	iter := l.cassandraSession.Query(GetNumOfLikesForPost, postId, postBy).Iter()

	for iter.Scan(&numOfLikes) {
		numOfLikes = numOfLikes + 1
	}
	err = l.cassandraSession.Query(AddLikeToPost, numOfLikes, postId, postBy).Exec()


	if err != nil {
		return err
	}

	return nil
}

func (l likeRepository) DislikePost(postId string, postBy string, profile domain.Profile, ctx context.Context) error {
	err := l.cassandraSession.Query(InsertDislikeStatement, postId, time.Now(), profile.Id).Exec()
	if err != nil {
		return err
	}
	var numOfDislikes int
	iter := l.cassandraSession.Query(GetNumOfDislikesForPost, postId, postBy).Iter()

	for iter.Scan(&numOfDislikes) {
		numOfDislikes = numOfDislikes + 1
	}
	err = l.cassandraSession.Query(AddDislikeToPost, numOfDislikes, postId, postBy).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (l likeRepository) RemoveLike(postId string, postBy string, profile domain.Profile, ctx context.Context) error {
	err := l.cassandraSession.Query(RemoveLike, postBy, profile.Id).Exec()
	if err != nil {
		return err
	}
	var numOfLikes int
	iter := l.cassandraSession.Query(GetNumOfLikesForPost, postId, postBy).Iter()

	for iter.Scan(&numOfLikes) {
		numOfLikes = numOfLikes - 1
	}

	err = l.cassandraSession.Query(RemoveLikeFromPost, numOfLikes, postId, postBy).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (l likeRepository) RemoveDislike(postId string, postBy string, profile domain.Profile, ctx context.Context) error {
	err := l.cassandraSession.Query(RemoveDislike, postBy, profile.Id).Exec()
	if err != nil {
		return err
	}
	var numOfDislikes int
	iter := l.cassandraSession.Query(GetNumOfDislikesForPost, postId, postBy).Iter()

	for iter.Scan(&numOfDislikes) {
		numOfDislikes = numOfDislikes - 1
	}

	err = l.cassandraSession.Query(RemoveLikeFromPost, numOfDislikes, postId, postBy).Exec()

	if err != nil {
		return err
	}

	return nil
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
