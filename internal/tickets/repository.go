package tickets

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucas-soria/desafio-goweb-lucassoria/internal/domain"
	"github.com/lucas-soria/desafio-goweb-lucassoria/pkg/store"
)

type Repository interface {
	GetAll(ctx *gin.Context) (ts []domain.Ticket, err error)
	GetTicketByDestination(ctx *gin.Context, destination string) (ts []domain.Ticket, err error)
}

type repository struct {
	db store.Store
}

func NewRepository(store *store.Store) Repository {
	return &repository{
		db: *store,
	}
}

func (r *repository) GetAll(ctx *gin.Context) (ts []domain.Ticket, err error) {
	ts, err = r.db.Read()
	if err != nil {
		return
	}
	if len(ts) == 0 {
		err = fmt.Errorf("empty list of tickets")
		return
	}
	return
}

func (r *repository) GetTicketByDestination(ctx *gin.Context, destination string) (tsByDestination []domain.Ticket, err error) {
	var ts []domain.Ticket
	ts, err = r.GetAll(ctx)
	if err != nil {
		return
	}
	if destination != "" {
		for _, t := range ts {
			if t.Country == destination {
				tsByDestination = append(tsByDestination, t)
			}
		}
	} else {
		tsByDestination = ts
	}
	return
}
