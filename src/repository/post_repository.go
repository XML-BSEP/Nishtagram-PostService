package repository

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	logger "github.com/jelena-vlajkov/logger/logger"
	"post-service/dto"
	"post-service/gateway"
	"time"
)

const (
 	CreatePostTable = "CREATE TABLE if not exists post_keyspace.Posts (id text, profile_id text, description text, timestamp timestamp, num_of_likes int, num_of_dislikes int, num_of_comments int, banned boolean, location_name text, location_lat double," +
 		"location_long double, mentions list<text>, hashtags list<text>, media list<text>, type text, deleted boolean, PRIMARY KEY (profile_id, id));"
	CreatePostsTimestampTable = "CREATE TABLE if not exists post_keyspace.PostsTimestamp (post_id text, profile_id text, timestamp timestamp, PRIMARY KEY (profile_id, timestamp)) WITH CLUSTERING ORDER BY (timestamp ASC);"
 	InsertIntoPostTable = "INSERT INTO post_keyspace.Posts (id, profile_id, description, timestamp, num_of_likes, " +
 		"num_of_dislikes, num_of_comments, banned, location_name, location_lat, location_long, mentions, hashtags, media, type, deleted) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) IF NOT EXISTS;"
	InsertIntoPostsTimestampTable = "INSERT INTO post_keyspace.PostsTimestamp (post_id, profile_id, timestamp) VALUES (?, ?, ?);"
	GetPostsInLastThreeDays = "SELECT post_id from post_keyspace.PostsTimestamp WHERE profile_id = ? AND timestamp >= ?;"
 	AddLikeToPost = "UPDATE post_keyspace.Posts SET num_of_likes = ? WHERE id = ? and profile_id = ?;"
	AddDislikeToPost = "UPDATE post_keyspace.Posts SET num_of_dislikes = ? WHERE id = ? and profile_id = ?;"
	AddCommentToPost = "UPDATE post_keyspace.Posts SET num_of_comments = ?  WHERE id = ? and profile_id = ?;"
	RemoveLikeFromPost = "UPDATE post_keyspace.Posts SET num_of_likes = ? WHERE id = ? and profile_id = ?;"
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
	UpdatePost        = "UPDATE post_keyspace.Posts SET description = ?, mentions = ?, hashtags = ?, location_name = ?, location_lat = ?, location_long = ? WHERE profile_id = ? and id = ?;"
	DeletePost        = "UPDATE post_keyspace.Posts SET deleted = true WHERE profile_id = ? AND id = ?;"
	IfDeletedOrBanned = "SELECT banned, deleted FROM post_keyspace.Posts WHERE profile_id = ? AND id = ?;"
	SeeIfPostExists   = "SELECT count(*) FROM post_keyspace.Posts WHERE id = ? AND profile_id = ?;"
	GetPostByIdForSearch = "SELECT id, profile_id, type, media, deleted, banned FROM post_keyspace.Posts WHERE profile_id = ? and id = ?;"

	)
type PostRepo interface {
	CreatePost(req dto.CreatePostDTO, ctx context.Context) error
	EditPost(req dto.UpdatePostDTO, ctx context.Context) error
	DeletePost(req dto.DeletePostDTO, ctx context.Context) error
	GetPostsByUserId(userId string, ctx context.Context) ([]dto.PostDTO, error)
	GetPostsById(userId string, postId string, ctx context.Context) (dto.PostDTO, error)
	SeeIfPostDeletedOrBanned(userId string, postId string, ctx context.Context) bool
	GetPostsInDateTimeRange(userId string, timeRange time.Time, ctx context.Context) []string
	GetPostByIdForSearch(profileId string, id string, ctx context.Context) (dto.PostSearchDTO, string)
	GetPostForAdmin(userId string, postId string, ctx context.Context) (dto.PostDTO, error)

}

type postRepository struct {
	cassandraSession *gocql.Session
	logger *logger.Logger
}

func (p postRepository) GetPostForAdmin(userId string, postId string, ctx context.Context) (dto.PostDTO, error) {
	var id, profileId, description, location, postType string
	var numOfLikes, numOfDislikes, numOfComments int
	var banned, deleted bool
	var timestamp time.Time

	iter := p.cassandraSession.Query(GetPostsForById, userId, postId).Iter()
	var post dto.PostDTO

	if iter == nil {
		p.logger.Logger.Errorf("error while getting post %v by user %v\n", postId, userId)
		return post, fmt.Errorf("no such element")
	}

	var hashtags, media, mentions []string
	for iter.Scan(&id, &profileId, &description, &timestamp, &numOfLikes,
		&numOfDislikes, &numOfComments, &banned, &location, &mentions, &hashtags, &media, &postType, &deleted) {
		return dto.NewPost(id, description, timestamp, numOfLikes, numOfDislikes,
			numOfComments, profileId, location, mentions, hashtags, media, postType), nil

	}

	return post, fmt.Errorf("no such element")}

