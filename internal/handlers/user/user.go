package user

import (
	"main/internal/validate"

	"github.com/go-playground/locales/en"
	// *"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	// ut "github.com/go-playground/universal-translator"

	"fmt"

	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type User struct {
	Name     string   `validate:"required,min=4,max=15"`
	Nickname string   `validate:"required,min=4,max=15"`
	Birth    string   `validate:"required,min=4,max=15"`
	Stack    []string `validate:"required,min=4,max=15"`
}

func HandlerUser(c *fiber.Ctx) error {
	user := User{}

	validate := validate.Builder()
	validate.
		Language(en.New()).
		Translator("en")

	enTranslations.RegisterDefaultTranslations(
		validate.GetValidator(),
		validate.GetTranslator(),
	)

	err := validate.GetValidator().Struct(user)
	if err == nil {
		return c.SendStatus(200)
	}

	var errStr string = fmt.Sprintf("✋ Errors!\n\n\n")
	for _, e := range validate.TranslateError(err) {
		errStr += fmt.Sprintf("❌ Error: %s\n", e)
	}

	return c.SendString(errStr)
}
