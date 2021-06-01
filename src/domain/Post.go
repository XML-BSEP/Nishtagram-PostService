package domain

import "time"

type Post struct {
	Id string
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
