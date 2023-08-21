package user

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"main/database/mongo"
)

func HandlerGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := mongodb.Builder().
		UseDatabase("backend").
		UseCollection("users")
	user, err := db.GetUserByUUID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.SendString("🔎 ❌ not found")
			return c.SendStatus(fiber.StatusNotFound)
		}
		c.SendString("💥 internal error")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
  var responseBody = struct{
    Name string
    Nickname string
    Birth string
    Stack []string
  }{
      user.Name,
      user.Nickname,
      user.Birth,
      user.Stack,
    }
	return c.JSON(responseBody)
}
