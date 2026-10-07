package main

import (
	"profil_service/internal/user"
	"profil_service/internal/user/handlerUser"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(userHandler handlerUser.UserHandler, jwtMiddleware fiber.Handler) *fiber.App {
	app := fiber.New()
	user.RegisterRoutes(app, userHandler, jwtMiddleware) // oper ke domain
	return app
}
