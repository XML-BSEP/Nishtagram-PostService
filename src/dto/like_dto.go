package dto

type LikeDislikeDTO struct {
	UserId string `json:"user_id" validate:"required"`
	PostId string `json:"post_id" validate:"required"`
	PostBy string `json:"post_by" validate:"required"`
}
