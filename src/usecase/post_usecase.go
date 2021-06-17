package usecase

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	logger "github.com/jelena-vlajkov/logger/logger"
	"io/ioutil"
	"os"
	"post-service/domain"
	"post-service/dto"
	"post-service/gateway"
	"post-service/repository"
	"strings"
	"time"
)

type PostUseCase interface {
	AddPost(postDTO dto.CreatePostDTO, ctx context.Context) error
	DeletePost(postDTO dto.DeletePostDTO, ctx context.Context) error
	EditPost(postDTO dto.UpdatePostDTO, ctx context.Context) error
	GetPostsByUser(userId string, userRequestedId string, ctx context.Context) ([]dto.PostDTO, error)
	GetPost(postId string, userId string, userRequestedId string, ctx context.Context) (dto.PostPreviewDTO, error)
	GetPostDTO(postId string, userId string, userRequestedId string, ctx context.Context) (dto.PostDTO, error)
	GenerateUserFeed(userId string, userRequestedId string, ctx context.Context) ([]dto.PostPreviewDTO, error)
	EncodeBase64(media string, userId string, ctx context.Context) (string, error)
	DecodeBase64(media string, userId string, ctx context.Context) (string, error)
	GetPostsOnProfile(profileId string, userRequested string, ctx context.Context) ([]dto.PostInDTO, error)
	GetAllLikedMedia(profileId string, ctx context.Context) ([]dto.PostDTO, error)
	GetAllDislikedMedia(profileId string, ctx context.Context) ([]dto.PostDTO, error)
	GetPostByIdForSearch(profileId string, id string, ctx context.Context) dto.PostSearchDTO
}

type postUseCase struct {
	postRepository repository.PostRepo
	likeRepository repository.LikeRepo
	collectionRepository repository.CollectionRepo
	favoriteRepository repository.FavoritesRepo
	logger *logger.Logger
}

func (p postUseCase) GetAllLikedMedia(profileId string, ctx context.Context) ([]dto.PostDTO, error) {
	likedMedia, _ := p.likeRepository.GetLikedMedia(profileId, ctx)

	var retVal []dto.PostDTO

	for _, m := range likedMedia {
		post, err := p.GetPostDTO(m.PostId, m.PostBy.Id, m.Profile.Id, ctx)
		if err != nil {
			continue
		}
		retVal = append(retVal, post)
	}

	return retVal, nil
}

func (p postUseCase) GetAllDislikedMedia(profileId string, ctx context.Context) ([]dto.PostDTO, error) {
	likedMedia, _ := p.likeRepository.GetDislikedMedia(profileId, ctx)

	var retVal []dto.PostDTO

	for _, m := range likedMedia {
		post, err := p.GetPostDTO(m.PostId, m.PostBy.Id, m.Profile.Id, ctx)
		if err != nil {
			continue
		}
		retVal = append(retVal, post)
	}

	return retVal, nil
}

func (p postUseCase) GetPostDTO(postId string, userId string, userRequestedId string, ctx context.Context) (dto.PostDTO, error) {
	p.logger.Logger.Infof("getting post %v by user %v\n", postId, userId)
	post, err := p.postRepository.GetPostsById(userId, postId, context.Background())
	if err != nil {
		return dto.PostDTO{}, err
	}
	var mediaToAppend []string

	for _, s := range post.Media {
		base64Image, err := p.DecodeBase64(s, userId, context.Background())
		if err != nil {
			continue
		}

		mediaToAppend = append(mediaToAppend, base64Image)
	}

	appendToDescHashtags := ""
	appendToTags := ""
	if len(post.Hashtags) > 0 {
		for _, s := range post.Hashtags {
			if s != "" {
				appendToDescHashtags = appendToDescHashtags + "#" + s
			}
		}
	}

	if len(post.Mentions) > 0 {
		for _, s := range post.Mentions {
			if s != "" {
				appendToTags = appendToTags + "@" + s
			}
		}
	}

	p.logger.Logger.Infof("getting user info for %v from user ms\n", userId)
	profile, err := gateway.GetUser(context.Background(), post.Profile.Id)
	if err != nil {
		p.logger.Logger.Errorf("error while getting info for %v from user ms, error: %v\n", userId, err)
	}
	post.Profile = domain.Profile{Id: post.Profile.Id, Username: profile.Username, ProfilePhoto: profile.ProfilePhoto}

	post.Description = post.Description + "\n\n" + appendToTags + "\n\n" + appendToDescHashtags
	post.Media = mediaToAppend

	return post, nil
}

