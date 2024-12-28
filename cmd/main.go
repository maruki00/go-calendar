package main

import (
	"fmt"
	"go-calendar/internal/calender/app/services"
	repositories "go-calendar/internal/calender/infra/Repositories"
	"go-calendar/internal/calender/userinterface/controllers"
	pkg "go-calendar/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func main() {


	db := pkg.NewDBHandler("../db/main.db")
	repo := repositories.NewEventRepository(db)
	srv := services.NewEventService(repo)
	ctl := controllers.NewEventController(srv)
	server  := gin.New()
	server.POST("/api/v1/event/create", ctl.Create)  
	server.POST("/api/v1/event/delete", ctl.Delete)  
	server.POST("/api/v1/event/update", ctl.Update)
	server.POST("/api/v1/events", ctl.Get)
	fmt.Println("server running on :5600") 
	server.Run(":5600") 
}
