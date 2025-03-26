package dtos

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Dto struct{}

type Validatable interface {
	Validate(v *validator.Validate) error
}
type Sendable interface {
	Send(w io.Writer) error
}

func New() Dto {
	return Dto{}
}

var vali = validator.New()

func (Dto) Get(req *http.Request, dto Validatable) error {
	if err := json.NewDecoder(req.Body).Decode(dto); err != nil {
		return err
	}
	if err := dto.Validate(vali); err != nil {
		return err
	}

	return nil
}

func (Dto) Send(w http.ResponseWriter, dto interface{}) error {
	if sendable, ok := dto.(Sendable); ok {
		if err := sendable.Send(w); err != nil {
			return err
		}
		return nil
	}
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		return err
	}
	return nil
}
