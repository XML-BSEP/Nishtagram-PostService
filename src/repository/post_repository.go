package repository

import (
	"github.com/gocql/gocql"
	"post-service/domain"
)

const (
 	CreatePostTable = "CREATE TABLE if not exists post_service.Posts (id, profile_id, description, timestamp, num_of_likes, num_of_dislikes, num_of_comments, banned, location_name, location_lat" +
 		"location_long, hashtags list<text>, media list<text>, type, PRIMARY KEY (id, profile_id)) " +
 		"WITH CLUSTERING ORDER BY (timestamp DESC);;"
)
type PostRepo interface {
	CreatePost(req *domain.Post) error
	EditPost(req *domain.Post) error
	DeletePost(req *domain.Post) error
	GetPostsForUserFeed(userId uint) ([]domain.Post, error)
}

type postRepository struct {
	cassandraSession *gocql.Session
}

func (p postRepository) GetPostsForUserFeed(userId uint) ([]domain.Post, error) {
	panic("implement me")
}

func (p postRepository) CreatePost(req *domain.Post) error {
	panic("implement me")
}

func (p postRepository) EditPost(req *domain.Post) error {
	panic("implement me")
}

func (p postRepository) DeletePost(req *domain.Post) error {
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