func (p postUseCase) GetPostsOnProfile(profileId string, userRequested string, ctx context.Context) ([]dto.PostInDTO, error) {
	p.logger.Logger.Infof("getting posts on user profile with id %v\n", profileId)
	if profileId != userRequested {
		userFollowing, _ := gateway.GetAllUserFollowing(context.Background(), userRequested)
		isOkay := false
		for  _, u := range userFollowing {
			if u.Id == profileId {
				isOkay = true
				break
			}
		}
		if !isOkay {
			p.logger.Logger.Errorf("user %v does not follow user %v\n", userRequested, profileId)
			return nil, fmt.Errorf("oh no i hope i don't fall")
		}
	}

	p.logger.Logger.Infof("getting all posts for user %v\n", profileId)
	posts, err := p.GetPostsByUser(profileId, userRequested, context.Background())
	if err != nil {
		return nil, err
	}
	var retVal []dto.PostInDTO
	for _, post := range posts {
		dto := dto.PostInDTO{PostId: post.Id, Posts: post.Media[0], User: post.Profile.Id}
		if post.MediaType.Type == "VIDEO" {
			dto.IsVideo = true
		}
		retVal = append(retVal, dto)
	}

	return retVal, nil
}

func (p postUseCase) DecodeBase64(media string, userId string, ctx context.Context) (string, error) {
	p.logger.Logger.Infof("decoding base64 for image %v and user %v\n", media, userId)
	workingDirectory, _ := os.Getwd()
	if !strings.HasSuffix(workingDirectory, "src") {
		firstPart := strings.Split(workingDirectory, "src")
		value := firstPart[0] + "/src"
		workingDirectory = value
		os.Chdir(workingDirectory)
	}

	path1 := "./assets/images/"
	err := os.Chdir(path1)
	fmt.Println(err)
	spliced := strings.Split(media, "/")
	var f *os.File
	if len(spliced) > 1 {
		err = os.Chdir(userId)
		f, _ = os.Open(spliced[1])
	} else {
		f, _ = os.Open(spliced[0])
	}

	defer f.Close()
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	encoded := base64.StdEncoding.EncodeToString(content)

	fmt.Println("ENCODED: " + encoded)
	os.Chdir(workingDirectory)

	return "data:image/jpg;base64," + encoded, nil
}

func (p postUseCase) GenerateUserFeed(userId string, userRequestedId string, ctx context.Context) ([]dto.PostPreviewDTO, error) {
	p.logger.Logger.Infof("generating user feed for user %v\n", userId)
	p.logger.Logger.Infof("getting all user followings for user %v from follow ms\n", userId)
	userFollowing, err := gateway.GetAllUserFollowing(context.Background(), userRequestedId)
	if err != nil {
		p.logger.Logger.Errorf("error while getting all followings for user %v from follow ms, error: %v\n", userId, err)
		return nil, err
	}


	inTimeRange := time.Now().Add(-72*time.Hour)
	postsToShow := make(map[string][]string, len(userFollowing))
	for _, user := range userFollowing {
		posts := p.postRepository.GetPostsInDateTimeRange(user.Id, inTimeRange, context.Background())
		if posts != nil {
			postsToShow[user.Id] = append(postsToShow[user.Id], posts...)
		}
	}

	var postsPreview []dto.PostPreviewDTO
	for idFollowing := range postsToShow {
		for _, postById := range postsToShow[idFollowing] {
			post, err := p.postRepository.GetPostsById(idFollowing, postById, context.Background())
			if err != nil {
				p.logger.Logger.Errorf("error while getting post %v by user %v, error: %v\n", idFollowing, postById, err)
				continue
			}
			var mediaToAppend []string
			for _, media := range post.Media {
				base64Image, err := p.DecodeBase64(media, post.Profile.Id, context.Background())
				if err != nil {
					continue
				}
				mediaToAppend = append(mediaToAppend, base64Image)
			}

			post.Media = mediaToAppend
			if post.MediaType.Type == "VIDEO" {
				post.IsVideo = true
			}
			if post.MediaType.Type == "IMAGE" {
				if len(post.Media) > 1 {
					post.IsAlbum = true
				}
			}

			appendToDescHashtags := ""
			appendToTags := ""
			if len(post.Hashtags) > 0 {
				for _, s := range post.Hashtags {
					if s != "" {
						appendToDescHashtags = appendToDescHashtags + "#" + s
					}
				}
			}

			if len(post.Mentions) > 0 {
				for _, s := range post.Mentions {
					if s != "" {
						appendToTags = appendToTags + "@" + s
					}
				}
			}

			if p.likeRepository.SeeIfLikeExists(post.Id, userRequestedId, context.Background()) {
				post.IsLiked = true
			}

			if p.likeRepository.SeeIfDislikeExists(post.Id, userRequestedId, context.Background()) {
				post.IsDisliked = true
			}

			favorites, err := p.favoriteRepository.GetFavorites(userRequestedId)
			if _, ok := favorites[post.Id]; ok {
				post.IsBookmarked = true
			}



			profile, err := gateway.GetUser(context.Background(), post.Profile.Id)
			if err != nil {
				p.logger.Logger.Errorf("error while getting user info for %v, error: %v\n", post.Profile.Id, err)
			}
			post.Profile = domain.Profile{Id: post.Profile.Id, Username: profile.Username, ProfilePhoto: profile.ProfilePhoto}
			post.Description = post.Description + "\n\n" + appendToTags + "\n\n" + appendToDescHashtags

			postsPreview = append(postsPreview, dto.NewPostPreviewDTO(post))

		}
	}

	return postsPreview, nil
}

