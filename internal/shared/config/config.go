package config

import (
	"newProjek/internal/shared/logger"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

func Load() *Config {

	log := logger.New()
	err := godotenv.Load()
	if err != nil {

		log.WithFields(logrus.Fields{
			"operation": "Load_Config",
			"layer":     "config",
		}).WithError(err).Error("data  env tidak ditemukan")
	}

	return &Config{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
	}
}
