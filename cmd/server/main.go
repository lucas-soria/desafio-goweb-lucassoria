package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-soria/desafio-goweb-lucassoria/cmd/server/handler"
	"github.com/lucas-soria/desafio-goweb-lucassoria/internal/tickets"
	"github.com/lucas-soria/desafio-goweb-lucassoria/pkg/store"
)

func main() {
	ticketStore := store.NewStore("./tickets.csv")
	ticketRepository := tickets.NewRepository(ticketStore)
	ticketService := tickets.NewService(ticketRepository)
	ticketController := handler.NewController(ticketService)
	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	// Rutas a desarrollar:
	ts := engine.Group("/tickets")
	{
		// GET - “/tickets/getByCountry/:dest”
		ts.GET("/getbycountry/:destination", ticketController.GetTicketsByCountry)
		// GET - “/tickets/getAverage/:dest”
		ts.GET("/getaverage/:destination", ticketController.AverageDestination)
		// GET - “/tickets”
		ts.GET("/", ticketController.GetTicketsByCountry)
	}
	if err := engine.Run(); err != nil {
		panic(err)
	}
}
