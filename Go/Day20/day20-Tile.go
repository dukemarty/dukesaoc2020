package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Border int
type Direction int

const (
	UP    Border = 0
	RIGHT Border = 1
	DOWN  Border = 2
	LEFT  Border = 3

	CLOCKWISE        Direction = 0
	COUNTERCLOCKWISE Direction = 1
)

// Borders: 0 is Up, 1 is Right, 2 is Down, 3 is Left
type Tile struct {
	Id               int
	RawTile          []string
	Borders          [][]string
	BordersClockwise [4]string
	BordersCounterCw [4]string
	Connections      []TileConnection
	Neighbors        []Connection
}

func (t Tile) String() string {
	return fmt.Sprintf("Tile %d (Connections: %v)", t.Id, t.Connections)
}

func MakeTile(tile string) Tile {
	lines := strings.Split(tile, "\r\n")

	regex, _ := regexp.Compile(`^Tile (\d+):$`)
	matches := regex.FindStringSubmatch(lines[0])
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", matches[1], err)
		return Tile{}
	}

	resTile := Tile{
		Id:      id,
		RawTile: lines[1:len(lines)],
		Borders: make([][]string, 2),
	}

	resTile.Borders[CLOCKWISE] = make([]string, 4)
	resTile.Borders[COUNTERCLOCKWISE] = make([]string, 4)

	resTile.Borders[CLOCKWISE][UP] = lines[1]
	resTile.Borders[COUNTERCLOCKWISE][DOWN] = lines[len(lines)-1]

	for _, side := range []Border{UP, RIGHT, DOWN, LEFT} {
		if resTile.Borders[CLOCKWISE] == ""
	}


	sb1, sb3 := strings.Builder{}, strings.Builder{}
	for i := 0; i < len(lines)-1; i++ {
		sb1.WriteByte(lines[1+i][len(lines[1])-1])
		sb3.WriteByte(lines[1+i][0])
	}
	resTile.BordersClockwise[RIGHT] = sb1.String()
	resTile.BordersCounterCw[RIGHT] = Reversed(resTile.BordersClockwise[1])
	resTile.BordersCounterCw[LEFT] = sb3.String()
	resTile.BordersClockwise[LEFT] = Reversed(resTile.BordersCounterCw[3])

	return resTile
}
