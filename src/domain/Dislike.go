package domain

import "time"

type Dislike struct {
	Id uint
	PostId uint
	Profile Profile
	Timestamp time.Time
}
