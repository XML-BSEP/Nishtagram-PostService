package usecase

import (
	"context"
	"post-service/dto"
	"post-service/repository"
)

type FavoriteUseCase interface {
	AddPostToFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error
	RemovePostFromFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error
}

type favoriteUseCase struct {
	favoriteRepository repository.FavoritesRepo
}

func (f favoriteUseCase) AddPostToFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error {
	return f.favoriteRepository.AddPostToFavorites(favoriteDTO.PostId, favoriteDTO.UserId, favoriteDTO.PostBy, context.Background())
}

func (f favoriteUseCase) RemovePostFromFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error {
	return f.favoriteRepository.RemovePostFromFavorites(favoriteDTO.PostId, favoriteDTO.UserId, context.Background())
}

func NewFavoriteUseCase(favoritesRepository repository.FavoritesRepo) FavoriteUseCase {
	return &favoriteUseCase{favoriteRepository: favoritesRepository}
}