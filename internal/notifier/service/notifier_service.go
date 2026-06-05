package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-calendar/internal/notifier/domain"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus/v5"
)

type NotifierService struct {
	queueMux sync.Mutex
}

func NewNotifierService() *NotifierService {
	return &NotifierService{}
}

func (_this *NotifierService) triggerLinuxNotification(title, body, hexColor string) error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return fmt.Errorf("failed linking to session bus: %w", err)
	}

	client, err := notify.New(conn)
	if err != nil {
		return fmt.Errorf("failed spawning notify engine client context: %w", err)
	}

	hints := map[string]dbus.Variant{}

	hints["urgency"] = dbus.MakeVariant(byte(2))
	notification := notify.Notification{
		AppName:       "GoScheduler",
		ReplacesID:    0,
		AppIcon:       "appointment-reminder",
		Summary:       title,
		Body:          body,
		Actions:       []notify.Action{},
		Hints:         hints,
		ExpireTimeout: 12000 * time.Millisecond,
	}

	_, err = client.SendNotification(notification)
	return err
}

func isBetween(start, end, target string) bool {
	_start, _ := time.Parse("15:04", start)
	_end, _ := time.Parse("15:04", end)
	_target, _ := time.Parse("15:04", target)
	return _target.After(_start) && _target.Before(_end)
}

func getFileContent(filePath string) []byte {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	return content
}

func (_this *NotifierService) BackgroundScheduler(ctx context.Context) {
	content := getFileContent("tasks.json")
	events := make(map[string]map[string]domain.Event)
	err := json.Unmarshal(content, &events)
	if err != nil {
		panic(err)
	}
	currDay := strings.ToLower(time.Now().Weekday().String())
	currDayEvent := events[currDay]
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	log.Println("[Engine] Asynchronous database scanning loop initialized...")
	fired := make(map[string]bool)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			now := time.Now().Format("15:04")
			if isBetween("00:00", "01:00", now) {
				fired = make(map[string]bool)
			}
			_this.queueMux.Lock()
			defer _this.queueMux.Unlock()
			alertsToFire := make([]domain.Event, 0)

			for name, task := range currDayEvent {
				if isBetween(task.Start, task.End, now) && fired[name] != true {
					alertsToFire = append(alertsToFire, task)
				}
				fired[name] = true
			}
			log.Println("New Event Executed")
			for _, alert := range alertsToFire {
				log.Printf("[Scheduler Alert] Processing target event tracking item: %s", alert.Title)
				if err := _this.sendToBackend(ctx, alert); err != nil {
					continue
				}
				go func(id int64, title, css, bg string) {
					msg := fmt.Sprintf("Time to pivot tasks!\nPeriod:  (%s:%s)", alert.Start, alert.End)
					if err := _this.triggerLinuxNotification(title, msg, bg); err != nil {
						log.Printf("[OS Link Error] DBus interaction failure: %v", err)
					}
				}(alert.ID, alert.Title, alert.CSS, alert.BackgroundColor)
			}
			_this.queueMux.Unlock()
		}
	}
}

func (_this *NotifierService) sendToBackend(ctx context.Context, payload domain.Event) error {

	url := "http://127.0.0.1:5600/api/v1/event/create"
	payload.Start = time.Now().Format("2006-01-02")
	payload.End = time.Now().Format("2006-01-02")
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Erreur d'encodage JSON:", err)
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erreur lors de la requête:", err)
		return err

	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur de lecture de la réponse:", err)
		return err

	}
	return nil
}
