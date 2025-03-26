package dtos

import (
	"encoding/json"
	"io"
	"time"

	"github.com/go-playground/validator/v10"
)

type Register struct {
	Name  string `json:"name" validate:"required,min=2,max=16,alphanumunicode"`
	Email string `json:"email" validate:"required,email,max=64"`
}

func (r Register) Validate(v *validator.Validate) error {
	err := v.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

type Exchange struct {
	Email string `json:"email" validate:"required,email,max=64"`
	Code  string `json:"code" validate:"required,min=6,max=6"`
}

func (r Exchange) Validate(v *validator.Validate) error {
	err := v.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

type Login struct {
	Email string `json:"email" validate:"required,email,max=64"`
}

func (r *Login) Validate(v *validator.Validate) error {
	err := v.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

type EmailChange struct {
	Email string `json:"email" validate:"required,email,max=64"`
}

func (r *EmailChange) Validate(v *validator.Validate) error {
	err := v.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	ImageUrl  string    `json:"image_url"`
	IsDraft   bool      `json:"is_draft"`
	CreatedAt time.Time `json:"created_at"`
	Exp       float64   `json:"exp"`
	Level     int       `json:"level"`
	Progress  float64   `json:"progress"`
}

func (u User) Send(w io.Writer) error {
	if err := json.NewEncoder(w).Encode(&u); err != nil {
		return err
	}
	return nil
}
