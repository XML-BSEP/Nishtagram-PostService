package dto

import (
	"post-service/domain"
	"time"
)

type CollectionPreview struct {
	Id string `json:"postid" validate:"required"`
	Media string `json:"image" validate:"required"`
	Type string
	PostBy string `json:"postBy" validate:"required"`
	UserName string
	UserSurname string
	UserUsername string
	User domain.Profile `json:"user" validate:"required"`
	Location string `json:"location" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsAlbum bool `json:"isAlbum" validate:"required"`
	Timestamp time.Time `json:"time" validate:"required"`
	NumOfLikes int `json:"numOfLikes" validate:"required"`
	NumOfDislikes int `json:"numOfDislikes" validate:"required"`
	NumOfComments int `json:"numOfComments" validate:"required"`
	Banned bool
	IsVideo bool `json:"isVideo" validate:"required"`
	//Profile domain.Profile `json:"user" validate:"required"`
	IsBookmarked bool `json:"isBookmarked" validate:"required"`
	IsDisliked bool `json:"isDisliked" validate:"required"`
	IsLiked bool `json:"isLiked" validate:"required"`
}