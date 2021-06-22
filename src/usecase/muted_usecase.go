package usecase

import (
	"context"
	"post-service/repository"
)

type MutedContentUseCase interface {
	SeeIfMuted(blockedFor string, blocked string, ctx context.Context) bool
	MuteUser(blockedFor string, blocked string, ctx context.Context) error
	UnmuteUser(blockedFor string, blocked string, ctx context.Context) error
}


type mutedContentUseCase struct {
	mutedContentRepo repository.MutedContentRepo
}

func (m mutedContentUseCase) SeeIfMuted(blockedFor string, blocked string, ctx context.Context) bool {
	return m.mutedContentRepo.SeeIfMuted(blockedFor, blocked, ctx)
}

func (m mutedContentUseCase) MuteUser(blockedFor string, blocked string, ctx context.Context) error {
	return m.mutedContentRepo.AddMuted(blockedFor, blocked, ctx)
}

func (m mutedContentUseCase) UnmuteUser(blockedFor string, blocked string, ctx context.Context) error {
	return m.mutedContentRepo.DeleteFrom(blockedFor, blocked, ctx)
}

func NewMutedContentUseCase(mutedContentRepo repository.MutedContentRepo) MutedContentUseCase {
	return &mutedContentUseCase{mutedContentRepo: mutedContentRepo}
}