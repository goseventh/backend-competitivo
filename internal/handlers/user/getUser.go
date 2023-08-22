package user

import (
	"fmt"
	"main/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandlerFindUser(c *fiber.Ctx) error {
	t := c.Query("t")
	db := mongodb.Builder().
		UseDatabase("backend").
		UseCollection("users")
	searchUsers, _ := db.SearchUsersByTerm(t)
  fmt.Printf("term: %v\n", t)
	type response struct {
		Name     string
		Nickname string
		Birth    string
		Stack    []string
	}
	responseBody := []response{}
	for _, u := range searchUsers {
		userBody := response{
			u.Name,
			u.Nickname,
			u.Birth, u.Stack,
		}
		responseBody = append(responseBody, userBody)
	}
  fmt.Println(searchUsers)
	return c.JSON(responseBody)

}

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
	var responseBody = struct {
		Name     string
		Nickname string
		Birth    string
		Stack    []string
	}{
		user.Name,
		user.Nickname,
		user.Birth,
		user.Stack,
	}
	return c.JSON(responseBody)
}
