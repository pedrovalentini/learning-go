package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const rock = "rock"
const paper = "paper"
const scissors = "scissors"

// rock paper scissors
func main() {
	app := fiber.New()

	Routes(app)

	app.Listen(":3000")
}

func Routes(app fiber.Router) {
	app.Get("/api/", GetGame)
	app.Get("/api/play/:card1/:card2", PlayGame)
}

func GetGame(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"api":  "/api/play/:card1/:card2",
		"deck": Deck(),
	})
}

func PlayGame(c *fiber.Ctx) error {
	deck := Deck()
	card1, valid1 := deck[c.Params("card1")]
	card2, valid2 := deck[c.Params("card2")]
	if valid1 && valid2 {
		result := Play(card1, card2)
		return c.JSON(&fiber.Map{
			"message": result.message,
			"winner":  result.winner.name,
		})
	}
	return c.JSON(&fiber.Map{
		"error": "Invalid play",
	})
}

// Game Module

type card struct {
	name string
	win  string
}

type playResult struct {
	message string
	winner  card
}

func Play(a card, b card) playResult {
	var result = playResult{}
	if a.win == b.name {
		result.winner = a
	}
	if b.win == a.name {
		result.winner = b
	}
	if result.winner == (card{}) {
		result.message = "No winners"
	} else {
		result.message = fmt.Sprintf("The winner is %s", result.winner.name)
	}
	return result
}

func Deck() map[string]card {
	return map[string]card{
		rock: {
			name: rock,
			win:  scissors,
		},
		paper: {
			name: paper,
			win:  rock,
		},
		scissors: {
			name: scissors,
			win:  paper,
		},
	}
}
