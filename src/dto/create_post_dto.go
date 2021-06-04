package dto

import "post-service/domain"

type CreatePostDTO struct {
	UserId string `json:"user_id" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location domain.Location `json:"location" validate:"required"`
	Hashtags []string `json:"hashtags" validate:"required"`
	Mentions []string `json:"mentions" validate:"required"`
	Media []string `json:"media" validate:"required"`
	MediaType string `json:"media_type" validate:"required"`
}