func (p postRepository) GetPostsInDateTimeRange(userId string, timeRange time.Time, ctx context.Context) []string {
	var posts []string
	iter := p.cassandraSession.Query(GetPostsInLastThreeDays, userId, timeRange).Iter().Scanner()
	var post string
	for iter.Next() {
		iter.Scan(&post)
		posts = append(posts, post)
	}
	return posts
}

func (p postRepository) SeeIfPostDeletedOrBanned(postId string, userId string, ctx context.Context) bool {
	var banned, deleted bool
	p.cassandraSession.Query(IfDeletedOrBanned, userId, postId).Iter().Scan(&banned, &deleted)
	return banned || deleted
}

func (p postRepository) GetPostsById(userId string, postId string, ctx context.Context) (dto.PostDTO, error) {
	var id, profileId, description, location, postType string
	var numOfLikes, numOfDislikes, numOfComments int
	var banned, deleted bool
	var timestamp time.Time

	iter := p.cassandraSession.Query(GetPostsForById, userId, postId).Iter()
	var post dto.PostDTO

	if iter == nil {
		p.logger.Logger.Errorf("error while getting post %v by user %v\n", postId, userId)
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
		p.logger.Logger.Errorf("error while getting posts for user %v, error: no posts\n", userId)
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
	postId := uuid.NewString()

	currentTime := time.Now()

	err := p.cassandraSession.Query(InsertIntoPostTable, postId, req.UserId.UserId, req.Caption, currentTime, 0, 0, 0, false, req.Location,
		0.0, 0.0, req.MentionsToAdd, req.Hashtags, req.Media, req.MediaType, false).Exec()

	err = p.cassandraSession.Query(InsertIntoPostsTimestampTable, postId, req.UserId.UserId, currentTime).Exec()

	if err != nil {
		p.logger.Logger.Errorf("error while saving post for user %v\n", req.UserId)
		return fmt.Errorf("error while saving post")
	}
	if req.Location != "" {
		postLocation := dto.PostLocationProfileDTO{PostId: postId, ProfileId: req.UserId.UserId, Location: req.Location}
		err := gateway.SaveNewPostLocation(postLocation, ctx)
		if err != nil {
			p.logger.Logger.Errorf("error while sending request to save new locations")
		}
	}

	if len(req.Hashtags) > 0 {
		postTag := dto.PostTagProfileDTO{PostId: postId, ProfileId: req.UserId.UserId, Hashtag: req.Hashtags}
		err := gateway.SaveNewPostTage(postTag, ctx)

		if err != nil {
			p.logger.Logger.Errorf("error while sending request to save new tags")
		}
	}
	return nil
}

func (p postRepository) EditPost(req dto.UpdatePostDTO, ctx context.Context) error {
	err := p.cassandraSession.Query(UpdatePost, req.Description, req.Mentions, req.Hashtags,
		req.Location.Location, req.Location.Latitude, req.Location.Longitude, req.UserId, req.PostId).Exec()

	if err != nil {
		p.logger.Logger.Errorf("error while editing post for user %v\n", req.UserId)
		return fmt.Errorf("error while updating post")
	}

	return nil
}

func (p postRepository) DeletePost(req dto.DeletePostDTO, ctx context.Context) error {
	err := p.cassandraSession.Query(DeletePost, req.UserId, req.PostId).Exec()

	if err != nil {
		p.logger.Logger.Errorf("error while deleting post with id %v for user %v, error: %v\n", req.PostId, req.UserId, err)
		return fmt.Errorf("error while deleting post")
	}

	return nil
}

func (p postRepository) GetPostByIdForSearch(profileId string, id string, ctx context.Context) (dto.PostSearchDTO, string) {

	iter := p.cassandraSession.Query(GetPostByIdForSearch, profileId, id).Iter().Scanner()
	var postSearch dto.PostSearchDTO
	var idp string
	var profileIdp string
	var types string
	var media []string
	var deleted bool
	var banned bool

	for iter.Next() {
		iter.Scan(&idp, &profileIdp, &types, &media, &banned, &deleted)
		if banned == false && deleted == false {
			postSearch = dto.NewPostSearchDTO(idp, types, media)
		}
	}

	return postSearch, profileIdp

}

func NewPostRepository(cassandraSession *gocql.Session, logger *logger.Logger) PostRepo {
	var p = &postRepository{
		cassandraSession: cassandraSession,
		logger: logger,
	}
	err := p.cassandraSession.Query(CreatePostTable).Exec()
	err = p.cassandraSession.Query(CreatePostsTimestampTable).Exec()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return p
}
