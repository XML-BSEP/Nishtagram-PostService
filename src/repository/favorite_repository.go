package repository

import "github.com/gocql/gocql"

const (
	CreateFavoritesTable = "CREATE TABLE if not exists post_service.Favorites (profile_id, time_of_creation, posts list<int>, " +
		"PRIMARY KEY (profile_id));"
)

type FavoritesRepo interface {
	AddPostToFavorites(postId uint, profileId uint) error
	RemovePostFromFavorites(postId uint, profileId uint) error
}

type favoritesRepository struct {
	cassandraSession *gocql.Session
}

func (f favoritesRepository) AddPostToFavorites(postId uint, profileId uint) error {
	panic("implement me")
}

func (f favoritesRepository) RemovePostFromFavorites(postId uint, profileId uint) error {
	panic("implement me")
}

func NewFavoritesRepository(cassandraSession *gocql.Session) FavoritesRepo {
	var f = favoritesRepository {
		cassandraSession: cassandraSession,
	}
	f.cassandraSession.Query(CreateFavoritesTable).Exec()
	return f
}

