package app

import (
	"go-calendar/internal/event/app/services"
	repositories "go-calendar/internal/event/infra/Repositories"
)

type App struct {
	repo    repositories.EventRepository
	cfg     any
	service services.EventService
}
