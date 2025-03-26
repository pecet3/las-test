package dtos

import (
	"encoding/json"
	"io"
	"time"
)

type PDF struct {
	UUID       string    `json:"uuid"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	LastOpenAt time.Time `json:"last_open_at"`
	URL        string    `json:"url"`
}

func (p PDF) Send(w io.Writer) error {
	if err := json.NewEncoder(w).Encode(&p); err != nil {
		return err
	}
	return nil
}
