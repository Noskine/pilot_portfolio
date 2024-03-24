package config

import (
	"log"
	"os"

	"github.com/go-playground/validator"
)

type Env struct {
	ENV_DATABASE_URI string `validate:"required"`
	ENV_WEB_PORT     string `validate:"required,min=4"`
}

func init() {
	env := NewEnvironments()

	valid := validator.New()

	if err := valid.Struct(env); err != nil {
		errors := err.(validator.ValidationErrors)
		log.Fatalf("Loading Error Envirioments: %s", errors)
	}
}

func NewEnvironments() *Env {
	return &Env{
		ENV_DATABASE_URI: os.Getenv("DATABASE_URI"),
		ENV_WEB_PORT:     os.Getenv("PORT"),
	}
}

func (e *Env) GetPort() string {
	return e.ENV_WEB_PORT
}

func (e *Env) GetDatabaseUri() string {
	return e.ENV_DATABASE_URI
}
