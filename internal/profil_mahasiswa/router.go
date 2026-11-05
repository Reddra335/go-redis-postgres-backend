package profilmahasiswa

import "github.com/gofiber/fiber/v3"

func AppProfil(h ProfilHandler, a *fiber.App) {

	api := a.Group("/api")
	api.Post("/register", h.Register)
	api.Post("/otp", h.VerifyOTP)
	api.Post("/clear-cookie", h.ClearCookie)
}
