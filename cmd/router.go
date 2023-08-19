package main

import (
	"github.com/gofiber/fiber/v2"
  "main/internal/handlers/user"
)

func Routers(app *fiber.App) {
	app.Post("/pessoas", user.HandlerUser)
	app.Get("/pessoas:id", user.HandlerUser)
	app.Get("/pessoas:id?", user.HandlerUser)

}
