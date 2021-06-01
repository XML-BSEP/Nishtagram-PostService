package domain

import "time"

type Dislike struct {
	Id string
	PostId uint
	Profile Profile
	Timestamp time.Time
}
