package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"go-calendar/internal/calender/domain"
	pkg "go-calendar/pkg/postgres"
	"strings"
)

type StatisticsRepository struct {
	db *pkg.DBHandler
}

func NewStatisticsRepository(db *pkg.DBHandler) *StatisticsRepository {
	return &StatisticsRepository{
		db: db,
	}
}

func (obj *StatisticsRepository) Home(ctx context.Context) (any, error) {
	
}
func (obj *StatisticsRepository) CreateCommonEvent(ctx context.Context, event domain.CommonEvent) (any, error) {
}


