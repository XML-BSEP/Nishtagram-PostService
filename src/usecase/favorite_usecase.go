package usecase

import (
	"context"
	logger "github.com/jelena-vlajkov/logger/logger"
	"post-service/dto"
	"post-service/repository"
)

type FavoriteUseCase interface {
	AddPostToFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error
	RemovePostFromFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error
	GetFavoritesForUser(userId string, ctx context.Context) ([]dto.PostInDTO, error)
}

type favoriteUseCase struct {
	favoriteRepository repository.FavoritesRepo
	postRepository repository.PostRepo
	postUseCase PostUseCase
	logger *logger.Logger

}

func (f favoriteUseCase) GetFavoritesForUser(userId string, ctx context.Context) ([]dto.PostInDTO, error) {
	f.logger.Logger.Infof("getting favorite posts for user %v\n", userId)
	favorites, err := f.favoriteRepository.GetFavorites(userId)
	if err != nil {
		return nil, err
	}

	var posts string
	var bannedPosts []string
	var retVal []dto.PostInDTO
	for favorite := range favorites {
		if f.postRepository.SeeIfPostDeletedOrBanned(favorite, favorites[favorite], context.Background()) {
			bannedPosts = append(bannedPosts, favorite)
			continue
		}
		post, err := f.postUseCase.GetPost( favorite, favorites[favorite], userId, context.Background())
		if err != nil {
			continue
		}
		posts = post.Media[0]
		retVal = append(retVal, dto.PostInDTO{User: favorites[favorite], Posts: posts, PostBy: favorites[favorite], PostId: favorite})
	}

	for _, s := range bannedPosts {
		err = f.favoriteRepository.RemovePostFromFavorites(s, favorites[s], context.Background())
		f.logger.Logger.Errorf("error while deleting banned posts from favorites for %v, for post id %v, error: %v\n", userId, s, err)
	}

	return retVal, nil

}

func (f favoriteUseCase) AddPostToFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error {
	f.logger.Logger.Infof("adding post %v to favorites for user %v\n", favoriteDTO.PostId, favoriteDTO.UserId)
	return f.favoriteRepository.AddPostToFavorites(favoriteDTO.PostId, favoriteDTO.UserId, favoriteDTO.PostBy, context.Background())
}

func (f favoriteUseCase) RemovePostFromFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error {
	f.logger.Logger.Infof("removing post %v from favorites for user %v\n", favoriteDTO.PostId, favoriteDTO.UserId)
	return f.favoriteRepository.RemovePostFromFavorites(favoriteDTO.PostId, favoriteDTO.UserId, context.Background())

}

func NewFavoriteUseCase(favoritesRepository repository.FavoritesRepo, postRepository repository.PostRepo, useCase PostUseCase, logger *logger.Logger) FavoriteUseCase {
	return &favoriteUseCase{favoriteRepository: favoritesRepository, postRepository: postRepository, postUseCase: useCase, logger: logger}
}