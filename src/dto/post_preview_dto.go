package dto

import (
	"post-service/domain"
	"time"
)

type PostPreviewDTO struct {
	Id string `json:"id" validate:"required"`
	Media []string `json:"images" validate:"required"`
	Type string
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

func NewPostPreviewDTO(post PostDTO) PostPreviewDTO {
	return PostPreviewDTO{
		Id : post.Id,
		Media: post.Media,
		User: post.Profile,
		Location: post.Location,
		Description: post.Description,
		IsAlbum: post.IsAlbum,
		IsDisliked: post.IsDisliked,
		IsLiked: post.IsLiked,
		IsVideo: post.IsVideo,
		IsBookmarked: post.IsBookmarked,
		NumOfComments: post.NumOfComments,
		NumOfDislikes: post.NumOfDislikes,
		NumOfLikes: post.NumOfLikes,
		Banned: post.Banned,
		//Profile: post.Profile,
		Type: post.MediaType.Type,
		UserName: "",
		UserUsername: "",
		UserSurname: "",
	}
}