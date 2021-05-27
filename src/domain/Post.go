package domain

import "time"

type Post struct {
	Id uint
	Description string
	Timestamp time.Time
	NumOfLikes int
	NumOfDislikes int
	NumOfComments int
	Banned bool
	Profile Profile
	Location Location
	Hashtags []Hashtag
	Tags []Profile
	Media []Media
	MediaType MediaType

}
