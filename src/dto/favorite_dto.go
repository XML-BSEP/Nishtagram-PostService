package dto

type FavoriteDTO struct {
	PostId string `json:"postid" validate:"required"`
	UserId string `json:"userid" validate:"required"`
	PostBy string `json:"postby" validate:"required"`


}
