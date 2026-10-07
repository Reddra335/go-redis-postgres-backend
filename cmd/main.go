package main

import (
	"newProjek/internal/shared/config"
	"newProjek/internal/shared/logger"

	"github.com/sirupsen/logrus"
)

func main() {
	// app := InitializeApp()
	// log.Fatal(app.Listen(":3000"))

	// rdb := redis.NewClient(&redis.Options{

	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	// ctx := context.Background()

	// resp, err := rdb.Ping(ctx).Result()

	// if err != nil {
	// 	log.Fatal(err.Error())

	// }

	// fmt.Println(resp)

	log := logger.New()

	db, err := config.InitDB()

	if err != nil {
		log.WithFields(logrus.Fields{
			"operation": "init_db",
			"layer":     "main",
		}).Fatal("Gagal menginisialisasi DB")

	}

	defer db.Close()

}
