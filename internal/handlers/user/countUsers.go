package user

import (
	"fmt"
	"main/database/mongo"

	"github.com/gofiber/fiber/v2"
)

func HandlerCountUsers(c *fiber.Ctx) error {
	db := mongodb.Builder().UseDatabase("backend")
	db.UseCollection("users")
	users, err := db.CountUsers()
	if err != nil {
		c.SendString("❌ internal error")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString(fmt.Sprintf("👌 Total users: %v", users))

}
