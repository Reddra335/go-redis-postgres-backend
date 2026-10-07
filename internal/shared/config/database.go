package config

import (
	"database/sql"
	"fmt"
	"log"
	"newProjek/internal/shared/logger"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB() (*sql.DB, error) {
	config := Load()

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPassword,
		config.DbName,
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
