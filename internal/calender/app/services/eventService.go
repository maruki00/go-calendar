package services

import (
	"context"
	"go-calendar/internal/calender/domain"
	repositories "go-calendar/internal/calender/infra/Repositories"
	"go-calendar/internal/calender/userinterface/requests"

	"github.com/google/uuid"
)

type EventService struct {
	repo *repositories.EventRepository
}

func NewEventService(repo *repositories.EventRepository) *EventService {
	return &EventService{
		repo: repo,
	}
}

// Create method  
// Create a new Event
func (srv *EventService) Create(ctx context.Context, req *requests.CreateRequest) (any, error) {
	_ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return srv.repo.Create(_ctx, domain.Event{
		Id:              uuid.NewString(),
		Title:           req.Title,
		StartAt:         req.StartAt,
		EndAt:           req.EndAt,
		AllDay:          true,
		BackgroundColor: req.BackgroundColor,
		BorderColor:     req.BorderColor,
	})
}

func (srv *EventService) CreateCommonEvent(ctx context.Context, req *requests.CreateCommonRequest) (any, error) {
	_ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return srv.repo.CreateCommonEvent(_ctx, domain.CommonEvent{
		Id:              uuid.NewString(),
		Title:           req.Title,
		BackgroundColor: req.BackgroundColor,
		BorderColor:     req.BorderColor,
	})
}

// Update The event
func (srv *EventService) Update(ctx context.Context, req *requests.UpdateRequest) (any, error) {
	_ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return srv.repo.Update(_ctx, req.Id, req.Fields)
}

// Delete  
func (srv *EventService) Delete(ctx context.Context, req *requests.DeleteRequest) (any, error) {
	_ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return srv.repo.Delete(_ctx, req.Id)
}

func (srv *EventService) DeleteCommunEvent(ctx context.Context, req *requests.DeleteCommonRequest) (any, error) {
	_ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return srv.repo.DeleteCommonEvent(_ctx, req.Id)
}

// Get  
func (srv *EventService) Get(ctx context.Context, req *requests.GetRequest) (any, error) {
	_ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return srv.repo.GetAll(_ctx)
}

// Get  
func (srv *EventService) Home(ctx context.Context, req *requests.GetRequest) (any, error) {
	_ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return srv.repo.Home(_ctx)
}
