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
	Media []string
	MediaType MediaType
}

func NewPost(id string, desc string, timestamp time.Time, numOfLikes int, numOfDislikes int,
	numOfComments int, profileId string, locationName string, mentions []string, hashtags []string, media []string, mediaType string) Post {

	var mentionsSlice []Profile
	var hashtagsSlice []Hashtag

	for _, s := range mentions {
		mentionsSlice = append(mentionsSlice, Profile{Id: s})
	}

	for _, s := range hashtags {
		hashtagsSlice = append(hashtagsSlice, Hashtag{Tag: s})
	}

	return Post{
		Id: id,
		Description: desc,
		Timestamp: timestamp,
		NumOfLikes: numOfLikes,
		NumOfDislikes: numOfDislikes,
		NumOfComments: numOfComments,
		Profile: Profile{Id: profileId},
		Location: Location{Location: locationName},
		MediaType: MediaType{Type: mediaType},
		Hashtags: hashtagsSlice,
		Tags: mentionsSlice,
		Media: media,
	}
}
