package tickets

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-soria/desafio-goweb-lucassoria/internal/domain"
)

type Service interface {
	GetTotalTickets(*gin.Context, string) ([]domain.Ticket, error)
	AverageDestination(*gin.Context, string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(ctx *gin.Context, destination string) (ts []domain.Ticket, err error) {
	ts, err = s.repository.GetTicketByDestination(ctx, destination)
	if len(ts) == 0 {
		return []domain.Ticket{}, nil
	}
	return
}

func (s *service) AverageDestination(ctx *gin.Context, destination string) (avg float64, err error) {
	var ticketsTotal []domain.Ticket
	ticketsTotal, err = s.GetTotalTickets(ctx, destination)
	if err != nil {
		return
	}
	var ticketsDestino []domain.Ticket
	ticketsDestino, err = s.GetTotalTickets(ctx, "")
	if err != nil {
		return
	}
	avg = float64(len(ticketsTotal)) / float64(len(ticketsDestino))
	return
}
