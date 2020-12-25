package main

import (
	"container/ring"
	"fmt"
	"strings"
)

// test
//var day23Data = [...]int{3, 8, 9, 1, 2, 5, 4, 6, 7}

// puzzle
var day23Data = [...]int{3, 1, 8, 9, 4, 6, 5, 7, 2}

func main() {
	fmt.Println("Day 23: Crab Cups\n=====================")

	gameData := prepareGameData(day23Data[:])
	fmt.Printf("Initial game data: %v\n", gameData)

	fmt.Println("\nPart 1: Order after 100 moves\n--------------------------")
	solvePart1(gameData)

	gameData2 := prepareExtendedGameData(day23Data[:])

	fmt.Println("\nPart 2: Product of numbers after 1\n----------------------------------")
	solvePart2(gameData2)
}

type GameData struct {
	InitialNumbers []int
	Min            int
	Max            int
	Numbers        *ring.Ring
}

func (gd GameData) String() string {
	s := make([]string, 0, len(gd.InitialNumbers))
	for i := 0; i < gd.Numbers.Len(); i++ {
		s = append(s, fmt.Sprintf("%d", gd.Numbers.Value))
		gd.Numbers = gd.Numbers.Next()
	}

	return fmt.Sprintf("GameData state: [ %s ]", strings.Join(s, " "))
}

func (gd *GameData) Reset() {
	gd.Numbers = ring.New(len(gd.InitialNumbers))

	for _, i := range gd.InitialNumbers {
		gd.Numbers.Value = i
		gd.Numbers = gd.Numbers.Next()
	}
}

func MakeGameData(initialNumbers []int) GameData {
	res := GameData{
		InitialNumbers: initialNumbers,
		Min:            initialNumbers[0],
		Max:            initialNumbers[0],
		Numbers:        ring.New(len(initialNumbers)),
	}

	for _, i := range initialNumbers {
		res.Numbers.Value = i
		res.Numbers = res.Numbers.Next()

		if res.Min > i {
			res.Min = i
		}
		if res.Max < i {
			res.Max = i
		}
	}

	return res
}

func prepareGameData(initialNumbers []int) GameData {
	gameData := MakeGameData(initialNumbers)

	return gameData
}

func prepareExtendedGameData(initialNumbers []int) GameData {
	max := initialNumbers[0]
	for _, i := range initialNumbers {
		if i > max {
			max = i
		}
	}
	for len(initialNumbers) < 1000000 {
		max++
		initialNumbers = append(initialNumbers, max)
	}

	gameData := MakeGameData(initialNumbers)

	return gameData
}

func (gd *GameData) MakeMove2() {
	current := gd.Numbers.Value.(int)
	pickup := gd.Numbers.Unlink(3)

	target := current
	for found := false; !found; {
		target--
		if target < gd.Min {
			target = gd.Max
		}

		if !ringbufferContains(pickup, target) {
			found = true
		}
	}

	dist := 0
	for gd.Numbers.Value.(int) != target {
		gd.Numbers = gd.Numbers.Next()
		dist++
	}

	gd.Numbers.Link(pickup)

	gd.Numbers = gd.Numbers.Move(1000000 - dist)
}

func (gd *GameData) MakeMove() {
	current := gd.Numbers.Value.(int)
	pickup := gd.Numbers.Unlink(3)

	target := current
	for found := false; !found; {
		target--
		if target < gd.Min {
			target = gd.Max
		}

		if !ringbufferContains(pickup, target) {
			found = true
		}
	}

	for gd.Numbers.Value.(int) != target {
		gd.Numbers = gd.Numbers.Next()
	}

	gd.Numbers.Link(pickup)

	for gd.Numbers.Value.(int) != current {
		gd.Numbers = gd.Numbers.Next()
	}
	gd.Numbers = gd.Numbers.Next()
}

func ringbufferContains(rb *ring.Ring, value int) bool {
	res := false
	for i := 0; i < 3; i++ {
		if value == rb.Value {
			res = true
		}
		rb = rb.Next()
	}

	return res
}

func solvePart1(gameData GameData) {
	for i := 0; i < 100; i++ {
		gameData.MakeMove2()
		fmt.Printf("%d: %v\n", i, gameData)
	}
	fmt.Printf("Game data after 100 rounds: %v\n", gameData)

	for gameData.Numbers.Value.(int) != 1 {
		gameData.Numbers = gameData.Numbers.Next()
	}

	fmt.Printf("Ordered result after 100 moves: %v\n", gameData)

	resNumbers := make([]string, 0, len(gameData.InitialNumbers)-1)
	for i := 0; i < gameData.Numbers.Len()-1; i++ {
		gameData.Numbers = gameData.Numbers.Next()
		resNumbers = append(resNumbers, fmt.Sprintf("%d", gameData.Numbers.Value.(int)))
	}
	res := strings.Join(resNumbers, "")
	fmt.Printf("Result in required format: %s\n", res)
}

func solvePart2(gameData GameData) {
	for j := 0; j < 10000; j++ {
		for i := 0; i < 1000; i++ {
			gameData.MakeMove2()
		}
		println(j)
	}

	for gameData.Numbers.Value.(int) != 1 {
		gameData.Numbers = gameData.Numbers.Next()
	}

	gameData.Numbers = gameData.Numbers.Next()
	rn1 := gameData.Numbers.Value.(int)
	gameData.Numbers = gameData.Numbers.Next()
	rn2 := gameData.Numbers.Value.(int)
	fmt.Printf("Result numbers after 10000000 moves: %d  %d\n", rn1, rn2)

	fmt.Printf("Product of the result numbers: %d\n", rn1*rn2)
}
