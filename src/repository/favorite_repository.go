package repository

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	logger "github.com/jelena-vlajkov/logger/logger"
	"time"
)

const (
	CreateFavoritesTable = "CREATE TABLE if not exists post_keyspace.Favorites (profile_id text, time_of_creation timestamp, posts map<text, text>, " +
		"PRIMARY KEY (profile_id));"
	InsertFavoriteStatement = "INSERT INTO post_keyspace.Favorites (profile_id, time_of_creation, posts) VALUES (?, ?, ?) IF NOT EXISTS;"
	GetFavoritesForUser     = "SELECT posts FROM post_keyspace.Favorites WHERE profile_id = ?;"
	UpdateFavorites         = "UPDATE post_keyspace.Favorites SET posts = ? WHERE profile_id = ?;"
	SeeIfExists = "SELECT count(*) from post_keyspace.Favorites WHERE profile_id = ?;"
	GetFavorites = "SELECT posts FROM post_keyspace.Favorites WHERE profile_id = ?;"
)

type FavoritesRepo interface {
	AddPostToFavorites(postId string, profileId string, postBy string, ctx context.Context) error
	RemovePostFromFavorites(postId string, profileId string, ctx context.Context) error
	GetFavorites(userId string) (map[string]string, error)
}

type favoritesRepository struct {
	cassandraSession *gocql.Session
	logger *logger.Logger
}


func (f favoritesRepository) GetFavorites(userId string) (map[string]string, error) {
	var posts map[string]string
	iter := f.cassandraSession.Query(GetFavoritesForUser, userId).Iter()
	if iter == nil {
		f.logger.Logger.Errorf("no favorites for user %v\n", userId)
		return nil, fmt.Errorf("no such element")
	}

	iter.Scan(&posts)

	return  posts, nil
}

func (f favoritesRepository) AddPostToFavorites(postId string, profileId string, postBy string, ctx context.Context) error {
	iter := f.cassandraSession.Query(GetFavoritesForUser, profileId).Iter()

	if iter == nil {
		f.logger.Logger.Errorf("no favorites for user %v\n", profileId)
		return fmt.Errorf("no such element")
	}
	var posts map[string]string
	for iter.Scan(&posts) {
		posts[postId] = postBy
	}
	if len(posts) == 0 {
		posts = make(map[string]string, 1)
		posts[postId] = postBy
	}

	var ifExists int
	var err error
	f.cassandraSession.Query(SeeIfExists, profileId).Iter().Scan(&ifExists)
	if ifExists == 0 {
		err = f.cassandraSession.Query(InsertFavoriteStatement, profileId, time.Now(), posts).Exec()
	} else {
		err = f.cassandraSession.Query(UpdateFavorites, posts, profileId).Exec()
	}
	if err != nil {
		f.logger.Logger.Errorf("error while adding post %v to favorites for user %v, error: %v\n", postId, profileId, err)
		return err
	}

	return nil
}

func (f favoritesRepository) RemovePostFromFavorites(postId string, profileId string, ctx context.Context) error {
	var posts map[string]string
	f.cassandraSession.Query(GetFavoritesForUser, profileId).Iter().Scan(&posts)
	delete(posts, postId)
	err := f.cassandraSession.Query(UpdateFavorites, posts, profileId).Exec()

	if err != nil {
		f.logger.Logger.Errorf("error while removin post %v from favorites for user %v, error: %v\n", postId, profileId, err)
		return err
	}

	return nil
}

func NewFavoritesRepository(cassandraSession *gocql.Session, logger *logger.Logger) FavoritesRepo {
	var f = favoritesRepository {
		cassandraSession: cassandraSession,
		logger: logger,
	}
	err := f.cassandraSession.Query(CreateFavoritesTable).Exec()
	if err != nil {
		return nil
	}
	return f
}

