package usecase

import (
	"context"
	"post-service/domain"
	"post-service/dto"
	"post-service/repository"
)

type LikeUseCase interface {
	LikePost(dto dto.LikeDTO, ctx context.Context) error
	DislikePost(postId string, profile *domain.Profile, ctx context.Context) error
	RemoveDislike(postId string, profile *domain.Profile, ctx context.Context) error
	GetLikesForPost(postId string, ctx context.Context) ([]domain.Like, error)
	GetDislikesForPost(postId string, ctx context.Context) ([]domain.Dislike, error)
	GetNumOfLikesForPost(postId string, ctx context.Context) (uint64, error)
	GetNumOfDislikesForPost(postId string, ctx context.Context) (uint64, error)
}

type likeUseCase struct {
	likeRepository repository.LikeRepo
}

func (l likeUseCase) LikePost(dto dto.LikeDTO, ctx context.Context) error {
	panic("implement me")
}

func (l likeUseCase) DislikePost(postId string, profile *domain.Profile, ctx context.Context) error {
	panic("implement me")
}

func (l likeUseCase) RemoveDislike(postId string, profile *domain.Profile, ctx context.Context) error {
	panic("implement me")
}

func (l likeUseCase) GetLikesForPost(postId string, ctx context.Context) ([]domain.Like, error) {
	panic("implement me")
}

func (l likeUseCase) GetDislikesForPost(postId string, ctx context.Context) ([]domain.Dislike, error) {
	panic("implement me")
}

func (l likeUseCase) GetNumOfLikesForPost(postId string, ctx context.Context) (uint64, error) {
	panic("implement me")
}

func (l likeUseCase) GetNumOfDislikesForPost(postId string, ctx context.Context) (uint64, error) {
	panic("implement me")
}

func NewLikeUseCase(likeRepository repository.LikeRepo) LikeUseCase {
	return &likeUseCase{likeRepository: likeRepository}
}