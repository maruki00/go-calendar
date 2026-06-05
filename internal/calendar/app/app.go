package app

import (
	"go-calendar/internal/calendar/app/services"
	repositories "go-calendar/internal/calendar/infra/Repositories"
)

type App struct {
	repo    repositories.EventRepository
	cfg     any
	service services.EventService
}
