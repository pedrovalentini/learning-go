package main

import "github.com/gofiber/fiber/v2"

// rock paper scissors
func main() {
	app := fiber.New()

	Routes(app)

	app.Listen(":3000")
}

type TodoResponse struct {
	ID        uint   `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func Routes(app fiber.Router) {
	app.Get("/api/", GetGame)
}

func GetGame(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"r": "rock",
		"p": "paper",
		"c": "scissors",
	})
}
