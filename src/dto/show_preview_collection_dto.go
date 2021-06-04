package dto

type PreviewCollectionDTO struct {
	CollectionName string
	UserId string
	Posts []PostPreviewDTO

}

func NewPreviewCollectionDTO() PreviewCollectionDTO {
	return PreviewCollectionDTO{

	}
}
func NewPreviewCollectionParDTO(name string, user string, posts []PostPreviewDTO) PreviewCollectionDTO {
	return PreviewCollectionDTO{
		CollectionName: name,
		UserId: user,
		Posts: posts,
	}
}