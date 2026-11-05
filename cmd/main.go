package main

import (
	"newProjek/internal"
	profilmahasiswa "newProjek/internal/Profil_Mahasiswa"
	"newProjek/internal/shared/config"
	"newProjek/internal/shared/logger"
	sendemail "newProjek/internal/shared/send_email"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

// func main() {
// 	// app := InitializeApp()
// 	// log.Fatal(app.Listen(":3000"))

// 	// rdb := redis.NewClient(&redis.Options{

// 	// 	Addr:     "localhost:6379",
// 	// 	Password: "",
// 	// 	DB:       0,
// 	// })

// 	// ctx := context.Background()

// 	// resp, err := rdb.Ping(ctx).Result()

// 	// if err != nil {
// 	// 	log.Fatal(err.Error())

// 	// }

// 	// fmt.Println(resp)

// 	logger.New()
// 	con := config.Load()
// 	conn, _ := config.InitDB(con)
// 	userrepo := profilmahasiswa.NewRepoMhsImpln(conn)
// 	v := validator.New()
// 	profilmahasiswa.R
// 	profilmahasiswa.NewProfileServiceImpln(userrepo, v)
// 	defer db.Close()

// }
func NewRouter(depen internal.Depedencies) *fiber.App {
	app := fiber.New()

	internal.RegisterRoutes(app, depen)

	return app

}

func main() {

	log := logger.New()
	dataget := config.Load()
	db, err := config.InitDB(dataget)
	defer db.Close()
	redis := config.InitRedis(*dataget)
	email := sendemail.NewImplnSendEmail(dataget)
	v := validator.New()
	if err != nil {

		log.WithError(err).Fatal("Koneksi gagal")

	}
	reds := profilmahasiswa.NewImplnRedisUser(redis)
	repo := profilmahasiswa.NewRepoMhsImpln(db)
	service := profilmahasiswa.NewProfileServiceImpln(repo, v, reds, email)
	mhsprofil := profilmahasiswa.NewProfilHandler(service)
	b := internal.Depedencies{

		Profil: mhsprofil,
	}

	app := NewRouter(b)

	err = app.Listen(":3000")

	if err != nil {
		log.WithError(err).Fatal("Koneksi gagal pada server go")
	}

}
