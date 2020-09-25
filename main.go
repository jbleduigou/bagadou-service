package main

import (
	"github.com/gofiber/fiber"
	fiberprometheus "github.com/hepsiburada/fiber-prometheus"
)

func main() {
	app := fiber.New()

	p8sMiddleware := fiberprometheus.NewMiddleware("fiber", "http", "/metrics")
	p8sMiddleware.Register(app)

	app.Get("/:year", func(c *fiber.Ctx) {
		c.Send("Bagad " + c.Params("year"))
	})

	app.Listen(7000)
}
