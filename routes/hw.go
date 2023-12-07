package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Hw(c *fiber.Ctx) error {
	
	return c.SendString("Hello world!")
}