package repository

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"post-service/dto"
	"time"
)

const (
 	CreatePostTable = "CREATE TABLE if not exists post_keyspace.Posts (id text, profile_id text, description text, timestamp timestamp, num_of_likes int, num_of_dislikes int, num_of_comments int, banned boolean, location_name text, location_lat double," +
 		"location_long double, mentions list<text>, hashtags list<text>, media list<text>, type text, deleted boolean, PRIMARY KEY (profile_id, id));"
 	InsertIntoPostTable = "INSERT INTO post_keyspace.Posts (id, profile_id, description, timestamp, num_of_likes, " +
 		"num_of_dislikes, num_of_comments, banned, location_name, location_lat, location_long, mentions, hashtags, media, type, deleted) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) IF NOT EXISTS;"
	AddLikeToPost = "UPDATE post_keyspace.Posts SET num_of_likes = ? WHERE id = ? and profile_id = ?;"
	AddDislikeToPost = "UPDATE post_keyspace.Posts SET num_of_dislikes = ? WHERE id = ? and profile_id = ?;"
	AddCommentToPost = "UPDATE post_keyspace.Posts SET num_of_comments = ?  WHERE id = ? and profile_id = ?;"
	RemoveLikeFromPost = "UPDATE post_keyspace.Posts SET num_of_likes = ? WHERE id = ? and profile_id = ? and;"
	RemoveDislikeFromPost = "UPDATE post_keyspace.Posts SET num_of_dislikes = ?  WHERE id = ? and profile_id = ?;"
	RemoveCommentFromPost = "UPDATE post_keyspace.Posts SET num_of_comments = ? WHERE id = ? and profile_id = ?;"
	GetNumOfLikesForPost = "SELECT num_of_likes FROM post_keyspace.Posts WHERE id = ? AND profile_id = ?;"
	GetNumOfDislikesForPost = "SELECT num_of_dislikes FROM post_keyspace.Posts WHERE id = ? AND profile_id = ?;"
	GetNumOfCommentsForPost = "SELECT num_of_comments FROM post_keyspace.Posts WHERE id = ? AND profile_id = ?;"
	GetPrimaryKeysById = "SELECT id, profile_id, timestamp FROM post_keyspace.Posts where id = ?;"
	GetPostsByUserId = "SELECT id, profile_id, description, timestamp, num_of_likes, num_of_dislikes, num_of_comments, banned, location_name, mentions, " +
		"hashtags, media, type, deleted FROM post_keyspace.Posts where profile_id = ?;"
	GetPostsForById = "SELECT id, profile_id, description, timestamp, num_of_likes, num_of_dislikes, num_of_comments, banned, location_name, mentions, " +
		"hashtags, media, type, deleted  FROM post_keyspace.Posts where profile_id = ? and id = ?;"
	UpdatePost = "UPDATE post_keyspace.Posts SET description = ?, mentions = ?, hashtags = ?, location_name = ?, location_lat = ?, location_long = ? WHERE profile_id = ? and id = ?;"
	DeletePost = "UPDATE post_keyspace.Posts SET deleted = true WHERE profile_id = ? AND id = ?;"
	)
type PostRepo interface {
	CreatePost(req dto.CreatePostDTO, ctx context.Context) error
	EditPost(req dto.UpdatePostDTO, ctx context.Context) error
	DeletePost(req dto.DeletePostDTO, ctx context.Context) error
	GetPostsByUserId(userId string, ctx context.Context) ([]dto.PostDTO, error)
	GetPostsById(userId string, postId string) (dto.PostDTO, error)
}

type postRepository struct {
	cassandraSession *gocql.Session
}

func (p postRepository) GetPostsById(userId string, postId string) (dto.PostDTO, error) {
	var id, profileId, description, location, postType string
	var numOfLikes, numOfDislikes, numOfComments int
	var banned, deleted bool
	var timestamp time.Time

	iter := p.cassandraSession.Query(GetPostsForById, userId, postId).Iter()
	var post dto.PostDTO

	if iter == nil {
		return post, fmt.Errorf("no such element")
	}

	var hashtags, media, mentions []string
	for iter.Scan(&id, &profileId, &description, &timestamp, &numOfLikes,
		&numOfDislikes, &numOfComments, &banned, &location, &mentions, &hashtags, &media, &postType, &deleted) {
		if !deleted && !banned {
			return dto.NewPost(id, description, timestamp, numOfLikes, numOfDislikes,
				numOfComments, profileId, location, mentions, hashtags, media, postType), nil
		}
	}

	return post, fmt.Errorf("no such element")
}

func (p postRepository) GetPostsByUserId(userId string, ctx context.Context) ([]dto.PostDTO, error) {

	var id, profileId, description, location, postType string
	var numOfLikes, numOfDislikes, numOfComments int
	var banned, deleted bool
	var timestamp time.Time


	iter := p.cassandraSession.Query(GetPostsByUserId, userId).Iter().Scanner()
	var posts []dto.PostDTO
	if iter == nil {
		return nil, iter.Err()
	}

	for iter.Next() {
		var hashtags, media, mentions []string
		iter.Scan(&id, &profileId, &description, &timestamp, &numOfLikes,
			&numOfDislikes, &numOfComments, &banned, &location, &mentions, &hashtags, &media, &postType, &deleted)
		if !deleted && !banned {
			posts = append(posts, dto.NewPost(id, description, timestamp, numOfLikes, numOfDislikes,
				numOfComments, profileId, location, mentions, hashtags, media, postType))
		}
	}

	return posts, nil
}


func (p postRepository) CreatePost(req dto.CreatePostDTO, ctx context.Context) error {
	postId, err := uuid.NewUUID()

	if err != nil {
		return err
	}
	currentTime := time.Now()
	err = p.cassandraSession.Query(InsertIntoPostTable, postId, req.UserId, req.Description, currentTime, 0, 0, 0, false, req.Location.Location,
		req.Location.Latitude, req.Location.Longitude, req.Mentions, req.Hashtags, req.Media, req.MediaType, false).Exec()

	if err != nil {
		return fmt.Errorf("error while saving post")
	}
	return nil
}

func (p postRepository) EditPost(req dto.UpdatePostDTO, ctx context.Context) error {
	err := p.cassandraSession.Query(UpdatePost, req.Description, req.Mentions, req.Hashtags,
		req.Location.Location, req.Location.Latitude, req.Location.Longitude).Exec()

	if err != nil {
		return fmt.Errorf("error while updating post")
	}

	return nil
}

func (p postRepository) DeletePost(req dto.DeletePostDTO, ctx context.Context) error {
	err := p.cassandraSession.Query(DeletePost, req.UserId, req.PostId).Exec()

	if err != nil {
		return fmt.Errorf("error while deleting post")
	}

	return nil
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