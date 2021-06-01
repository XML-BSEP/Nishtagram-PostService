package repository

import (
	"context"
	"github.com/gocql/gocql"
	"post-service/domain"
)

const (
 	CreatePostTable = "CREATE TABLE if not exists post_keyspace.Posts (id text, profile_id text, description text, timestamp timestamp, num_of_likes int, num_of_dislikes int, num_of_comments int, banned boolean, location_name text, location_lat double," +
 		"location_long double, hashtags list<text>, media list<text>, type text, PRIMARY KEY (id, profile_id));"
 	InsertIntoPostTable = "INSERT INTO post_keyspace.Posts (id, profile_id, description, timestamp, num_of_likes, num_of_dislikes, num_of_comments, banned, location_name, location_lat, location_long, hashtags, media, type) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) IF NOT EXISTS;"
	AddLikeToPost = "UPDATE post_keyspace.Posts SET num_of_likes = ? WHERE id = ? and profile_id = ?;"
	AddDislikeToPost = "UPDATE post_keyspace.Posts SET num_of_dislikes = ? WHERE id = ? and profile_id = ?;"
	AddCommentToPost = "UPDATE post_keyspace.Posts SET num_of_comments = ?  WHERE id = ? and profile_id = ?;"
	RemoveLikeFromPost = "UPDATE post_keyspace.Posts SET num_of_likes = ? WHERE id = ? and profile_id = ? and;"
	RemoveDislikeFromPost = "UPDATE post_keyspace.Posts SET num_of_dislikes = ?  WHERE id = ? and profile_id = ?;"
	RemoveCommentFromPost = "UPDATE post_keyspace.Posts SET num_of_comments = ? WHERE id = ? and profile_id = ?;"
	GetPrimaryKeysById = "SELECT id, profile_id, timestamp FROM post_keyspace.Posts where id = ?;"
	)
type PostRepo interface {
	CreatePost(req *domain.Post, ctx context.Context) error
	EditPost(req *domain.Post, ctx context.Context) error
	DeletePost(req *domain.Post, ctx context.Context) error
	GetPostsForUserFeed(userId string, ctx context.Context) ([]domain.Post, error)
}

type postRepository struct {
	cassandraSession *gocql.Session
}

func (p postRepository) GetPostsForUserFeed(userId string, ctx context.Context) ([]domain.Post, error) {
	panic("implement me")
}

func (p postRepository) CreatePost(req *domain.Post, ctx context.Context) error {
	panic("implement me")
}

func (p postRepository) EditPost(req *domain.Post, ctx context.Context) error {
	panic("implement me")
}

func (p postRepository) DeletePost(req *domain.Post, ctx context.Context) error {
	panic("implement me")
}

func NewPostRepository(cassandraSession *gocql.Session) PostRepo {
	var p = &postRepository{
		cassandraSession: cassandraSession,
	}
	err := p.cassandraSession.Query(CreatePostTable).Exec()
	if err != nil {
		return nil
	}
	return p
}
