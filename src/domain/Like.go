package domain

import "time"

type Like struct {
	Id string
	PostId string
	Profile Profile
	Timestamp time.Time
}
