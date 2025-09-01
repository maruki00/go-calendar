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
	res := make(map[string]any, 0)

	sqlCommunEvents := `SELECT *  
						FROM common_events`
	sqlTodayEvents := `	SELECT id, title, start_at, end_at, background_color, border_color, all_day
				 		FROM events 
						-- WHERE strftime('%Y-%m', start_at) = strftime('%Y-%m', 'now') 
						-- OR strftime('%Y-%m', end_at) = strftime('%Y-%m', 'now') 
						-- OR (start_at <= 'now' AND end_at >= 'now')
	`
	rows1, err := obj.db.GetDB().Query(sqlCommunEvents)
	if err != nil {
		return nil, err
	}
	rows2, err := obj.db.GetDB().Query(sqlTodayEvents)
	if err != nil {
		return nil, err
	}

	res["common_events"] = getCommonEvents(rows1)
	res["events"] = getEvents(rows2)

	return res, nil

}
func (obj *StatisticsRepository) CreateCommonEvent(ctx context.Context, event domain.CommonEvent) (any, error) {
	sql := `
						INSERT INTO common_events (id, title, background_color, border_color) 
						VALUES (?,?,?,?)
	`
	_, err := obj.db.GetDB().Exec(sql, event.Id, event.Title, event.BackgroundColor, event.BorderColor)
	if err != nil {
		return nil, err
	}
	return event, nil
}


