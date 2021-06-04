package dto

import "post-service/domain"

type UpdatePostDTO struct {
	UserId string `json:"user_id" validate:"required"`
	PostId string `json:"post_id" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location domain.Location `json:"location" validate:"required"`
	Hashtags []string `json:"hashtags" validate:"required"`
	Mentions []string `json:"mentions" validate:"required"`
}