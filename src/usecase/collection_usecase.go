package usecase

import (
	"context"
	"post-service/dto"
	"post-service/repository"
)

type CollectionUseCase interface {
	CreateCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error
	AddPostToCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error
	RemovePostFromCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error
}

type collectionUseCase struct {
	collectionRepository repository.CollectionRepo
}

func (c collectionUseCase) CreateCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error {
	panic("implement me")
}

func (c collectionUseCase) AddPostToCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error {
	panic("implement me")
}

func (c collectionUseCase) RemovePostFromCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error {
	panic("implement me")
}

func NewCollectionUseCase(collectionRepository repository.CollectionRepo) CollectionUseCase {
	return &collectionUseCase{collectionRepository: collectionRepository}
}
