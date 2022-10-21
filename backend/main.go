package main

import (
	"fmt"

	"github.com/JStanislav/ochoquesos-data/config"
	"github.com/JStanislav/ochoquesos-data/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	config.Connect()

	app.Get("/hi", func(req *fiber.Ctx) error {
		return req.JSON(&fiber.Map{
			"data": "hi from back",
		})
	})
	app.Get("/players", handlers.GetPlayers)

	app.Listen(":3000")
	fmt.Println("Server listening on :3000")

}
