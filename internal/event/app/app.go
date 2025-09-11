package app

import (
	"go-calendar/internal/calender/app/services"
	repositories "go-calendar/internal/calender/infra/Repositories"
)

type App struct {
	repo    repositories.EventRepository
	cfg     any
	service services.EventService
}
