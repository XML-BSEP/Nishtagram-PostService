package dto

type PreviewCollectionDTO struct {
	CollectionName string `json:"name"`
	UserId string `json:"user"`
	Posts []PostInDTO `json:"posts"`
	PostBy string `json:"postby"`
	PostId string `json:"postid"`
}

func NewPreviewCollectionDTO() PreviewCollectionDTO {
	return PreviewCollectionDTO{

	}
}
func NewPreviewCollectionParDTO(name string, user string, posts []PostInDTO) PreviewCollectionDTO {
	return PreviewCollectionDTO{
		CollectionName: name,
		UserId: user,
		Posts: posts,
		PostBy: user,
	}
}