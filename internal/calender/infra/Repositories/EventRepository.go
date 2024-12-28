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

func (obj *EventRepository) Home(ctx context.Context) (map[string][]domain.Event, error) {
	res := make(map[string][]domain.Event, 0)

	sqlCommunEvents := `SELECT id,title,start_at,end_at,all_day,background_color,border_color,css from   FROM commun_events`
	sqlTodayEvents := `SELECT  id,title,start_at,end_at,all_day,background_color,border_color,css from  
											FROM events
											WHERE DATE(end_at) <= DATE('now')
											AND DATE(start_at) >= DATE('now');
	`
	rows1, err := obj.db.GetDB().Query(sqlCommunEvents)
	if err != nil {
		return nil, err
	}
	rows2, err := obj.db.GetDB().Query(sqlTodayEvents)
	if err != nil {
		return nil, err
	}

	res["commun_events"] = getEvents(rows1)
	res["toddday_events"] = getEvents(rows2)

	return res, nil

}
func (obj *EventRepository) CreateCommunEvent(ctx context.Context, event domain.Event) (any, error) {
	sql := `
						INSERT INTO commun_events (id, title, start_at, end_at, background_color, border_color, all_day, css) 
						VALUES (?,?,?,?,?,?,?,?)
	`
	_, err := obj.db.GetDB().Exec(sql, event.Id, event.Title, event.StartAt, event.EndAt, event.BackgroundColor, event.BorderColor, event.AllDay, event.Css)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (obj *EventRepository) Create(ctx context.Context, event domain.Event) (any, error) {
	sql := `
						INSERT INTO events (id, title, start_at, end_at, background_color, border_color, all_day, css) 
						VALUES (?,?,?,?,?,?,?,?)
	`
	_, err := obj.db.GetDB().Exec(sql, event.Id, event.Title, event.StartAt, event.EndAt, event.BackgroundColor, event.BorderColor, event.AllDay, event.Css)
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

func (obj *EventRepository) DeleteCommunEvent(ctx context.Context, id string) (any, error) {

	sql := "DELETE FROM commun_events WHERE id  = ?"
	_, err := obj.db.GetDB().Exec(sql, id)
	if err != nil {
		return nil, err
	}
	return true, nil
}

func (obj *EventRepository) GetAll(ctx context.Context) ([]domain.Event, error) {
	sql := "select id,title,start_at,end_at,all_day,background_color,border_color,css from events"
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
		rows.Scan(&event.Id, &event.Title, &event.StartAt, &event.EndAt, &event.AllDay, &event.BackgroundColor, &event.BorderColor, &event.Css)
		events = append(events, event)
	}
	return events

}
