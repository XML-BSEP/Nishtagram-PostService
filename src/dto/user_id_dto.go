package dto

type UserIdDTO struct {
	UserId string `json:"user_id" validate:"required"`
}
