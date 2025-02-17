package main

import (
	"fmt"

	"github.com/fmelihh/decentralized-poker-game-go/deck"
)

func main() {
	card := deck.NewCard(deck.Spades, 1)
	fmt.Println(card)
}
