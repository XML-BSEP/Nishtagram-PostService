package dto

type UserIdDTO struct {
	Id string `json:"id" validate:"required"`
}
