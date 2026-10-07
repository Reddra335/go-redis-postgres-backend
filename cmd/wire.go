//go:build wireinject
// +build wireinject

package main

import (
	"profil_service/internal/shared/config"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func InitializeApp() *fiber.App {
	wire.Build(
		set,
		config.InitDB,
		ProvideValidator,
		NewRouter,
		ProvideJWTMiddleware,
	)
	return nil
}
