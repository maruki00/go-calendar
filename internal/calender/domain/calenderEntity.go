package domain

/**
{
  	// title: 'Lunch',
  	// start: new Date(y, m, d, 12, 0),
  	// end: new Date(y, m, d, 14, 0),
  	// allDay: false,
  	// backgroundColor: '#00c0ef', //Info (aqua)
  	// borderColor: '#00c0ef' //Info (aqua)
}
*/

type Event struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	StartAt         string `json:"start"`
	EndAt           string `json:"end"`
	BackgroundColor string `json:"backgroundColor"`
	BorderColor     string `json:"borderColor"`
	AllDay          bool   `json:"allDay"`
	Css             string `json:"css"`
}
