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

	//fmt.Println("\nPart 2: Manhattan distance to target via way point\n-------------------------------------------------------")
	//solvePart2(navigationInstructions)
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

//func solvePart2(navInstructions []NavInstruction) {
//	x, y := 0, 0
//	wayPoint := WayPoint{X: 10, Y: 1}
//
//	for _, inst := range navInstructions {
//		switch inst.Command {
//		case 'N', 'E', 'S', 'W', 'R', 'L':
//			wayPoint.Move(inst)
//		case 'F':
//			x = x + wayPoint.X*inst.Value
//			y = y + wayPoint.Y*inst.Value
//		}
//	}
//
//	fmt.Printf("Final position: %d/%d (waypoint at %v)\n", x, y, wayPoint)
//	fmt.Printf("Distance to pos: %d\n", absInt(x)+absInt(y))
//}
