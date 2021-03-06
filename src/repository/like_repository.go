package repository

import (
	"context"
	"github.com/gocql/gocql"
	logger "github.com/jelena-vlajkov/logger/logger"
	"post-service/domain"
	"time"
)

const (
	CreateLikeTable = "CREATE TABLE if not exists post_keyspace.Likes (post_id text, timestamp timestamp, profile_id text, " +
		"PRIMARY KEY (post_id, profile_id));"
	CreateShowLikesTable = "CREATE TABLE IF NOT EXISTS post_keyspace.LikesToShow (profile_id text, post_id text, post_by text, timestamp timestamp, " +
		"PRIMARY KEY (profile_id, post_id)) WITH CLUSTERING ORDER BY (post_id ASC);"
	CreateShowDislikesTable = "CREATE TABLE IF NOT EXISTS post_keyspace.DislikesToShow (profile_id text, post_id text, post_by text, timestamp timestamp, " +
		"PRIMARY KEY (profile_id, post_id)) WITH CLUSTERING ORDER BY (post_id ASC);"
	InsertIntoShowDislikesTable = "INSERT INTO post_keyspace.DislikesToShow (profile_id, post_id, post_by, timestamp) VALUES (?, ?, ?, ?) IF NOT EXISTS;"
	InsertIntoShowLikesTable = "INSERT INTO post_keyspace.LikesToShow (profile_id, post_id, post_by, timestamp) VALUES (?, ?, ?, ?) IF NOT EXISTS;"
	DeleteFromShowDislikesTable = "DELETE FROM post_keyspace.DislikesToShow WHERE profile_id = ? AND post_id = ?;"
	DeleteFromShowLikesTable = "DELETE FROM post_keyspace.LikesToShow WHERE profile_id = ? AND post_id = ?;"
	GetAllLikedMedia = "SELECT profile_id, post_id, post_by, timestamp FROM post_keyspace.LikesToShow WHERE profile_id = ?"
	GetAllDislikedMedia = "SELECT profile_id, post_id, post_by, timestamp FROM post_keyspace.DislikesToShow WHERE profile_id = ?"
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
	GetLikedMedia(profileId string, ctx context.Context) ([]domain.Like, error)
	GetDislikedMedia(profileId string, ctx context.Context) ([]domain.Dislike, error)
}

type likeRepository struct {
	cassandraSession *gocql.Session
	logger *logger.Logger
}

func (l likeRepository) GetLikedMedia(profileId string, ctx context.Context) ([]domain.Like, error) {
	var postId, likeBy, postBy string
	var timestamp time.Time
	var retVal []domain.Like
	iter := l.cassandraSession.Query(GetAllLikedMedia, profileId).Iter().Scanner()

	for iter.Next() {
		iter.Scan(&likeBy, &postId, &postBy, &timestamp)
		retVal = append(retVal, domain.Like{PostId: postId, Profile: domain.Profile{Id: likeBy}, PostBy: domain.Profile{Id: postBy}, Timestamp: timestamp})
	}

	return retVal, nil
}

func (l likeRepository) GetDislikedMedia(profileId string, ctx context.Context) ([]domain.Dislike, error) {
	var postId, likeBy, postBy string
	var timestamp time.Time
	var retVal []domain.Dislike
	iter := l.cassandraSession.Query(GetAllDislikedMedia, profileId).Iter().Scanner()

	for iter.Next() {
		iter.Scan(&likeBy, &postId, &postBy, &timestamp)
		retVal = append(retVal, domain.Dislike{PostId: postId, Profile: domain.Profile{Id: likeBy}, PostBy: domain.Profile{Id: postBy}, Timestamp: timestamp})
	}

	return retVal, nil
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

	err = l.cassandraSession.Query(InsertIntoShowLikesTable, profile.Id, postId, postBy, time.Now() ).Exec()
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
		l.logger.Logger.Errorf("error while ading like for user %v on post %v, error: %v\n", profile.Id, postId, err)
		return err
	}

	return nil
}

func (l likeRepository) DislikePost(postId string, postBy string, profile domain.Profile, ctx context.Context) error {
	err := l.cassandraSession.Query(InsertDislikeStatement, postId, time.Now(), profile.Id).Exec()
	if err != nil {
		return err
	}


	err = l.cassandraSession.Query(InsertIntoShowDislikesTable, profile.Id, postId, postBy, time.Now() ).Exec()
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
		l.logger.Logger.Errorf("error while removing dislike for user %v on post %v, error: %v\n", profile.Id, postId, err)
		return err
	}

	return nil
}

func (l likeRepository) RemoveLike(postId string, postBy string, profile domain.Profile, ctx context.Context) error {
	err := l.cassandraSession.Query(RemoveLike, postId, profile.Id).Exec()
	if err != nil {
		return err
	}


	err = l.cassandraSession.Query(DeleteFromShowLikesTable, profile.Id, postId).Exec()
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
		l.logger.Logger.Errorf("error while removing like for user %v on post %v, error %v\n", profile.Id, postId, err)
		return err
	}

	return nil
}

func (l likeRepository) RemoveDislike(postId string, postBy string, profile domain.Profile, ctx context.Context) error {
	err := l.cassandraSession.Query(RemoveDislike, postId, profile.Id).Exec()
	if err != nil {
		return err
	}



	err = l.cassandraSession.Query(DeleteFromShowDislikesTable, profile.Id, postId).Exec()
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
		l.logger.Logger.Errorf("error while removing dislike for user %v on post %v, error %v\n", profile.Id, postId, err)
		return err
	}

	return nil
}

func NewLikeRepository(cassandraSession *gocql.Session, logger *logger.Logger) LikeRepo {
	var l =  &likeRepository{
		cassandraSession : cassandraSession,
		logger: logger,
	}
	err := l.cassandraSession.Query(CreateLikeTable).Exec()
	if err != nil {
		return nil
	}
	err = l.cassandraSession.Query(CreateDislikeTable).Exec()
	if err != nil {
		return nil
	}

	err = l.cassandraSession.Query(CreateShowLikesTable).Exec()
	if err != nil {
		return nil
	}

	err = l.cassandraSession.Query(CreateShowDislikesTable).Exec()
	if err != nil {
		return nil
	}

	return l
}
