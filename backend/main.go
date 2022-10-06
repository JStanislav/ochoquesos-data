package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/hi", func(req *fiber.Ctx) error {
		return req.JSON(&fiber.Map{
			"data": "hi from back",
		})
	})

	app.Listen(":3000")
	fmt.Println("Server listening on :3000")

}
