package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 24: Lobby Layout\n=====================")

	tileList := readTileList("RawData.txt")
	fmt.Printf("Read tile list: %v\n", tileList)

	fmt.Println("\nPart 1: Count of black tiles\n-----------------------------------------")
	solvePart1(tileList)

	fmt.Println("\nPart 2: Count of black tiles after 100 days\n-------------------------------------------------------")
	solvePart2(tileList)
}

func readTileList(filename string) []Tile {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	tiles := strings.Split(string(buf), "\r\n")

	var res []Tile
	for _, t := range tiles {
		tile := MakeTile(t)
		res = append(res, tile)
	}

	return res
}

func solvePart1(tiles []Tile) {
	tileMap := make(map[Position]TileColor)

	for _, t := range tiles {
		tileMap[t.Pos] = 1 - tileMap[t.Pos]
	}

	res := 0
	for _, color := range tileMap {
		res = res + int(color)
	}

	fmt.Printf("Count of black tiles: %d\n", res)
}

func solvePart2(tiles []Tile) {
	tileMap := make(map[Position]TileColor)

	for _, t := range tiles {
		tileMap[t.Pos] = 1 - tileMap[t.Pos]
		if tileMap[t.Pos] == WHITE {
			delete(tileMap, t.Pos)
		}
	}

	for i := 0; i < 100; i++ {
		tileMap = performDayFlip(tileMap)
		fmt.Printf("Number of black tiles on day %d: %d\n", i+1, len(tileMap))
	}

	res := 0
	for _, color := range tileMap {
		res = res + int(color)
	}

	fmt.Printf("Count of black tiles: %d\n", res)
}

func performDayFlip(tileMap map[Position]TileColor) map[Position]TileColor {
	resMap := make(map[Position]TileColor)

	for pos, color := range tileMap {
		if color == WHITE {
			continue
		}
		blacksBlackNeighborCount := 0
		for _, adjPos := range pos.GetAdjacentPositions() {
			if tileMap[adjPos] == BLACK {
				blacksBlackNeighborCount++
			} else {
				whitesBlackNeighborCount := 0
				for _, wAdjPos := range adjPos.GetAdjacentPositions() {
					if tileMap[wAdjPos] == BLACK {
						whitesBlackNeighborCount++
					}
				}
				if whitesBlackNeighborCount == 2 {
					resMap[adjPos] = BLACK
				}
			}
		}
		if blacksBlackNeighborCount == 1 || blacksBlackNeighborCount == 2 {
			resMap[pos] = BLACK
		}
	}

	for pos, color := range resMap {
		if color == WHITE {
			delete(tileMap, pos)
		}
	}

	return resMap
}
