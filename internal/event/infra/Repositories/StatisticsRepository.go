package repositories

import (
	"context"
	pkg "go-calendar/pkg/postgres"
)

type StatisticsRepository struct {
	db *pkg.DBHandler
}

func NewStatisticsRepository(db *pkg.DBHandler) *StatisticsRepository {
	return &StatisticsRepository{
		db: db,
	}
}

func (obj *StatisticsRepository) GetStatistics(ctx context.Context, domain any) (any, error) {
	return nil, nil
}

func (obj *StatisticsRepository) GetProdactivity(ctx context.Context, domain any) (any, error) {
	return nil, nil
}

func (obj *StatisticsRepository) GetStatisticsByDay(ctx context.Context, domain any) (any, error) {
	return nil, nil
}

func (obj *StatisticsRepository) GetStatisticsByMonth(ctx context.Context, domain any) (any, error) {
	return nil, nil
}
