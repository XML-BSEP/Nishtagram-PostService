package domain

import "time"

type Collection struct {
	Id string
	Name string
	Timestamp time.Time
	Profile Profile
	Posts []Post
}
