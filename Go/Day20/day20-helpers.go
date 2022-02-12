package main

import (
	"errors"
	"strings"
)

func Reversed(s string) string {
	sb := strings.Builder{}

	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}

	return sb.String()
}

func FindTileById(tiles []Tile, id int) (Tile, error) {
	for _, t := range tiles {
		if t.Id == id {
			return t, nil
		}
	}

	return Tile{}, errors.New("could not find tile with provided ID")
}
