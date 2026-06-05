package domain

type CommonEvent struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	BackgroundColor string `json:"backgroundColor"`
	BorderColor     string `json:"borderColor"`
}
