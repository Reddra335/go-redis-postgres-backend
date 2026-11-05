package config

import (
	"database/sql"
	"fmt"
	"log"
	"newProjek/internal/shared/logger"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB(config *Config) (*sql.DB, error) {

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB.DbHost,
		config.DB.DbPort,
		config.DB.DbUser,
		config.DB.DbPassword,
		config.DB.DbName,
	)

	DB, err := sql.Open("pgx", dns)

	if err != nil {
		log.Fatalf("Gagal memuat database %v", err)
	}
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(25)
	DB.SetConnMaxLifetime(30 * time.Minute)
	DB.SetConnMaxIdleTime(5 * time.Minute)
	log := logger.New()

	err = DB.Ping()
	if err != nil {
		DB.Close()
		return nil, err
	}

	log.Info("Database Berhasil Terhubung")
	return DB, nil

}
