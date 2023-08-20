package main

import (
	mongodb "main/database/mongo"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
    StreamRequestBody: true,
    Concurrency: 256*2048,
		ServerHeader:  "github.com/goseventh",
		AppName:       "goseventh backend",
    ColorScheme: fiber.DefaultColors,
    
	})
	Routers(app)
  mongodb.Connect()
	app.Listen(":3000")
}
