package requests

type CreateRequest struct {
	Title           string `validate:"required" json:"title"`
	StartAt         string `validate:"required" json:"start_at,omitempty"`
	EndAt           string `validate:"required" json:"end_at,omitempty"`
	AllDay          bool   `validate:"required" json:"all_day,omitempty"`
	BackgroundColor string `validate:"required" json:"background_color,omitempty"`
	BorderColor     string `validate:"required" json:"border_color,omitempty"`
	Css             string `validate:"required" json:"css,omitempty"`
}
