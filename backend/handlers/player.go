package handlers

import (
	"fmt"

	"github.com/JStanislav/ochoquesos-data/config"
	"github.com/JStanislav/ochoquesos-data/models"
	"github.com/gofiber/fiber/v2"
)

func GetPlayers(c *fiber.Ctx) error {
	var players []models.Player

	config.Database.Find(&players)
	fmt.Println(players)
	return c.Status(200).JSON(players)
}

// func getPlayer(c *fiber.Ctx) error {
// 	id :=
// }
