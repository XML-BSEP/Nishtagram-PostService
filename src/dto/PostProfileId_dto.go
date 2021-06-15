package dto

type PostProfileId struct {
	PostId string `bson:"post_id" json:"post_id"`
	ProfileId string `bson:"profile_id" json:"profile_id"`
}