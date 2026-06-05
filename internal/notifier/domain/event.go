package domain

type Event struct {
	ID              int64  `json:"id,omitempty"`
	Title           string `json:"title"`
	BackgroundColor string `json:"backgroundColor"`
	BorderColor     string `json:"borderColor"`
	Start           string `json:"start"`
	End             string `json:"end"`
	CSS             string `json:"css"`
	Description     string `json:"description"`
}
