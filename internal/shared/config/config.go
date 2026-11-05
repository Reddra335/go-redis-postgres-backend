package config

import (
	"newProjek/internal/shared/logger"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// SQL
type ConfigPostgreSQL struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

// Redis
// "redis://<user>:<pass>@localhost:6379"
// User, Host, Port, Password
type ConfigRedis struct {
	User     string
	Password string
	Host     string
	Port     string
}

// send email
// Host,Port,Email,Password
type ConfigSendEmail struct {
	Host     string
	Port     string
	Email    string
	Password string
}

type Config struct {
	DB        ConfigPostgreSQL
	Redis     ConfigRedis
	SendEmail ConfigSendEmail
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
		DB: ConfigPostgreSQL{
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbName:     os.Getenv("DB_NAME"),
		},

		Redis: ConfigRedis{
			User:     os.Getenv("CACHE_USER"),
			Password: os.Getenv("CACHE_PASS"),
			Host:     os.Getenv("CACHE_HOST"),
			Port:     os.Getenv("CACHE_PORT"),
		},
		SendEmail: ConfigSendEmail{
			Host:     os.Getenv("SMTP_HOST"),
			Port:     os.Getenv("SMTP_PORT"),
			Email:    os.Getenv("SMTP_EMAIL"),
			Password: os.Getenv("SMTP_PASS"),
		},
	}
}
