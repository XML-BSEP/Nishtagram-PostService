package dto


type PostTagProfileDTO struct {
	PostId	string `json:"post_id"`
	Hashtag []string `json:"hashtag"`
	ProfileId string `json:"profile_id"`
}

