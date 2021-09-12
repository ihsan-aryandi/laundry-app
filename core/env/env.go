package env

import (
	"github.com/joho/godotenv"
)

type env struct{}

func NewEnv() *env {
	return &env{}
}

func (*env) ReadEnv(fileNames ...string) error {
	return godotenv.Load(fileNames...)
}
