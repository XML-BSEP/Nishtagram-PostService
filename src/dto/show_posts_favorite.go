package dto

type ShowFavoritePostsDTO struct {
	Posts []PostPreviewDTO
	UserId string
}

func NewShowFavoritePostsDTO(userId string, posts []PostPreviewDTO) ShowFavoritePostsDTO {
	return ShowFavoritePostsDTO{UserId: userId, Posts: posts}
}

func NewShowFavoriteNoParamsDTO() ShowFavoritePostsDTO {
	return ShowFavoritePostsDTO{}
}