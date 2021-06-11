package usecase

import (
	"context"
	logger "github.com/jelena-vlajkov/logger/logger"
	"post-service/dto"
	"post-service/repository"
)

type CollectionUseCase interface {
	CreateCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error
	AddPostToCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error
	RemovePostFromCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error
	DeleteCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error
	GetCollection(userId string, collectionName string, ctx context.Context) (dto.PreviewCollectionDTO, error)
	GetAllCollectionsPerUser(userId string, ctx context.Context) ([]dto.CollectionDTO, error)

}

type collectionUseCase struct {
	collectionRepository repository.CollectionRepo
	postRepository repository.PostRepo
	postUseCase PostUseCase
	logger *logger.Logger
}

func (c collectionUseCase) GetAllCollectionsPerUser(userId string, ctx context.Context) ([]dto.CollectionDTO, error) {
	c.logger.Logger.Infof("getting all collection for user %v\n", userId)
	var retVal []dto.CollectionDTO

	collections, err :=  c.collectionRepository.GetAllCollectionNames(userId, context.Background())

	if err != nil {
		return nil, err
	}

	for _, col := range collections {
		retVal = append(retVal, dto.CollectionDTO{CollectionName: col})
	}
	return retVal, nil
}

func (c collectionUseCase) GetCollection(userId string, collectionName string, ctx context.Context) (dto.PreviewCollectionDTO, error) {
	c.logger.Logger.Infof("getting collection with name %v for user %v\n", collectionName, userId)
	posts, err := c.collectionRepository.GetCollection(userId, collectionName, context.Background())
	if err != nil {
		return dto.NewPreviewCollectionDTO(), err
	}
	var bannedPosts []string
	var postsPreview []dto.PostPreviewDTO
	var retVal []dto.PostInDTO

	for s := range posts {
		if c.postRepository.SeeIfPostDeletedOrBanned(s, posts[s], context.Background()) {
			bannedPosts = append(bannedPosts, s)
			continue
		}
		
		post, err := c.postUseCase.GetPost(s, posts[s], userId, context.Background())
		post.PostBy = posts[s]
		post.Id = s

		if err != nil {
			continue
		}
		postsPreview = append(postsPreview, post)
		isVideo := false
		if post.Type == "VIDEO" {
			isVideo = true
		}
		retVal = append(retVal, dto.PostInDTO{PostId: s, Posts: post.Media[0], PostBy: posts[s], IsVideo: isVideo, User: posts[s]})
	}

	for _, s := range bannedPosts {
		c.collectionRepository.RemovePostFromCollection(userId, collectionName, s, context.Background())
	}

	return dto.NewPreviewCollectionParDTO(collectionName, userId, retVal), nil
}

func (c collectionUseCase) DeleteCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error {
	c.logger.Logger.Infof("deleting collection with name %v for user %v\n", collectionDTO.CollectionName, collectionDTO.UserId)
	return c.collectionRepository.DeleteCollection(collectionDTO.UserId, collectionDTO.CollectionName, context.Background())
}

func (c collectionUseCase) CreateCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error {
	c.logger.Logger.Infof("creating collection with name %v for user %v\n", collectionDTO.CollectionName, collectionDTO.UserId)
	return c.collectionRepository.CreateCollection(collectionDTO.UserId, collectionDTO.CollectionName, context.Background())
}

func (c collectionUseCase) AddPostToCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error {
	c.logger.Logger.Infof("adding post %v to collection with name %v for user %v\n", collectionDTO.PostId, collectionDTO.CollectionName, collectionDTO.UserId)
	return c.collectionRepository.AddPostToCollection(collectionDTO.UserId, collectionDTO.CollectionName, collectionDTO.PostId, collectionDTO.PostBy, context.Background())
}

func (c collectionUseCase) RemovePostFromCollection(collectionDTO dto.CollectionDTO, ctx context.Context) error {
	c.logger.Logger.Infof("removing post %v from collection with name %v for user %v\n", collectionDTO.PostId, collectionDTO.CollectionName, collectionDTO.UserId)
	return c.collectionRepository.RemovePostFromCollection(collectionDTO.UserId, collectionDTO.CollectionName, collectionDTO.PostId, context.Background())
}

func NewCollectionUseCase(collectionRepository repository.CollectionRepo, postRepository repository.PostRepo, useCase PostUseCase, logger *logger.Logger ) CollectionUseCase {
	return &collectionUseCase{collectionRepository: collectionRepository, postRepository: postRepository, postUseCase: useCase, logger: logger}
}
