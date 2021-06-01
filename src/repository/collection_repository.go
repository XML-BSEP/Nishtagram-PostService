package repository

import "github.com/gocql/gocql"

const (
	CreateCollectionTable = "CREATE TABLE if not exists post_service.Collections (profile_id, name, time_of_creation, posts list<int>, " +
		"PRIMARY KEY (profile_id, name));"
)

type CollectionRepo interface {
	CreateCollection(userId uint, collectionName string) error
	AddPostToCollection(userId uint, collectionName string, postId uint) error
	RemovePostFromCollection(userId uint, collectionName string, postId uint) error
}

type collectionRepository struct {
	cassandraSession *gocql.Session
}

func (c collectionRepository) CreateCollection(userId uint, collectionName string) error {
	panic("implement me")
}

func (c collectionRepository) AddPostToCollection(userId uint, collectionName string, postId uint) error {
	panic("implement me")
}

func (c collectionRepository) RemovePostFromCollection(userId uint, collectionName string, postId uint) error {
	panic("implement me")
}

func NewCollectionRepository(cassandraSession *gocql.Session) CollectionRepo {
	var c = &collectionRepository{
		cassandraSession: cassandraSession,
	}
	c.cassandraSession.Query(CreateCollectionTable).Exec()
	return c
}
