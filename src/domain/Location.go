package domain

type Location struct {
	Location string `json:"location" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Latitude float64 `json:"latitude" validate:"required"`
}
