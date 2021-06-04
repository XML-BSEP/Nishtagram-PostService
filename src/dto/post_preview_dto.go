package dto

type PostPreviewDTO struct {
	Media []string
	Type string
	UserName string
	UserSurname string
	UserUsername string
}

func NewPostPreviewDTO(post PostDTO) PostPreviewDTO {
	return PostPreviewDTO{
		Media: post.Media,
		Type: post.MediaType.Type,
		UserName: "",
		UserUsername: "",
		UserSurname: "",
	}
}