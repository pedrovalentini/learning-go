package main

import (
	"github.com/gofiber/fiber/v2"
)

// rock paper scissors
func main() {
	app := fiber.New()

	Routes(app)

	app.Listen(":3000")
}

func Routes(app fiber.Router) {
	app.Post("/api/table/:table/round/finish", FinishRound)
}

func FinishRound(c *fiber.Ctx) error {
	hand1 := c.Params("hand1")
	hand2 := c.Params("hand2")
	board := c.Params("board")
	// if valid1 && valid2 && valid3 {
	result := GetResult([]string{hand1, hand2}, board)
	return c.JSON(&fiber.Map{
		"message": result.message,
	})
	// }
	// return c.JSON(&fiber.Map{
	// 	"error": "Invalid play",
	// })
}

func GetResult(hands []string, board string) gameResult {
	return gameResult{
		message: "Game result...",
	}
}

type gameResult struct {
	message string
}
