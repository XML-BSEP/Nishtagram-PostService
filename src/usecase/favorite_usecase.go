package usecase

import (
	"context"
	"post-service/dto"
	"post-service/repository"
)

type FavoriteUseCase interface {
	AddPostToFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error
	RemovePostFromFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error
	GetFavoritesForUser(userId string, ctx context.Context) (dto.ShowFavoritePostsDTO, error)
}

type favoriteUseCase struct {
	favoriteRepository repository.FavoritesRepo
	postRepository repository.PostRepo
}

func (f favoriteUseCase) GetFavoritesForUser(userId string, ctx context.Context) (dto.ShowFavoritePostsDTO, error) {
	favorites, err := f.favoriteRepository.GetFavorites(userId)
	if err != nil {
		return dto.NewShowFavoriteNoParamsDTO(), err
	}

	var posts []dto.PostDTO
	var bannedPosts []string

	for favorite := range favorites {
		if f.postRepository.SeeIfPostDeletedOrBanned(favorite, favorites[favorite], context.Background()) {
			bannedPosts = append(bannedPosts, favorite)
			continue
		}
		post, err := f.postRepository.GetPostsById(favorites[favorite], favorite, context.Background())
		if err != nil {
			continue
		}
		posts = append(posts, post)
	}

	for _, s := range bannedPosts {
		err = f.favoriteRepository.RemovePostFromFavorites(s, favorites[s], context.Background())
	}

	var postsPreview []dto.PostPreviewDTO
	for _, post := range posts {
		postsPreview = append(postsPreview, dto.NewPostPreviewDTO(post))
	}

	return dto.NewShowFavoritePostsDTO(userId, postsPreview), nil

}

func (f favoriteUseCase) AddPostToFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error {
	return f.favoriteRepository.AddPostToFavorites(favoriteDTO.PostId, favoriteDTO.UserId, favoriteDTO.PostBy, context.Background())
}

func (f favoriteUseCase) RemovePostFromFavorites(favoriteDTO dto.FavoriteDTO, ctx context.Context) error {
	return f.favoriteRepository.RemovePostFromFavorites(favoriteDTO.PostId, favoriteDTO.UserId, context.Background())
}

func NewFavoriteUseCase(favoritesRepository repository.FavoritesRepo, postRepository repository.PostRepo) FavoriteUseCase {
	return &favoriteUseCase{favoriteRepository: favoritesRepository, postRepository: postRepository}
}