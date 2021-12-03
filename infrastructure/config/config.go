package config

import (
	"os"

	"giapps/servisin/infrastructure/exception"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
	GetOrDefault(key string, valueDefault string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func (config *configImpl) GetOrDefault(key string, valueDefault string) string {
	value := config.Get(key)
	if value != "" {
		return value
	} else {
		return valueDefault
	}
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)

	exception.PanicIfNeeded(err)
	return &configImpl{}
}
