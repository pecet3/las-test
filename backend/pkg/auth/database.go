package auth

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pecet3/las-test-pdf/data"
)

func (a Auth) AddUserTablesDb(name, email string) (*data.User, error) {
	ctx := context.Background()
	u, err := a.d.AddUser(ctx, data.AddUserParams{
		Salt: uuid.NewString(),
		Uuid: uuid.NewString(),
		Name: name,
		Email: sql.NullString{
			String: email,
			Valid:  true,
		},
		IsDraft: true,
	})
	if err != nil {
		return nil, err
	}

	return &u, nil
}
