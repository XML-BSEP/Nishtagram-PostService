package usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"os"
	"post-service/dto"
	"post-service/gateway"
	"post-service/repository"
	"strconv"
	"strings"
	"time"
)

type PostUseCase interface {
	AddPost(postDTO dto.CreatePostDTO, ctx context.Context) error
	DeletePost(postDTO dto.DeletePostDTO, ctx context.Context) error
	EditPost(postDTO dto.UpdatePostDTO, ctx context.Context) error
	GetPostsByUser(userId string, ctx context.Context) ([]dto.PostDTO, error)
	GetPost(postId string, userId string, ctx context.Context) (dto.PostDTO, error)
	EncodeBase64Images(images []string, userId string) ([]string, error)
	GenerateUserFeed(userId string, ctx context.Context) ([]dto.PostPreviewDTO, error)
	EncodeBase64(media string, userId string, ctx context.Context) (string, error)
}

type postUseCase struct {
	postRepository repository.PostRepo
}

func (p postUseCase) GenerateUserFeed(userId string, ctx context.Context) ([]dto.PostPreviewDTO, error) {
	userFollowing, err := gateway.GetAllUserFollowing(context.Background(), userId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}


	inTimeRange := time.Now().Add(-72*time.Hour)
	postsToShow := make(map[string][]string, len(userFollowing))
	for _, user := range userFollowing {
		posts := p.postRepository.GetPostsInDateTimeRange(user.Id, inTimeRange, context.Background())
		postsToShow[user.Id] = append(postsToShow[user.Id], posts...)
	}

	var postsPreview []dto.PostPreviewDTO
	for idFollowing := range postsToShow {
		for _, postById := range postsToShow[idFollowing] {
			post, err := p.postRepository.GetPostsById(idFollowing, postById, context.Background())
			if err != nil {
				continue
			}
			postsPreview = append(postsPreview, dto.NewPostPreviewDTO(post))
		}
	}

	return postsPreview, nil
}

func (p postUseCase) EncodeBase64Images(images []string, userId string) ([]string, error) {
	path2, _ := os.Getwd()
	fmt.Println(path2)

	path1 := "./src/assets"
	os.Chdir(path1)

	os.Mkdir(userId, 0755)

	os.Chdir(userId)

	imagesToSave := make([]string, len(images))

	if len(images) > 0{
		for i,_ := range images {

			s := strings.Split(images[i], ",")
			a := strings.Split(s[0], "/")
			format := strings.Split(a[1], ";")

			dec, err := base64.StdEncoding.DecodeString(s[1])

			if err != nil {
				return nil, err
			}
			f, err := os.Create(strconv.Itoa(i) + "." + format[0])

			if err != nil {
				return nil, err
			}

			defer f.Close()

			if _, err := f.Write(dec); err != nil {
				return nil, err
			}
			if err := f.Sync(); err != nil {
				return nil, err
			}

			imagesToSave = append(imagesToSave, "/" + strconv.Itoa(i) + "." + format[0])
		}
	}


	os.Chdir(path2)
	return imagesToSave, nil
}

func (p postUseCase) EncodeBase64(media string, userId string, ctx context.Context) (string, error) {

	workingDirectory, _ := os.Getwd()
	path1 := "././assets/images"
	os.Chdir(path1)
	err := os.Mkdir("././assets/images" + userId, 0755)

	os.Chdir("././assets/images" + userId)


	s := strings.Split(media, ",")
	a := strings.Split(s[0], "/")
	format := strings.Split(a[1], ";")

	dec, err := base64.StdEncoding.DecodeString(s[1])

	if err != nil {
		panic(err)
	}
	uuid := uuid.NewString()
	f, err := os.Create(uuid + format[0])

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
	return path1 + uuid + format[0], nil
}


func (p postUseCase) AddPost(postDTO dto.CreatePostDTO, ctx context.Context) error {
	var media []string

	if postDTO.IsImage || postDTO.IsVideo {
		media = make([]string, 1)
		mediaToAttach, err := p.EncodeBase64(postDTO.Image, postDTO.UserId.UserId, context.Background())
		if err != nil {
			return fmt.Errorf("error while decoding base64")
		}
		media[0] = mediaToAttach
		postDTO.Media = media
	} else {
		media = make([]string, len(postDTO.Album))
		for _, s := range postDTO.Album {
			mediaToAttach, err := p.EncodeBase64(s, postDTO.UserId.UserId, context.Background())
			if err != nil {
				return fmt.Errorf("error while decoding base64")
			}
			media = append(media, mediaToAttach)
		}

		postDTO.Media = media
	}

	return p.postRepository.CreatePost(postDTO, context.Background())
}

func (p postUseCase) DeletePost(postDTO dto.DeletePostDTO, ctx context.Context) error {
	return p.postRepository.DeletePost(postDTO, context.Background())
}

func (p postUseCase) EditPost(postDTO dto.UpdatePostDTO, ctx context.Context) error {
	return p.postRepository.EditPost(postDTO, context.Background())
}

func (p postUseCase) GetPostsByUser(userId string, ctx context.Context) ([]dto.PostDTO, error) {
	return p.postRepository.GetPostsByUserId(userId, context.Background())
}

func (p postUseCase) GetPost(postId string, userId string, ctx context.Context) (dto.PostDTO, error) {
	return p.postRepository.GetPostsById(userId, postId, context.Background())
}

func NewPostUseCase(postRepository repository.PostRepo) PostUseCase {
	return &postUseCase{postRepository: postRepository}
}
