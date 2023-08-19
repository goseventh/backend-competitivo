package user

import (
	"main/internal/validate"
	"time"

	"fmt"

	"github.com/go-playground/locales/en"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name     string   `json:"name"     validate:"required,alpha,min=4,max=15"`
	Nickname string   `json:"nickname" validate:"required,alphanum,min=4,max=15"`
	Birth    string   `json:"birth"    validate:"required"`
	Stack    []string `json:"stack"    validate:"max=50"`
}

func HandlerUser(c *fiber.Ctx) error {
	var err error
	var errStr string = fmt.Sprintf("✋ Errors!\n\n\n")
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

	if !validate.IsDate(user.Birth) || !validate.IsDateValid(user.Birth) {
		errStr += fmt.Sprintf("❌ Error: Birth must be of type yyyy-mm-dd\n")
	}
	err = validator.GetValidator().Struct(user)
	if err == nil && validate.IsDate(user.Birth) && validate.IsDateValid(user.Birth) {
		return c.SendStatus(201)
	}

	for _, e := range validator.TranslateError(err) {
		errStr += fmt.Sprintf("❌ Error: %s\n", e)
	}

	if !validate.IsDateValid(user.Birth) {
		date, _ := time.Parse("2006-01-02", user.Birth)
		errStr += fmt.Sprintf("❌ Error: Birth cannot be: \"%s\" \n", date.Format("January 02, 2006"))
	}
	// if birthTime

	errStr += fmt.Sprintf("\n\n\n📥 Json: \n\n\n")
	errStr += fmt.Sprintf("🏷️ Atribute: %s, 📦  Value: %v\n", "user", user.Name)
	errStr += fmt.Sprintf("🏷️ Atribute: %s, 📦  Value: %v\n", "nickname", user.Nickname)
	errStr += fmt.Sprintf("🏷️ Atribute: %s, 📦  Value: %v\n", "birth", user.Birth)
	errStr += fmt.Sprintf("🏷️ Atribute: %s, 📦 Value: %v\n", "stack", user.Stack)
	c.WriteString(errStr)
	return c.SendStatus(fiber.StatusNotAcceptable)

}
