package usecase

import (
	"context"
	"post-service/domain"
	"post-service/dto"
	"post-service/repository"
)

type LikeUseCase interface {
	LikePost(dto dto.LikeDislikeDTO, ctx context.Context) error
	DislikePost(dto dto.LikeDislikeDTO, ctx context.Context) error
	RemoveLike(dto dto.LikeDislikeDTO, ctx context.Context) error
	RemoveDislike(dto dto.LikeDislikeDTO, ctx context.Context) error
	GetLikesForPost(postId string, ctx context.Context) ([]domain.Like, error)
	GetDislikesForPost(postId string, ctx context.Context) ([]domain.Dislike, error)
}

type likeUseCase struct {
	likeRepository repository.LikeRepo
}

func (l likeUseCase) RemoveLike(dto dto.LikeDislikeDTO, ctx context.Context) error {
	return l.likeRepository.RemoveLike(dto.PostBy, dto.PostBy, domain.Profile{Id: dto.UserId}, context.Background())
}

func (l likeUseCase) LikePost(dto dto.LikeDislikeDTO, ctx context.Context) error {
	return l.likeRepository.LikePost(dto.PostId, dto.PostBy, domain.Profile{Id: dto.UserId}, context.Background())
}

func (l likeUseCase) DislikePost(dto dto.LikeDislikeDTO, ctx context.Context) error {
	return l.likeRepository.LikePost(dto.PostId, dto.PostBy, domain.Profile{Id: dto.UserId}, context.Background())
}

func (l likeUseCase) RemoveDislike(dto dto.LikeDislikeDTO, ctx context.Context) error {
	return l.likeRepository.RemoveDislike(dto.PostBy, dto.PostBy, domain.Profile{Id: dto.UserId}, context.Background())
}

func (l likeUseCase) GetLikesForPost(postId string, ctx context.Context) ([]domain.Like, error) {
	return l.likeRepository.GetLikesForPost(postId, context.Background())
}

func (l likeUseCase) GetDislikesForPost(postId string, ctx context.Context) ([]domain.Dislike, error) {
	return l.likeRepository.GetDislikesForPost(postId, context.Background())
}


func NewLikeUseCase(likeRepository repository.LikeRepo) LikeUseCase {
	return &likeUseCase{likeRepository: likeRepository}
}