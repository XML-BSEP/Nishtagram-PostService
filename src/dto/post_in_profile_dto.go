package dto

type PostInDTO struct {
	User string `json:"user" validate:"required"`
	Posts string `json:"images" validate:"required"`
	PostId string `json:"postid" validate:"required"`
	IsVideo bool `json:"isVideo" validate:"required"`
	PostBy string `json:"postby"`
	NotFollowing bool `json:"notFollowing"`
}
