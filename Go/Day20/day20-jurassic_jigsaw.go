package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 20: Jurassic Jigsaw\n========================")

	tiles := readTiles("RawData.txt")
	fmt.Printf("Tiles: %q\n", tiles)

	fmt.Println("\nPart 1: Product of corner-tile ids\n----------------------------------")
	solvePart1(tiles)

	//fmt.Println("\nPart 2: Allergen-alphabetic ingredients list\n-------------------------------------")
	//solvePart2(foodList)
}

func readTiles(filename string) []Tile {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	tiles := strings.Split(string(buf), "\r\n\r\n")
	res := make([]Tile, 0, len(tiles))
	for _, t := range tiles {
		res = append(res, MakeTile(t))
	}

	return res
}

func solvePart1(tiles []Tile) {
	connectedTiles := findPossibleConnections(tiles)
	fmt.Printf("Connected tiles: %v\n", connectedTiles)

	cornerCandidates := make([]Tile, 0)
	for _, t := range connectedTiles {
		if len(t.Connections) == 2 {
			cornerCandidates = append(cornerCandidates, t)
		}
	}
	fmt.Printf("Corner candidates: %v\n", cornerCandidates)

	res := -1
	if len(cornerCandidates) == 4 {
		res = 1
		for _, ct := range cornerCandidates {
			res = res * ct.Id
		}
	}

	fmt.Printf("Product of corner ids: %d\n", res)
}

func findPossibleConnections(tiles []Tile) []Tile {
	//processedTiles := make(map[int]bool)

	for i, t := range tiles {
		for j, u := range tiles[i+1:] {
			for ti := range [...]int{0, 1, 2, 3} {
				for ui := range [...]int{0, 1, 2, 3} {
					if t.BordersClockwise[ti] == u.BordersClockwise[ui] {
						//fmt.Printf("Cw<->Cw: %v <-> %v\n", t.BordersClockwise[ti], u.BordersClockwise[ui])
						newConn := TileConnection{
							IdFrom:        t.Id,
							FromSide:      ti,
							FromDirection: 0,
							IdTo:          u.Id,
							ToSide:        ui,
							ToDirection:   0,
						}
						tiles[i].Connections = append(tiles[i].Connections, newConn)
						tiles[i+1+j].Connections = append(tiles[i+1+j].Connections, newConn)
					}
					if t.BordersClockwise[ti] == u.BordersCounterCw[ui] {
						newConn := TileConnection{
							IdFrom:        t.Id,
							FromSide:      ti,
							FromDirection: 0,
							IdTo:          u.Id,
							ToSide:        ui,
							ToDirection:   1,
						}
						tiles[i].Connections = append(tiles[i].Connections, newConn)
						tiles[i+1+j].Connections = append(tiles[i+1+j].Connections, newConn)
					}
				}
			}
		}
	}

	return tiles
}

//func solvePart2(foodList []Food) {
//	_, allergens := constructContentLists(foodList)
//	determineAllergenContainers(allergens)
//
//	fmt.Printf("Allergens: %v\n", allergens)
//
//	allAllergens := make([]string, 0, len(allergens))
//	for n := range allergens {
//		allAllergens = append(allAllergens, n)
//	}
//	sort.Strings(allAllergens)
//
//	orderedIngredients := make([]string, 0, len(allAllergens))
//	for _, a := range allAllergens {
//		orderedIngredients = append(orderedIngredients, allergens[a].Container)
//	}
//
//	fmt.Printf("Ordered container elements: %s\n", strings.Join(orderedIngredients, ","))
//}
//
