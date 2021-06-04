package dto

import (
	"post-service/domain"
	"time"
)

type PostDTO struct {
	Id string
	Description string
	Timestamp time.Time
	NumOfLikes int
	NumOfDislikes int
	NumOfComments int
	Banned bool
	Profile domain.Profile
	Location domain.Location
	Hashtags []domain.Hashtag
	Tags []domain.Profile
	Media []string
	MediaType domain.MediaType
}

func NewPost(id string, desc string, timestamp time.Time, numOfLikes int, numOfDislikes int,
	numOfComments int, profileId string, locationName string, mentions []string, hashtags []string, media []string, mediaType string) PostDTO {

	var mentionsSlice []domain.Profile
	var hashtagsSlice []domain.Hashtag

	for _, s := range mentions {
		mentionsSlice = append(mentionsSlice, domain.Profile{Id: s})
	}

	for _, s := range hashtags {
		hashtagsSlice = append(hashtagsSlice, domain.Hashtag{Tag: s})
	}

	return PostDTO{
		Id: id,
		Description: desc,
		Timestamp: timestamp,
		NumOfLikes: numOfLikes,
		NumOfDislikes: numOfDislikes,
		NumOfComments: numOfComments,
		Profile: domain.Profile{Id: profileId},
		Location: domain.Location{Location: locationName},
		MediaType: domain.MediaType{Type: mediaType},
		Hashtags: hashtagsSlice,
		Tags: mentionsSlice,
		Media: media,
	}
}
