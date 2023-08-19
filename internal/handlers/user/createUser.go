package user

import (
	"main/internal/validate"

	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/gofiber/fiber/v2"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type User struct {
	Name     string   `json:"name"     validate:"required,alpha,min=4,max=15"`
	Nickname string   `json:"nickname" validate:"required,alphanum,min=4,max=15"`
	Birth    string   `json:"birth"    validate:"required"`
	Stack    []string `json:"stack"    validate:"max=50"`
}

func HandlerUser(c *fiber.Ctx) error {
	var err error
	c.Accepts("application/json")
	user := User{}
	c.BodyParser(&user)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	validator := validate.Builder()
	validator.
		Language(en.New()).
		Translator("en")

	enTranslations.RegisterDefaultTranslations(
		validator.GetValidator(),
		validator.GetTranslator(),
	)

	err = validator.GetValidator().Struct(user)
	if err == nil && validate.IsDate(user.Birth){
		return c.SendStatus(201)
	}

	var errStr string = fmt.Sprintf("✋ Errors!\n\n\n")
	for _, e := range validator.TranslateError(err) {
		errStr += fmt.Sprintf("❌ Error: %s\n", e)
	}

	if !validate.IsDate(user.Birth) {
		errStr += fmt.Sprintf("❌ Error: Birth must be of type yyyy-mm-dd\n")
	}

	errStr += fmt.Sprintf("\n\n\n📥 Json: \n\n\n")
	errStr += fmt.Sprintf("🏷️ Atribute: %s, 📦  Value: %v\n", "user", user.Name)
	errStr += fmt.Sprintf("🏷️ Atribute: %s, 📦  Value: %v\n", "nickname", user.Nickname)
	errStr += fmt.Sprintf("🏷️ Atribute: %s, 📦  Value: %v\n", "birth", user.Birth)
	errStr += fmt.Sprintf("🏷️ Atribute: %s, 📦 Value: %v\n", "stack", user.Stack)
	c.WriteString(errStr)
	return c.SendStatus(fiber.StatusNotAcceptable)
}
