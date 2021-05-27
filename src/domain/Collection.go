package domain

import "time"

type Collection struct {
	Id uint
	Name string
	Timestamp time.Time
	Profile Profile
	Posts []Post
}
