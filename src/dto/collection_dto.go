package dto


type CollectionDTO struct {
	UserId string `json:"user_id" validate:"required"`
	PostId string `json:"post_id" validate:"required"`
	CollectionName string `json:"collection_name" validate:"required"`

}
