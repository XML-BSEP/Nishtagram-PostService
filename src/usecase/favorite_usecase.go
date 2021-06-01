package usecase

import (
	"post-service/dto"
	"post-service/repository"
)

type FavoriteUseCase interface {
	AddPostToFavorites(favoriteDTO dto.FavoriteDTO) error
	RemovePostFromFavorites(favoriteDTO dto.FavoriteDTO) error
}

type favoriteUseCase struct {
	favoriteRepository repository.FavoritesRepo
}

func (f favoriteUseCase) AddPostToFavorites(favoriteDTO dto.FavoriteDTO) error {
	panic("implement me")
}

func (f favoriteUseCase) RemovePostFromFavorites(favoriteDTO dto.FavoriteDTO) error {
	panic("implement me")
}

func NewFavoriteUseCase(favoritesRepository repository.FavoritesRepo) FavoriteUseCase {
	return &favoriteUseCase{favoriteRepository: favoritesRepository}
}