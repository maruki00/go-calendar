package requests

type CreateCommonRequest struct {
	Title           string `validate:"required" json:"title"`
	BackgroundColor string `validate:"required" json:"backgroundColor,omitempty"`
	BorderColor     string `validate:"required" json:"borderColor,omitempty"`
}
