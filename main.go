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
	palmares["2000"] = "Bagad Kemper"
	palmares["2002"] = "Kerlenn Pondi"
	palmares["2002"] = "Bagad Kemper"
	palmares["2003"] = "Bagad Roñsed-Mor"
	palmares["2004"] = "Bagad Kemper"
	palmares["2005"] = "Kevrenn Alre"
	palmares["2006"] = "Kevrenn Alre"
	palmares["2007"] = "Bagad Brieg"
	palmares["2008"] = "Bagad Cap Caval"
	palmares["2009"] = "Bagad Cap Caval"
	palmares["2010"] = "Bagad Cap Caval"
	palmares["2011"] = "Bagad Kemper"
	palmares["2012"] = "Bagad Kemper"
	palmares["2013"] = "Bagad Kemper"
	palmares["2014"] = "Bagad Kemper"
	palmares["2015"] = "Bagad Cap Caval"
	palmares["2016"] = "Bagad Cap Caval"
	palmares["2017"] = "Bagad Cap Caval"
	palmares["2018"] = "Bagad Cap Caval"
	palmares["2019"] = "Bagad Cap Caval"

	app.Get("/bagad/:year", func(c *fiber.Ctx) {
		y := c.Params("year")
		winner, found := palmares[y]
		if !found {
			c.Status(404)
		} else {
			champ := &champion{Name: winner, Year: y}
			c.JSON(champ)
		}
	})

	app.Listen(7000)
}

type champion struct {
	Year string
	Name string
}
