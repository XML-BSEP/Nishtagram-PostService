package domain

import "time"

type Favorites struct {
	Id uint
	Profile Profile
	TimeOfCreation time.Time
	Posts []Post
}
