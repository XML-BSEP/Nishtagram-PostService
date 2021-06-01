package domain

import "time"

type Favorites struct {
	Id string
	Profile Profile
	TimeOfCreation time.Time
	Posts []Post
}
