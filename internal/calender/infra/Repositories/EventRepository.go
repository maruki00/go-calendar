package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"go-calendar/internal/calender/domain"
	pkg "go-calendar/pkg/postgres"
	"strings"
)

type EventRepository struct {
	db *pkg.DBHandler
}

func NewEventRepository(db *pkg.DBHandler) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (obj *EventRepository) Home(ctx context.Context) (any, error) {
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
func (obj *EventRepository) CreateCommonEvent(ctx context.Context, event domain.CommonEvent) (any, error) {
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

func (obj *EventRepository) Create(ctx context.Context, event domain.Event) (any, error) {
	sql := `
						INSERT INTO events (id, title, start_at, end_at, background_color, border_color, all_day) 
						VALUES (?,?,?,?,?,?,?)
	`
	_, err := obj.db.GetDB().Exec(sql, event.Id, event.Title, event.StartAt, event.EndAt, event.BackgroundColor, event.BorderColor, event.AllDay)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (obj *EventRepository) Update(ctx context.Context, id string, data map[string]any) (any, error) {

	values := make([]any, 0)
	sql := "UPDATE events SET "
	index := 1
	for col, val := range data {
		sql += fmt.Sprintf(" %s = ?, ", col)
		values = append(values, val)
		index++
	}
	sql = strings.TrimRight(sql, ", ")
	sql += " WHERE id = ? "

	fmt.Println("SQL : ", sql)
	values = append(values, id)
	_, err := obj.db.GetDB().Exec(sql, values...)
	if err != nil {
		return nil, err
	}
	return true, nil
}

func (obj *EventRepository) Delete(ctx context.Context, id string) (any, error) {

	sql := "DELETE FROM events WHERE id  = ?"
	_, err := obj.db.GetDB().Exec(sql, id)
	if err != nil {
		return nil, err
	}
	return true, nil
}

func (obj *EventRepository) DeleteCommonEvent(ctx context.Context, id string) (any, error) {

	sql := "DELETE FROM common_events WHERE id  = ?"
	_, err := obj.db.GetDB().Exec(sql, id)
	if err != nil {
		return nil, err
	}
	return true, nil
}

func (obj *EventRepository) GetAll(ctx context.Context) ([]domain.Event, error) {
	sql := "select id,title,start_at,end_at,all_day,background_color,border_color, from events"
	rows, err := obj.db.GetDB().Query(sql)
	if err != nil {
		return nil, err
	}
	events := getEvents(rows)
	return events, nil
}

// Helpers
func getEvents(rows *sql.Rows) []domain.Event {
	events := make([]domain.Event, 0)
	for rows.Next() {
		event := domain.Event{}
		// id, title, start_at, end_at, background_color, border_color, all_day
		rows.Scan(&event.Id, &event.Title, &event.StartAt, &event.EndAt, &event.BackgroundColor, &event.BorderColor, &event.AllDay)
		events = append(events, event)
	}
	return events

}

func getCommonEvents(rows *sql.Rows) []domain.CommonEvent {
	events := make([]domain.CommonEvent, 0)
	for rows.Next() {
		event := domain.CommonEvent{}
		rows.Scan(&event.Id, &event.Title, &event.BackgroundColor, &event.BorderColor)
		events = append(events, event)
	}
	return events

}
