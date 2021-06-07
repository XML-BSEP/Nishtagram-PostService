package dto

type PostInDTO struct {
	User string `json:"user" validate:"required"`
	Posts string `json:"image" validate:"required"`
	PostId string `json:"postid" validate:"required"`
	IsVideo bool `json:"isVideo" validate:"required"`
}
