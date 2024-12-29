package requests

type CreateRequest struct {
	Title           string `validate:"required" json:"title"`
	BackgroundColor string `validate:"required" json:backgroundColor",omitempty"`
	BorderColor     string `validate:"required" json:borderColor",omitempty"`
	StartAt         string `validate:"required" json:"start,omitempty"`
	EndAt           string `validate:"required" json:"end,omitempty"`
}
