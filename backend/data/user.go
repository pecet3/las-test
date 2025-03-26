package data

import (
	"github.com/pecet3/las-test-pdf/data/dtos"
)

func (u User) ToDto(d *Queries) (*dtos.User, error) {

	return &dtos.User{
		Name:      u.Name,
		UUID:      u.Uuid,
		CreatedAt: u.CreatedAt.Time,
	}, nil
}
