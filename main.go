package main

import (
	"aziz/k8s/config"
	"aziz/k8s/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.GetClient()
	//routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/node/:label?", controller.GetNodes)

	app.Listen(":3000")
}
