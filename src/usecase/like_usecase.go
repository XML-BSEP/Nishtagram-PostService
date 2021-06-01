package usecase

import (
	"context"
	"post-service/domain"
	"post-service/dto"
	"post-service/repository"
)

type LikeUseCase interface {
	LikePost(dto dto.LikeDTO, ctx context.Context) error
	DislikePost(postId uint, profile *domain.Profile, ctx context.Context) error
	RemoveDislike(postId uint, profile *domain.Profile, ctx context.Context) error
	GetLikesForPost(postId uint, ctx context.Context) ([]domain.Like, error)
	GetDislikesForPost(postId uint, ctx context.Context) ([]domain.Dislike, error)
	GetNumOfLikesForPost(postId uint, ctx context.Context) (uint64, error)
	GetNumOfDislikesForPost(postId uint, ctx context.Context) (uint64, error)
}

type likeUseCase struct {
	likeRepository repository.LikeRepo
}

func (l likeUseCase) LikePost(dto dto.LikeDTO, ctx context.Context) error {
	panic("implement me")
}

func (l likeUseCase) DislikePost(postId uint, profile *domain.Profile, ctx context.Context) error {
	panic("implement me")
}

func (l likeUseCase) RemoveDislike(postId uint, profile *domain.Profile, ctx context.Context) error {
	panic("implement me")
}

func (l likeUseCase) GetLikesForPost(postId uint, ctx context.Context) ([]domain.Like, error) {
	panic("implement me")
}

func (l likeUseCase) GetDislikesForPost(postId uint, ctx context.Context) ([]domain.Dislike, error) {
	panic("implement me")
}

func (l likeUseCase) GetNumOfLikesForPost(postId uint, ctx context.Context) (uint64, error) {
	panic("implement me")
}

func (l likeUseCase) GetNumOfDislikesForPost(postId uint, ctx context.Context) (uint64, error) {
	panic("implement me")
}

func NewLikeUseCase(likeRepository repository.LikeRepo) LikeUseCase {
	return &likeUseCase{likeRepository: likeRepository}
}