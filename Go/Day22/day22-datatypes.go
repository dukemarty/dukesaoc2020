package main

import (
	"fmt"
	"os"
	"strconv"
)

type Deck struct {
	Player string
	Cards  []int
}

func (d Deck) String() string {
	return fmt.Sprintf("%s: [%v]", d.Player, d.Cards)
}

func (d *Deck) PopTop() int {
	res := d.Cards[0]

	d.Cards = d.Cards[1:]

	return res
}

func (d *Deck) Append(c1 int, c2 int) {
	d.Cards = append(d.Cards, c1, c2)
}

func (d Deck) IsEmpty() bool {
	return len(d.Cards) == 0
}

func (d Deck) CardsId() string {
	return fmt.Sprintf("%v", d.Cards)
}

func MakeDeck(lines []string) Deck {
	resDeck := Deck{
		Player: lines[0],
	}

	for _, l := range lines[1:] {
		i, err := strconv.Atoi(l)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", l, err)
			os.Exit(1)
		}
		resDeck.Cards = append(resDeck.Cards, i)
	}

	return resDeck
}
