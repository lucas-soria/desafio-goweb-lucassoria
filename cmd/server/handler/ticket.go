package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-soria/desafio-goweb-lucassoria/internal/tickets"
	"net/http"
)

type Controller struct {
	service tickets.Service
}

func NewController(s tickets.Service) *Controller {
	return &Controller{
		service: s,
	}
}

func (c *Controller) GetTicketsByCountry(ctx *gin.Context) {
	destination := ctx.Param("destination")
	ts, err := c.service.GetTotalTickets(ctx, destination)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	ctx.JSON(http.StatusOK, ts)
	return
}

func (c *Controller) AverageDestination(ctx *gin.Context) {
	destination := ctx.Param("destination")
	avg, err := c.service.AverageDestination(ctx, destination)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	ctx.JSON(200, avg)
}
