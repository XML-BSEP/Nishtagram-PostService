package domain

import "time"

type Like struct {
	Id uint
	PostId uint
	Profile Profile
	Timestamp time.Time
}
