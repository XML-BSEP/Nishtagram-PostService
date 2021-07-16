package repository

import (
	"context"
	"github.com/gocql/gocql"
	"time"
)

const (
	CreateMutedContent = "CREATE TABLE IF NOT EXISTS post_keyspace.MutedContent (blocked_for text, blocked text, timestamp timestamp, PRIMARY KEY (blocked_for, blocked));"
	InsertIntoMutedContent = "INSERT INTO post_keyspace.MutedContent (blocked_for, blocked, timestamp) VALUES (?, ?, ?) IF NOT EXISTS;"
	DeleteMutedContent = "DELETE FROM post_keyspace.MutedContent WHERE blocked_for = ? AND blocked = ?;"
	SelectMutedContent = "SELECT count(*) FROM post_keyspace.MutedContent WHERE blocked_for = ? AND blocked = ?;"
)

type MutedContentRepo interface {
	AddMuted(blockedFor string, blocked string, ctx context.Context) error
	DeleteFrom(blockedFor string, blocked string, ctx context.Context) error
	SeeIfMuted(blockedFor string, blocked string, ctx context.Context) bool
}

type mutedContentRepository struct {
	cassandraClient *gocql.Session
}

func (m mutedContentRepository) SeeIfMuted(blockedFor string, blocked string, ctx context.Context) bool {
	var muted int
	m.cassandraClient.Query(SelectMutedContent, blockedFor, blocked).Iter().Scan(&muted)
	return muted > 0
}

func (m mutedContentRepository) AddMuted(blockedFor string, blocked string, ctx context.Context) error {
	return m.cassandraClient.Query(InsertIntoMutedContent, blockedFor, blocked, time.Now()).Exec()
}

func (m mutedContentRepository) DeleteFrom(blockedFor string, blocked string, ctx context.Context) error {
	return m.cassandraClient.Query(DeleteMutedContent, blockedFor, blocked).Exec()
}

func NewMutedContentRepo(cassandraClient *gocql.Session) MutedContentRepo {
	err := cassandraClient.Query(CreateMutedContent).Exec()
	if err != nil {
		return nil
	}
	return &mutedContentRepository{cassandraClient: cassandraClient}
}