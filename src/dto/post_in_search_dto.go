package dto
/*id, media, type*/
type PostSearchDTO struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Media []string `json:"image"`
}

func NewPostSearchDTO(id string, types string, media []string) PostSearchDTO {
	return PostSearchDTO{Id: id, Type: types, Media: media}
}
