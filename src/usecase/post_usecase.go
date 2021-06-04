package usecase

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"post-service/dto"
	"post-service/repository"
	"strconv"
	"strings"
)

type PostUseCase interface {
	AddPost(postDTO dto.CreatePostDTO, ctx context.Context) error
	DeletePost(postDTO dto.DeletePostDTO, ctx context.Context) error
	EditPost(postDTO dto.UpdatePostDTO, ctx context.Context) error
	GetPostsByUser(userId string, ctx context.Context) ([]dto.PostDTO, error)
	GetPost(postId string, userId string, ctx context.Context) (dto.PostDTO, error)
	EncodeBase64Images(images []string, userId string) ([]string, error)
}

type postUseCase struct {
	postRepository repository.PostRepo
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

func (p postUseCase) AddPost(postDTO dto.CreatePostDTO, ctx context.Context) error {
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
