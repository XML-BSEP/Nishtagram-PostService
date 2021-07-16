package repository

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	logger "github.com/jelena-vlajkov/logger/logger"
	"post-service/dto"
	"time"
)

const (
	CreateCollectionTable = "CREATE TABLE if not exists post_keyspace.Collections (id text, profile_id text, name text, time_of_creation timestamp, posts map<text, text>, " +
		"PRIMARY KEY (profile_id, name));"
	InsertCollectionStatement = "INSERT INTO post_keyspace.Collections (id, profile_id, name, time_of_creation, posts) VALUES (?, ?, ?, ?, ?) IF NOT EXISTS;"
	GetCollection = "SELECT posts FROM post_keyspace.Collections WHERE profile_id = ? AND name = ?;"
	UpdateCollection = "UPDATE post_keyspace.Collections SET posts = ? WHERE profile_id = ? AND name = ?;"
	DeleteCollection = "DELETE FROM post_keyspace.Collections WHERE profile_id = ? AND name = ?;"
	GetAllCollectionNames = "SELECT name FROM post_keyspace.Collections WHERE profile_id = ?;"
	)

type CollectionRepo interface {
	CreateCollection(userId string, collectionName string, ctx context.Context) error
	AddPostToCollection(userId string, collectionName string, postId string, postBy string, ctx context.Context) error
	RemovePostFromCollection(userId string, collectionName string, postId string, ctx context.Context) error
	DeleteCollection(userId string, collectionName string, ctx context.Context) error
	GetCollection(userId string, collectionName string, ctx context.Context) (map[string]string, error)
	GetAllCollectionNames(userId string, ctx context.Context) ([]string, error)
	GetAllCollections(userId string, ctx context.Context) ([]dto.CollectionDTO, error)
}

type collectionRepository struct {
	cassandraSession *gocql.Session
	logger *logger.Logger
}

func (c collectionRepository) GetAllCollections(userId string, ctx context.Context) ([]dto.CollectionDTO, error) {
	iter := c.cassandraSession.Query(GetAllCollectionNames, userId).Iter()
	if iter == nil {
		c.logger.Logger.Errorf("no collections for user %v\n", userId)
		return nil, fmt.Errorf("no collections")
	}

	var collections []dto.CollectionDTO

	scanner := iter.Scanner()
	for scanner.Next() {
		var name string
		err := scanner.Scan(&name)
		if err != nil {
			continue
		}
		collections = append(collections, dto.CollectionDTO{CollectionName: name})
	}

	return collections, nil
}

func (c collectionRepository) GetAllCollectionNames(userId string, ctx context.Context) ([]string, error) {
	iter := c.cassandraSession.Query(GetAllCollectionNames, userId).Iter()

	if iter == nil {
		c.logger.Logger.Errorf("no collections for user %v\n", userId)
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

func (c collectionRepository) GetCollection(userId string, collectionName string, ctx context.Context) (map[string]string, error) {
	var posts map[string]string
	iter := c.cassandraSession.Query(GetCollection, userId, collectionName).Iter()

	if iter == nil {
		c.logger.Logger.Errorf("no collection with name %v for user %v\n", collectionName, userId)
		return nil, fmt.Errorf("no such collection")
	}

	iter.Scan(&posts)
	return posts, nil
}

func (c collectionRepository) DeleteCollection(userId string, collectionName string, ctx context.Context) error {
	err := c.cassandraSession.Query(DeleteCollection, userId, collectionName).Exec()

	if err != nil {
		c.logger.Logger.Errorf("error while deleting collection with name %v for user %v, error: %v\n", collectionName, userId, err)
		return err
	}

	return nil
}

func (c collectionRepository) CreateCollection(userId string, collectionName string, ctx context.Context) error {
	var posts []string
	collectionId, err := uuid.NewUUID()
	if err != nil {
		c.logger.Logger.Errorf("error while saving collection with name %v for user %v, error: %v\n", collectionName, userId, err)
		return fmt.Errorf("error while saving collection")
	}
	err = c.cassandraSession.Query(InsertCollectionStatement, collectionId, userId, collectionName, time.Now(), posts).Exec()

	if err != nil {
		return err
	}

	return nil

}

func (c collectionRepository) AddPostToCollection(userId string, collectionName string, postId string, postBy string, ctx context.Context) error {
	var posts map[string]string

	iter := c.cassandraSession.Query(GetCollection, userId, collectionName).Iter()

	if iter == nil {
		c.logger.Logger.Errorf("error while saving post %v to collection with name %v for user %v\n", postId, collectionName, userId)
		return fmt.Errorf("error while saving post to collection")
	}

	for iter.Scan(&posts) {
		posts[postId] = postBy
	}

	err := c.cassandraSession.Query(UpdateCollection, posts, userId, collectionName).Exec()

	if err != nil {
		c.logger.Logger.Errorf("error while saving post %v to collection with name %v for user %v, error:  %v\n", postId, collectionName, userId, err)
		return err
	}

	return nil
}


func (c collectionRepository) RemovePostFromCollection(userId string, collectionName string, postId string, ctx context.Context) error {
	var posts map[string]string

	iter := c.cassandraSession.Query(GetCollection, userId, collectionName).Iter()

	if iter == nil {
		c.logger.Logger.Errorf("error while removing post %v to collection with name %v for user %v\n", postId, collectionName, userId)
		return fmt.Errorf("error while saving post to collection")
	}
	var newPosts map[string]string
	for iter.Scan(&posts) {
		for s := range posts {
			if s == postId {
				continue
			}
			newPosts[s] = posts[s]
		}
	}

	err := c.cassandraSession.Query(UpdateCollection, posts, userId, collectionName).Exec()

	if err != nil {
		c.logger.Logger.Errorf("error while removing post %v to collection with name %v for user %v, error:  %v\n", postId, collectionName, userId, err)
		return err
	}

	return nil
}

func NewCollectionRepository(cassandraSession *gocql.Session, logger *logger.Logger) CollectionRepo {
	var c = &collectionRepository{
		cassandraSession: cassandraSession,
		logger: logger,
	}
	err := c.cassandraSession.Query(CreateCollectionTable).Exec()
	if err != nil {
		return nil
	}
	return c
}
