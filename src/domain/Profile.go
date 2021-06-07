package domain

type Profile struct {
	Id string `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	ProfilePhoto string `json:"profilePhoto" validate:"required"`
}
