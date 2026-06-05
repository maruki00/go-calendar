package app

import (
	"context"
	"go-calendar/internal/notifier/service"
)

type App struct {
	Svc *service.NotifierService
}

func Init() (*App, error) {
	svc := service.NewNotifierService()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
		svc.BackgroundScheduler(ctx)
	}(ctx)
	return &App{}, nil
}
