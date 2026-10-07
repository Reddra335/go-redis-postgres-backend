package main

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"

	"profil_service/internal/shared/middleware"
	"profil_service/internal/user/handlerUser"
	"profil_service/internal/user/repositoryUser"
	"profil_service/internal/user/serviceUser"
)

var set = wire.NewSet(
	repositoryUser.NewUsersImpln,
	serviceUser.NewUserService,
	handlerUser.NewUserHandler,
)

func ProvideValidator() *validator.Validate {
	return validator.New()
}

func ProvideJWTMiddleware(userRepo repositoryUser.UsersRepo, db *sql.DB) fiber.Handler {
	return middleware.JWTMiddleware(userRepo, db)
}
