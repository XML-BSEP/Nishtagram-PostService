package dto

type DeletePostDTO struct {
	UserId string `json:"user_id" validate:"required"`
	PostId string `json:"user_id" validate:"required"`
}
