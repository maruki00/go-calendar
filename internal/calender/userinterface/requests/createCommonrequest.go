package requests

type CreateCommonRequest struct {
	Title           string `validate:"required" json:"title"`
	BackgroundColor string `validate:"required" json:"background_color,omitempty"`
	BorderColor     string `validate:"required" json:"border_color,omitempty"`
}
