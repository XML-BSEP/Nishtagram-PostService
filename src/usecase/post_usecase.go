package usecase

import (
	"post-service/dto"
	"post-service/repository"
)

type PostUseCase interface {
	AddPost(postDTO dto.PostDTO)
	DeletePost(postDTO dto.PostDTO)
	EditPost(postDTO dto.PostDTO)
	GetPostsForUserFeed(userId uint)
	GetPost(postId uint, userId uint)
}

type postUseCase struct {
	postRepository repository.PostRepo
}

func (p postUseCase) AddPost(postDTO dto.PostDTO) {
	panic("implement me")
}

func (p postUseCase) DeletePost(postDTO dto.PostDTO) {
	panic("implement me")
}

func (p postUseCase) EditPost(postDTO dto.PostDTO) {
	panic("implement me")
}

func (p postUseCase) GetPostsForUserFeed(userId uint) {
	panic("implement me")
}

func (p postUseCase) GetPost(postId uint, userId uint) {
	panic("implement me")
}

func NewPostUseCase(postRepository repository.PostRepo) PostUseCase {
	return &postUseCase{postRepository: postRepository}
}
