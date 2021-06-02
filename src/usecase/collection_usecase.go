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
	DeleteCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error

}

type collectionUseCase struct {
	collectionRepository repository.CollectionRepo
}

func (c collectionUseCase) DeleteCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error {
	return c.collectionRepository.DeleteCollection(collectionDTO.UserId, collectionDTO.CollectionName, context.Background())
}

func (c collectionUseCase) CreateCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error {
	return c.collectionRepository.CreateCollection(collectionDTO.UserId, collectionDTO.CollectionName, context.Background())
}

func (c collectionUseCase) AddPostToCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error {
	return c.collectionRepository.AddPostToCollection(collectionDTO.UserId, collectionDTO.CollectionName, collectionDTO.PostId, context.Background())
}

func (c collectionUseCase) RemovePostFromCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error {
	return c.collectionRepository.RemovePostFromCollection(collectionDTO.UserId, collectionDTO.CollectionName, collectionDTO.PostId, context.Background())
}

func NewCollectionUseCase(collectionRepository repository.CollectionRepo) CollectionUseCase {
	return &collectionUseCase{collectionRepository: collectionRepository}
}
