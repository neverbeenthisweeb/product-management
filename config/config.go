package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	GetString(key string) string
}

type configImpl struct{}

func (c *configImpl) GetString(key string) string {
	return os.Getenv(key)
}

func NewConfigImpl(files ...string) *configImpl {
	err := godotenv.Load(files...)
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	return &configImpl{}
}
