package store

import (
	"encoding/csv"
	"fmt"
	"github.com/lucas-soria/desafio-goweb-lucassoria/internal/domain"
	"os"
	"strconv"
)

type Store struct {
	filepath string
}

func NewStore(filepath string) *Store {
	return &Store{
		filepath: filepath,
	}
}

func (s *Store) Read() (ticketList []domain.Ticket, err error) {
	file, err := os.Open(s.filepath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}
	return
}
