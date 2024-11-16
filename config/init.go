package config

import (
	"github.com/joho/godotenv"
)

func LoadEnv(path string) {
	_ = godotenv.Load(path)
}
