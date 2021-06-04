package repository

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
)

const (
	CreateFavoritesTable = "CREATE TABLE if not exists post_keyspace.Favorites (profile_id text, time_of_creation timestamp, posts map<text, text>, " +
		"PRIMARY KEY (profile_id));"
	InsertFavoriteStatement = "INSERT INTO post_keyspace.Favorites (profile_id, time_of_creation, posts) VALUES (?, ?, ?) IF NOT EXISTS;"
	GetFavoritesForUser     = "SELECT posts FROM post_keyspace.Favorites WHERE profile_id = ?;"
	UpdateFavorites         = "UPDATE post_keyspace.Favorites SET posts = ? WHERE profile_id = ?;"
)

type FavoritesRepo interface {
	AddPostToFavorites(postId string, profileId string, postBy string, ctx context.Context) error
	RemovePostFromFavorites(postId string, profileId string, ctx context.Context) error
	GetFavorites(userId string) (map[string]string, error)
}

type favoritesRepository struct {
	cassandraSession *gocql.Session
}

func (f favoritesRepository) GetFavorites(userId string) (map[string]string, error) {
	var posts map[string]string
	iter := f.cassandraSession.Query(GetFavoritesForUser, userId).Iter()
	if iter == nil {
		return nil, fmt.Errorf("no such element")
	}

	iter.Scan(&posts)

	return  posts, nil
}

func (f favoritesRepository) AddPostToFavorites(postId string, profileId string, postBy string, ctx context.Context) error {
	iter := f.cassandraSession.Query(GetFavoritesForUser, profileId).Iter()

	if iter == nil {
		return fmt.Errorf("no such element")
	}
	var posts map[string]string
	for iter.Scan(&posts) {
		posts[postId] = postBy
	}

	err := f.cassandraSession.Query(UpdateFavorites, posts, profileId).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (f favoritesRepository) RemovePostFromFavorites(postId string, profileId string, ctx context.Context) error {
	iter := f.cassandraSession.Query(GetFavoritesForUser, profileId).Iter()

	if iter == nil {
		return fmt.Errorf("no such element")
	}
	var posts map[string]string
	for iter.Scan(&posts) {
		delete(posts, postId)
	}

	err := f.cassandraSession.Query(UpdateFavorites, posts, profileId).Exec()

	if err != nil {
		return err
	}

	return nil
}

func NewFavoritesRepository(cassandraSession *gocql.Session) FavoritesRepo {
	var f = favoritesRepository {
		cassandraSession: cassandraSession,
	}
	err := f.cassandraSession.Query(CreateFavoritesTable).Exec()
	if err != nil {
		return nil
	}
	return f
}