func (p postUseCase) EncodeBase64(media string, userId string, ctx context.Context) (string, error) {
	p.logger.Logger.Infof("encoding base64 image for userId %v\n", userId)
	workingDirectory, _ := os.Getwd()
	if !strings.HasSuffix(workingDirectory, "src") {
		firstPart := strings.Split(workingDirectory, "src")
		value := firstPart[0] + "/src"
		workingDirectory = value
		os.Chdir(workingDirectory)
	}
	path1 := "./assets/images/"
	err := os.Chdir(path1)
	if err != nil {
		p.logger.Logger.Errorf("error while encoding base64 image for userId %v, error: %v\n", userId, err)
	}
	err = os.Mkdir(userId, 0755)
	fmt.Println(err)

	err = os.Chdir(userId)
	fmt.Println(err)


	s := strings.Split(media, ",")
	a := strings.Split(s[0], "/")
	format := strings.Split(a[1], ";")

	dec, err := base64.StdEncoding.DecodeString(s[1])

	if err != nil {
		p.logger.Logger.Errorf("error while encoding base64 image for userId %v, error: %v\n", userId, err)
	}
	uuid := uuid.NewString()
	f, err := os.Create(uuid + "." + format[0])

	if err != nil {
		p.logger.Logger.Errorf("error while encoding base64 image for userId %v, error: %v\n", userId, err)
	}

	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		p.logger.Logger.Errorf("error while encoding base64 image for userId %v, error: %v\n", userId, err)
	}
	if err := f.Sync(); err != nil {
		p.logger.Logger.Errorf("error while encoding base64 image for userId %v, error: %v\n", userId, err)
	}

	os.Chdir(workingDirectory)
	return userId + "/" + uuid + "." + format[0], nil
}


func (p postUseCase) AddPost(postDTO dto.CreatePostDTO, ctx context.Context) error {
	p.logger.Logger.Infof("adding post for user %v\n", postDTO.UserId)
	var media []string

	if postDTO.IsImage  {
		media = make([]string, 1)
		mediaToAttach, err := p.EncodeBase64(postDTO.Image, postDTO.UserId.UserId, context.Background())
		if err != nil {
			return fmt.Errorf("error while decoding base64")
		}
		media[0] = mediaToAttach
		postDTO.Media = media
		postDTO.MediaType = "IMAGE"
	}
	if postDTO.IsVideo {
		media = make([]string, 1)
		mediaToAttach, err := p.EncodeBase64(postDTO.Video, postDTO.UserId.UserId, context.Background())
		if err != nil {
			return fmt.Errorf("error while decoding base64")
		}
		media[0] = mediaToAttach
		postDTO.Media = media
		postDTO.MediaType = "VIDEO"
	}
	if postDTO.IsAlbum{
		for _, s := range postDTO.Album {
			mediaToAttach, err := p.EncodeBase64(s, postDTO.UserId.UserId, context.Background())
			if err != nil {
				return fmt.Errorf("error while decoding base64")
			}
			media = append(media, mediaToAttach)
		}

		postDTO.Media = media
		postDTO.MediaType = "IMAGE"
	}

	return p.postRepository.CreatePost(postDTO, context.Background())
}

