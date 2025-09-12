package main

import (
	"fmt"
	"go-calendar/internal/calender/app/services"
	repositories "go-calendar/internal/calender/infra/Repositories"
	"go-calendar/internal/calender/userinterface/controllers"
	pkg "go-calendar/pkg/postgres"
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Expose-Headers", "Content-Length, X-Custom-Header")
			c.Header("Access-Control-Max-Age", "86400")
		}
		if c.Request.Method == http.MethodOptions {
			c.JSON(http.StatusOK, nil)
			return
		}
		c.Next()
	}
}

func CORS1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.Header("Access-Control-Allow-Origin", "http://127.0.0.1")
			headers := []string{"Content-Type", "Authorization"}
			ctx.Header("Access-Control-Allow-Headers", strings.Join(headers, ","))
			methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}
			ctx.Header("Access-Control-Allow-Methods", strings.Join(methods, ","))
			ctx.Status(200)
			return
		} 
		ctx.Next()
	}
}

func main() {
	db := pkg.NewDBHandler("
	db := pkg.NewDBHandler("./db.db")
	srv := services.NewEventService(repo)
	ctl := controllers.NewEventController(srv)
	server := gin.Default()
	server.Use(CORS())
	html := template.Must(template.ParseFiles("./index.html"))
	server.SetHTMLTemplate(html)
	server.Static("/static", "/opt/calendar/static")

	server.POST("/api/v1/event/create", ctl.Create)
	server.POST("/api/v1/event/delete", ctl.Delete)
	server.POST("/api/v1/event/update", ctl.Update)
	server.POST("/api/v1/events", ctl.Get)
	server.POST("/api/v1/event/common/create", ctl.CreateCommon)
	server.POST("/api/v1/event/common/delete", ctl.DeleteCommon)
	server.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	fmt.Println("server running on :5600")
	server.Run(":5600")
}
