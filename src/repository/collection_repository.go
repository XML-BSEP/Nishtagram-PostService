package repository

import (
	"context"
	"github.com/gocql/gocql"
)

const (
	CreateCollectionTable = "CREATE TABLE if not exists post_keyspace.Collections (id text, profile_id text, name text, time_of_creation timestamp, posts list<text>, " +
		"PRIMARY KEY (profile_id, name));"
)

type CollectionRepo interface {
	CreateCollection(userId string, collectionName string, ctx context.Context) error
	AddPostToCollection(userId string, collectionName string, postId string, ctx context.Context) error
	RemovePostFromCollection(userId string, collectionName string, postId string, ctx context.Context) error
}

type collectionRepository struct {
	cassandraSession *gocql.Session
}

func (c collectionRepository) CreateCollection(userId string, collectionName string, ctx context.Context) error {
	panic("implement me")
}

func (c collectionRepository) AddPostToCollection(userId string, collectionName string, postId string, ctx context.Context) error {
	panic("implement me")
}

func (c collectionRepository) RemovePostFromCollection(userId string, collectionName string, postId string, ctx context.Context) error {
	panic("implement me")
}

func NewCollectionRepository(cassandraSession *gocql.Session) CollectionRepo {
	var c = &collectionRepository{
		cassandraSession: cassandraSession,
	}
	err := c.cassandraSession.Query(CreateCollectionTable).Exec()
	if err != nil {
		return nil
	}
	return c
}
