package controllers

import (
	"go-calendar/internal/calender/app/services"
	"go-calendar/internal/calender/userinterface/requests"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EventController struct {
	srvc      *services.EventService
	validator *validator.Validate
}

func NewEventController(srv *services.EventService) *EventController {
	return &EventController{
		srvc:      srv,
		validator: validator.New(),
	}
}

func (o *EventController) Create(ctx *gin.Context) {
	req := new(requests.CreateRequest)
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if err := o.validator.Struct(req); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	res, err := o.srvc.Create(ctx, req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": nil,
		"data":    res,
	})
}

func (o *EventController) Delete(ctx *gin.Context) {
	req := new(requests.DeleteRequest)
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if err := o.validator.Struct(req); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	res, err := o.srvc.Delete(ctx, req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": nil,
		"data":    res,
	})

}

func (o *EventController) Update(ctx *gin.Context) {
	req := new(requests.UpdateRequest)
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if err := o.validator.Struct(req); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	res, err := o.srvc.Update(ctx, req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": nil,
		"data":    res,
	})

}

func (o *EventController) Get(ctx *gin.Context) {

	res, err := o.srvc.Get(ctx, nil)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": nil,
		"data":    res,
	})

}

func (o *EventController) CreateCommon(ctx *gin.Context) {

	res, err := o.srvc.Get(ctx, nil)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": nil,
		"data":    res,
	})

}

func (o *EventController) DeleteCommon(ctx *gin.Context) {

	res, err := o.srvc.Get(ctx, nil)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": nil,
		"data":    res,
	})

}
