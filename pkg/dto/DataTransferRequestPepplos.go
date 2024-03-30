package dto

type DTORequestPepplos struct {
	Title    string `json:"title" validate:"required,min=10"`
	SubTitle string `json:"subtitle" validate:"required,min=5"`
	Text     string `json:"text" validate:"required,min=10"`
	Author   string `json:"author" validate:"required"`
	Image    string `json:"img_url" validate:"required"`
}
