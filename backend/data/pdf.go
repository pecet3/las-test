package data

import (
	"github.com/pecet3/las-test-pdf/data/dtos"
)

func (p Pdf) ToDto(d *Queries) *dtos.PDF {
	return &dtos.PDF{
		Name:       p.Name,
		UUID:       p.Uuid,
		CreatedAt:  p.CreatedAt.Time,
		LastOpenAt: p.LastOpenAt.Time,
	}
}
