package internal

import (
	profilmahasiswa "newProjek/internal/Profil_Mahasiswa"

	"github.com/gofiber/fiber/v3"
)

type Depedencies struct {
	Profil profilmahasiswa.ProfilHandler

	Jwt fiber.Handler
}

func RegisterRoutes(app *fiber.App, dependecies Depedencies) {

	RegisterProfilRouter(app, dependecies)

}

func RegisterProfilRouter(app *fiber.App, depedencies Depedencies) {

	profilmahasiswa.AppProfil(depedencies.Profil, app)

}
func RegisterJwtRouter(app *fiber.App, depedencies Depedencies) {

	profilmahasiswa.AppProfil(depedencies.Profil, app)

}
