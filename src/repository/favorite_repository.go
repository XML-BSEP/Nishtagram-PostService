package repository

import (
	"context"
	"github.com/gocql/gocql"
)

const (
	CreateFavoritesTable = "CREATE TABLE if not exists post_keyspace.Favorites (profile_id text, time_of_creation timestamp, posts list<text>, " +
		"PRIMARY KEY (profile_id));"
)

type FavoritesRepo interface {
	AddPostToFavorites(postId string, profileId string, ctx context.Context) error
	RemovePostFromFavorites(postId string, profileId string, ctx context.Context) error
}

type favoritesRepository struct {
	cassandraSession *gocql.Session
}

func (f favoritesRepository) AddPostToFavorites(postId string, profileId string, ctx context.Context) error {
	panic("implement me")
}

func (f favoritesRepository) RemovePostFromFavorites(postId string, profileId string, ctx context.Context) error {
	panic("implement me")
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

