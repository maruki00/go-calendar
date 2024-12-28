package app

import repositories "go-calendar/internal/calender/infra/Repositories"





type App struct {

	repo repositories.EventRepository
	cfg any
	service services.EventService

}
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


