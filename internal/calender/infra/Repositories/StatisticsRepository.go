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

func (obj *StatisticsRepository) GetStatistics(ctx context.Context, event domain.CommonEvent) (any, error) {
	return nil, nil
}

func (obj *StatisticsRepository) GetProdactivity(ctx context.Context, event domain.CommonEvent) (any, error) {
	return nil, nil
}


