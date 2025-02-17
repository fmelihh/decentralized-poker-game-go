package deck

import "fmt"

type Suit int

func (s Suit) String() string {
	switch s {
	case Spades:
		return "SPADES"
	case Harts:
		return "HARTS"
	case Diamonds:
		return "DIAMONDS"
	case Clubs:
		return "CLUBS"
	default:
		panic("invalid card suit.")
	}
}

const (
	Spades Suit = iota
	Harts
	Diamonds
	Clubs
)

type Card struct {
	suit  Suit
	value int
}

func (c Card) String() string {
	return fmt.Sprintf("%d of %s %s", c.value, c.suit, suitToUnicode(c.suit))
}

func NewCard(s Suit, v int) Card {
	if v > 13 {
		panic("the value of cards cannot be higher then 13.")
	}
	return Card{
		suit:  s,
		value: v,
	}
}

func suitToUnicode(s Suit) string {
	switch s {
	case Spades:
		return "♠"
	case Harts:
		return "♥"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		panic("invalid card suit.")
	}
}
