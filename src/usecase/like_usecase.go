package usecase

import (
	"post-service/domain"
	"post-service/dto"
	"post-service/repository"
)

type LikeUseCase interface {
	LikePost(dto dto.LikeDTO) error
	DislikePost(postId uint, profile *domain.Profile) error
	RemoveDislike(postId uint, profile *domain.Profile) error
	GetLikesForPost(postId uint) ([]domain.Like, error)
	GetDislikesForPost(postId uint) ([]domain.Dislike, error)
	GetNumOfLikesForPost(postId uint) (uint64, error)
	GetNumOfDislikesForPost(postId uint) (uint64, error)
}

type likeUseCase struct {
	likeRepository repository.LikeRepo
}

func (l likeUseCase) LikePost(dto dto.LikeDTO) error {
	panic("implement me")
}

func (l likeUseCase) DislikePost(postId uint, profile *domain.Profile) error {
	panic("implement me")
}

func (l likeUseCase) RemoveDislike(postId uint, profile *domain.Profile) error {
	panic("implement me")
}

func (l likeUseCase) GetLikesForPost(postId uint) ([]domain.Like, error) {
	panic("implement me")
}

func (l likeUseCase) GetDislikesForPost(postId uint) ([]domain.Dislike, error) {
	panic("implement me")
}

func (l likeUseCase) GetNumOfLikesForPost(postId uint) (uint64, error) {
	panic("implement me")
}

func (l likeUseCase) GetNumOfDislikesForPost(postId uint) (uint64, error) {
	panic("implement me")
}

func NewLikeUseCase(likeRepository repository.LikeRepo) LikeUseCase {
	return &likeUseCase{likeRepository: likeRepository}
}