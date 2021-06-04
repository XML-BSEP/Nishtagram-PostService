package repository

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"time"
)

const (
	CreateCollectionTable = "CREATE TABLE if not exists post_keyspace.Collections (id text, profile_id text, name text, time_of_creation timestamp, posts list<text>, " +
		"PRIMARY KEY (profile_id, name));"
	InsertCollectionStatement = "INSERT INTO post_keyspace.Collections (id, profile_id, name, time_of_creation, posts) VALUES (?, ?, ?, ?, ?) IF NOT EXISTS;"
	GetCollection = "SELECT posts FROM post_keyspace.Collections WHERE profile_id = ? AND name = ?;"
	UpdateCollection = "UPDATE post_keyspace.Collections SET posts = ? WHERE profile_id = ? AND name = ?;"
	DeleteCollection = "DELETE FROM post_keyspace.Collections WHERE profile_id = ? AND name = ?;"
	GetAllCollectionNames = "SELECT name FROM post_keyspace.Collections WHERE profile_id = ?;"
	)

type CollectionRepo interface {
	CreateCollection(userId string, collectionName string, ctx context.Context) error
	AddPostToCollection(userId string, collectionName string, postId string, ctx context.Context) error
	RemovePostFromCollection(userId string, collectionName string, postId string, ctx context.Context) error
	DeleteCollection(userId string, collectionName string, ctx context.Context) error
	GetCollection(userId string, collectionName string, ctx context.Context) ([]string, error)
	GetAllCollectionNames(userId string, ctx context.Context) ([]string, error)
}

type collectionRepository struct {
	cassandraSession *gocql.Session
}

func (c collectionRepository) GetAllCollectionNames(userId string, ctx context.Context) ([]string, error) {
	iter := c.cassandraSession.Query(GetAllCollectionNames, userId).Iter()
	if iter == nil {
		return nil, fmt.Errorf("no collections")
	}

	var collections []string
	scanner := iter.Scanner()
	for scanner.Next() {
		var name string
		err := scanner.Scan(&name)
		if err != nil {
			continue
		}
		collections = append(collections, name)
	}

	return collections, nil
}

func (c collectionRepository) GetCollection(userId string, collectionName string, ctx context.Context) ([]string, error) {
	var posts []string
	iter := c.cassandraSession.Query(GetCollection, userId, collectionName).Iter()

	if iter == nil {
		return nil, fmt.Errorf("no such collection")
	}

	iter.Scan(&posts)
	return posts, nil
}

func (c collectionRepository) DeleteCollection(userId string, collectionName string, ctx context.Context) error {
	err := c.cassandraSession.Query(DeleteCollection, userId, collectionName).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (c collectionRepository) CreateCollection(userId string, collectionName string, ctx context.Context) error {
	var posts []string
	collectionId, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("error while saving collection")
	}
	err = c.cassandraSession.Query(InsertCollectionStatement, collectionId, userId, collectionName, time.Now(), posts).Exec()

	if err != nil {
		return err
	}

	return nil

}

func (c collectionRepository) AddPostToCollection(userId string, collectionName string, postId string, ctx context.Context) error {
	var posts []string

	iter := c.cassandraSession.Query(GetCollection, userId, collectionName).Iter()

	if iter == nil {
		return fmt.Errorf("error while saving post to collection")
	}

	for iter.Scan(&posts) {
		posts = append(posts, postId)
	}

	err := c.cassandraSession.Query(UpdateCollection, posts, userId, collectionName).Exec()

	if err != nil {
		return err
	}

	return nil
}


func (c collectionRepository) RemovePostFromCollection(userId string, collectionName string, postId string, ctx context.Context) error {
	var posts []string

	iter := c.cassandraSession.Query(GetCollection, userId, collectionName).Iter()

	if iter == nil {
		return fmt.Errorf("error while saving post to collection")
	}

	for iter.Scan(&posts) {
		posts = remove(posts, postId)
	}

	err := c.cassandraSession.Query(UpdateCollection, posts, userId, collectionName).Exec()

	if err != nil {
		return err
	}

	return nil
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


func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}