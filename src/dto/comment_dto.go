package dto

import "post-service/domain"

type CommentDTO struct {
	Comment string `json:"comment" validate:"required"`
	PostId string `json:"post_id" validate:"required"`
	CommentBy domain.Profile `json:"comment_by" validate:"required"`
	Mentions []domain.Profile `json:"mentions" validate:"required"`
}