package usecase

import (
	"post-service/dto"
	"post-service/repository"
)

type CollectionUseCase interface {
	CreateCollection(collectionDTO dto.CollectionDTO) error
	AddPostToCollection(collectionDTO dto.CollectionDTO) error
	RemovePostFromCollection(collectionDTO dto.CollectionDTO) error
}

type collectionUseCase struct {
	collectionRepository repository.CollectionRepo
}

func (c collectionUseCase) CreateCollection(collectionDTO dto.CollectionDTO) error {
	panic("implement me")
}

func (c collectionUseCase) AddPostToCollection(collectionDTO dto.CollectionDTO) error {
	panic("implement me")
}

func (c collectionUseCase) RemovePostFromCollection(collectionDTO dto.CollectionDTO) error {
	panic("implement me")
}

func NewCollectionUseCase(collectionRepository repository.CollectionRepo) CollectionUseCase {
	return &collectionUseCase{collectionRepository: collectionRepository}
}
