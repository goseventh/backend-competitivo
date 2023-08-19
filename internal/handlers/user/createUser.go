package user

import "github.com/gofiber/fiber/v2"
import "fmt"

func HandlerCreate(c *fiber.Ctx) error {

	return c.SendString(fmt.Sprintf("Hello, World 👋! [%s]", c.Params("id")))
}
