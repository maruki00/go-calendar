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
	Id              string `yaml:"id" json:"id,omitempty"`
	Title           string `yaml:"title" json:"title"`
	StartAt         string `yaml:"start_at" json:"start_at,omitempty"`
	EndAt           string `yaml:"end_at" json:"end_at,omitempty"`
	AllDay          bool   `yaml:"all_day" json:"all_day,omitempty"`
	BackgroundColor string `yaml:"background_color" json:"background_color,omitempty"`
	BorderColor     string `yaml:"border_color" json:"border_color,omitempty"`
	Css             string `yaml:"css" json:"css,omitempty"`
}