func (p postUseCase) DeletePost(postDTO dto.DeletePostDTO, ctx context.Context) error {
	p.logger.Logger.Infof("deleting post with id %v for user %v\n", postDTO.PostId, postDTO.PostId)
	return p.postRepository.DeletePost(postDTO, context.Background())

}

func (p postUseCase) EditPost(postDTO dto.UpdatePostDTO, ctx context.Context) error {
	p.logger.Logger.Infof("editing post with id %v for user %v\n", postDTO.PostId, postDTO.UserId)
	return p.postRepository.EditPost(postDTO, context.Background())

}

func (p postUseCase) GetPostsByUser(userId string, userRequestedId string, ctx context.Context) ([]dto.PostDTO, error) {
	p.logger.Logger.Infof("getting posts for user %v\n", userId)
	posts, err := p.postRepository.GetPostsByUserId(userId, context.Background())
	var retVal []dto.PostDTO

	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		converted, err := p.GetPostDTO(post.Id, post.Profile.Id, userRequestedId, context.Background())
		if err != nil {
			continue
		}
		retVal = append(retVal, converted)
	}

	return retVal, nil
}

func (p postUseCase) GetPost(postId string, userId string, userRequestedId string, ctx context.Context) (dto.PostPreviewDTO, error) {
	p.logger.Logger.Infof("getting post with id %v by user %v\n", postId, userId)
	post, err := p.postRepository.GetPostsById(userId, postId, context.Background())
	if err != nil {
		return dto.PostPreviewDTO{}, err
	}
	var mediaToAppend []string

	for _, s := range post.Media {
		base64Image, err := p.DecodeBase64(s, userId, context.Background())
		if err != nil {
			continue
		}

		mediaToAppend = append(mediaToAppend, base64Image)
	}

	appendToDescHashtags := ""
	appendToTags := ""
	if len(post.Hashtags) > 0 {
		for _, s := range post.Hashtags {
			if s != "" {
				appendToDescHashtags = appendToDescHashtags + "#" + s
			}
		}
	}

	if len(post.Mentions) > 0 {
		for _, s := range post.Mentions {
			if s != "" {
				appendToTags = appendToTags + "@" + s
			}
		}
	}

	if post.MediaType.Type == "VIDEO" {
		post.IsVideo = true
	}

	if post.MediaType.Type == "ALBUM" || len(post.Media) > 1 {
		post.IsAlbum = true
	}

	if p.likeRepository.SeeIfLikeExists(post.Id, userRequestedId, context.Background()) {
		post.IsLiked = true
	}

	if p.likeRepository.SeeIfDislikeExists(post.Id, userRequestedId, context.Background()) {
		post.IsDisliked = true
	}

	favorites, err := p.favoriteRepository.GetFavorites(userRequestedId)
	if _, ok := favorites[post.Id]; ok {
		post.IsBookmarked = true
	}



	post.Description = post.Description + "\n\n" + appendToTags + "\n\n" + appendToDescHashtags
	post.Media = mediaToAppend


	return dto.NewPostPreviewDTO(post), nil

}

func (p postUseCase) GetPostByIdForSearch(profileId string, id string, ctx context.Context) dto.PostSearchDTO {
	post, profileId := p.postRepository.GetPostByIdForSearch(profileId, id, ctx)

	for i, postM := range post.Media {
		base64Image, err := p.DecodeBase64(postM, profileId, context.Background())
		if err != nil {
			panic(err)
		}
		post.Media[i] = base64Image
	}

	profile, err := gateway.GetUser(context.Background(), profileId)
	if err != nil {
		p.logger.Logger.Errorf("error while getting info for %v from user ms, error: %v\n", profileId, err)
	}

	post.Username = profile.Username
	post.ProfilePhoto = profile.ProfilePhoto



	return post

}

func NewPostUseCase(postRepository repository.PostRepo, repo repository.LikeRepo, favoritesRepo repository.FavoritesRepo, collectionRepo repository.CollectionRepo, logger *logger.Logger) PostUseCase {
	return &postUseCase{
		postRepository: postRepository,
		likeRepository: repo,
		favoriteRepository: favoritesRepo,
		collectionRepository: collectionRepo,
		logger: logger,
	}
}
