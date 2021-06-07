package usecase

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
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
	GetPost(postId string, userId string, userRequestedId string, ctx context.Context) (dto.PostDTO, error)
	GenerateUserFeed(userId string, userRequestedId string, ctx context.Context) ([]dto.PostPreviewDTO, error)
	EncodeBase64(media string, userId string, ctx context.Context) (string, error)
	DecodeBase64(media string, userId string, ctx context.Context) (string, error)
}

type postUseCase struct {
	postRepository repository.PostRepo
	likeRepository repository.LikeRepo
	collectionRepository repository.CollectionRepo
	favoriteRepository repository.FavoritesRepo
}

func (p postUseCase) DecodeBase64(media string, userId string, ctx context.Context) (string, error) {
	workingDirectory, _ := os.Getwd()

	path1 := "./assets/images/"
	err := os.Chdir(path1)
	fmt.Println(err)

	err = os.Chdir(userId)
	spliced := strings.Split(media, "/")
	var f *os.File
	if len(spliced) > 1 {
		f, _ = os.Open(spliced[1])
	} else {
		f, _ = os.Open(spliced[0])
	}


	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)


	encoded := base64.StdEncoding.EncodeToString(content)


	fmt.Println("ENCODED: " + encoded)
	os.Chdir(workingDirectory)

	return "data:image/jpg;base64," + encoded, nil
}

func (p postUseCase) GenerateUserFeed(userId string, userRequestedId string, ctx context.Context) ([]dto.PostPreviewDTO, error) {
	userFollowing, err := gateway.GetAllUserFollowing(context.Background(), userRequestedId)
	if err != nil {
		fmt.Println(err)
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


			post.Description = post.Description + "\n\n" + appendToTags + "\n\n" + appendToDescHashtags

			postsPreview = append(postsPreview, dto.NewPostPreviewDTO(post))

		}
	}

	return postsPreview, nil
}

func (p postUseCase) EncodeBase64(media string, userId string, ctx context.Context) (string, error) {

	workingDirectory, _ := os.Getwd()
	path1 := "./assets/images/"
	err := os.Chdir(path1)
	if err != nil {
		fmt.Println(err)
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
		panic(err)
	}
	uuid := uuid.NewString()
	f, err := os.Create(uuid + "." + format[0])

	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	os.Chdir(workingDirectory)
	return userId + "/" + uuid + "." + format[0], nil
}


func (p postUseCase) AddPost(postDTO dto.CreatePostDTO, ctx context.Context) error {
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
	return p.postRepository.DeletePost(postDTO, context.Background())
}

func (p postUseCase) EditPost(postDTO dto.UpdatePostDTO, ctx context.Context) error {
	return p.postRepository.EditPost(postDTO, context.Background())
}

func (p postUseCase) GetPostsByUser(userId string, userRequestedId string, ctx context.Context) ([]dto.PostDTO, error) {
	posts, err := p.postRepository.GetPostsByUserId(userId, context.Background())
	var retVal []dto.PostDTO

	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		converted, err := p.GetPost(post.Id, post.Profile.Id, userRequestedId, context.Background())
		if err != nil {
			continue
		}
		retVal = append(retVal, converted)
	}

	return retVal, nil
}

func (p postUseCase) GetPost(postId string, userId string, userRequestedId string, ctx context.Context) (dto.PostDTO, error) {
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
	for _, s := range post.Hashtags {
		appendToDescHashtags = appendToDescHashtags + "#" + s
	}

	for _, s := range post.Mentions {
		appendToTags = appendToTags + "@" + s
	}

	post.Description = post.Description + "\n\n" + appendToTags + "\n\n" + appendToDescHashtags
	post.Media = mediaToAppend

	return post, nil

}

func NewPostUseCase(postRepository repository.PostRepo, repo repository.LikeRepo, favoritesRepo repository.FavoritesRepo, collectionRepo repository.CollectionRepo) PostUseCase {
	return &postUseCase{
		postRepository: postRepository,
		likeRepository: repo,
		favoriteRepository: favoritesRepo,
		collectionRepository: collectionRepo,
	}
}
