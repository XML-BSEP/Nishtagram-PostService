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
	GetCollection(userId string, collectionName string, ctx context.Context) (dto.PreviewCollectionDTO, error)
	GetAllCollectionsPerUser(userId string, ctx context.Context) ([]string, error)

}

type collectionUseCase struct {
	collectionRepository repository.CollectionRepo
	postRepository repository.PostRepo
}

func (c collectionUseCase) GetAllCollectionsPerUser(userId string, ctx context.Context) ([]string, error) {
	return c.collectionRepository.GetAllCollectionNames(userId, context.Background())
}

func (c collectionUseCase) GetCollection(userId string, collectionName string, ctx context.Context) (dto.PreviewCollectionDTO, error) {
	posts, err := c.collectionRepository.GetCollection(userId, collectionName, context.Background())
	if err != nil {
		return dto.NewPreviewCollectionDTO(), err
	}
	var bannedPosts []string
	var postsPreview []dto.PostPreviewDTO

	for _, s := range posts {
		if c.postRepository.SeeIfPostDeletedOrBanned(userId, s, context.Background()) {
			bannedPosts = append(bannedPosts, s)
			continue
		}
		
		post, err := c.postRepository.GetPostsById(userId, s, context.Background())

		if err != nil {
			continue
		}
		postsPreview = append(postsPreview, dto.NewPostPreviewDTO(post))
	}

	for _, s := range bannedPosts {
		c.collectionRepository.RemovePostFromCollection(userId, collectionName, s, context.Background())
	}

	return dto.NewPreviewCollectionParDTO(collectionName, userId, postsPreview), nil
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

func NewCollectionUseCase(collectionRepository repository.CollectionRepo, postRepository repository.PostRepo) CollectionUseCase {
	return &collectionUseCase{collectionRepository: collectionRepository, postRepository: postRepository}
}
