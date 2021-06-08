package dto


type CollectionDTO struct {
	UserId string `json:"id" validate:"required"`
	PostId string `json:"postid" validate:"required"`
	CollectionName string `json:"name" validate:"required"`
	PostBy string `json:"postby"`

}
