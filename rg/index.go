package main

import (
	"github.com/gofiber/fiber/v2"
	"golang-basics/rg/models"
	"golang-basics/rg/utils"
)
import "fmt"

type AuthDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	utils.LoadEnv()
	app := fiber.New()
	db := models.SetupDataBase()

	app.Post("/login", func(c *fiber.Ctx) error {
		user := new(AuthDTO)

		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "The user must have a username and password",
			})
		}

		fmt.Println(user)

		return c.JSON(user)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")

}
