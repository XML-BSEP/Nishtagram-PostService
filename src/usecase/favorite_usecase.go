package usecase

import (
	"context"
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
}

func (f favoriteUseCase) GetFavoritesForUser(userId string, ctx context.Context) ([]dto.PostInDTO, error) {
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
	}


	return retVal, nil

}

func (f favoriteUseCase) AddPostToFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error {
	return f.favoriteRepository.AddPostToFavorites(favoriteDTO.PostId, favoriteDTO.UserId, favoriteDTO.PostBy, context.Background())
}

func (f favoriteUseCase) RemovePostFromFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error {
	return f.favoriteRepository.RemovePostFromFavorites(favoriteDTO.PostId, favoriteDTO.UserId, context.Background())
}

func NewFavoriteUseCase(favoritesRepository repository.FavoritesRepo, postRepository repository.PostRepo, useCase PostUseCase) FavoriteUseCase {
	return &favoriteUseCase{favoriteRepository: favoritesRepository, postRepository: postRepository, postUseCase: useCase}
}