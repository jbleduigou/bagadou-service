package main

import (
	"github.com/gofiber/fiber"
	fiberprometheus "github.com/hepsiburada/fiber-prometheus"
)

func main() {
	app := fiber.New()

	p8sMiddleware := fiberprometheus.NewMiddleware("fiber", "http", "/metrics")
	p8sMiddleware.Register(app)

	palmares := make(map[string]string)
	palmares["2010"] = "Bagad Cap Caval"

	app.Get("/bagad/:year", func(c *fiber.Ctx) {
		y := c.Params("year")
		winner, found := palmares[y]
		if !found {
			c.Send("Not found")
		} else {
			c.Send(winner)
		}
	})

	app.Listen(7000)
}
