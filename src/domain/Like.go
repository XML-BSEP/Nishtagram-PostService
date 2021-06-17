package domain

import "time"

type Like struct {
	PostId string
	Profile Profile
	Timestamp time.Time
	PostBy Profile
}

func NewLike(postId string, profileId string, timestamp time.Time) Like {
	return Like{
		PostId: postId,
		Profile: Profile{Id: profileId},
		Timestamp: timestamp,
	}
}
