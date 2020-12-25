package main

import (
	"fmt"
)

var day15Data = [...]int{2, 1, 10, 11, 0, 6}

//var day15Data = [...]int{0, 3, 6}

func main() {
	fmt.Println("Day 15: Report Repair\n=====================")

	gameData := prepareGameData(day15Data[:])
	fmt.Printf("Initial game data: %v\n", gameData)

	fmt.Println("\nPart 1: 2020th number\n--------------------------")
	solvePart1(gameData)

	gameData.Reset()

	fmt.Println("\nPart 2: 30000000 number\n--------------------------")
	solvePart2(gameData)
}

type Occurrences struct {
	PrePreRound int
	PreRound    int
}

type GameData struct {
	StartNumbers []int
	LastPosition int
	Numbers      map[int]Occurrences
	LastNumber   int
}

func (gd GameData) String() string {
	return fmt.Sprintf("GameData(last pos=%d with value %d): %v", gd.LastPosition, gd.LastNumber, gd.Numbers)
}

func (gd *GameData) Reset() {
	gd.Numbers = make(map[int]Occurrences, 0)

	for i, n := range gd.StartNumbers {
		if v, ok := gd.Numbers[n]; ok {
			gd.Numbers[n] = Occurrences{v.PreRound, i}
		} else {
			gd.Numbers[n] = Occurrences{-1, i}
		}
	}

	gd.LastNumber = gd.StartNumbers[len(gd.StartNumbers)-1]
	gd.LastPosition = len(gd.StartNumbers) - 1
}

func (gd *GameData) SetStartingNumbers(numbers []int) {
	gd.StartNumbers = numbers
	//
	//for i, n := range numbers {
	//	if v, ok := gd.Numbers[n]; ok {
	//		gd.Numbers[n] = Occurrences{v.PreRound, i}
	//	} else {
	//		gd.Numbers[n] = Occurrences{-1, i}
	//	}
	//}
	//
	//gd.LastNumber = numbers[len(numbers)-1]
	//gd.LastPosition = len(numbers) - 1

	gd.Reset()
}

func (gd *GameData) UpdateLastNumber() {
	if v, ok := gd.Numbers[gd.LastNumber]; ok {
		gd.Numbers[gd.LastNumber] = Occurrences{v.PreRound, gd.LastPosition}
	} else {
		gd.Numbers[gd.LastNumber] = Occurrences{-1, gd.LastPosition}
	}
}

func MakeGameData() GameData {
	res := GameData{
		LastPosition: 0,
		Numbers:      make(map[int]Occurrences),
	}

	return res
}

func prepareGameData(initialNumbers []int) GameData {
	gameData := MakeGameData()
	gameData.SetStartingNumbers(initialNumbers)

	return gameData
}

func (gd *GameData) DoGameRound() {
	gd.LastPosition++
	if v, ok := gd.Numbers[gd.LastNumber]; ok {
		//fmt.Printf("  Found for last number %d: %v\n", gameData.LastNumber, v)
		if v.PrePreRound > -1 {
			nextNumber := v.PreRound - v.PrePreRound
			gd.LastNumber = nextNumber
			//fmt.Printf("  -> Determined new last number(1): %d\n", gameData.LastNumber)
		} else {
			// case: previous number has not been spoken before
			gd.LastNumber = 0
			//fmt.Printf("  -> Determined new last number(2): %d\n", gameData.LastNumber)
		}
	} else {
		// case: previous number has not been spoken before, but it was not entered into number; should not be possible
		gd.LastNumber = 0
		fmt.Printf(" ERROR -> Determined new last number(3): %d\n", gd.LastNumber)
	}
	gd.UpdateLastNumber()
	//fmt.Printf("Game data after round %d: %v\n", i, gameData)
}

func solvePart1(gameData GameData) {
	remainingRounds := 2020 - gameData.LastPosition - 1
	for i := 0; i < remainingRounds; i++ {
		gameData.DoGameRound()
	}

	//fmt.Printf("Game data after 2020 rounds: %v\n", gameData)
	fmt.Printf("2020th number: %d\n", gameData.LastNumber)
}

func solvePart2(gameData GameData) {
	remainingRounds := 30000000 - gameData.LastPosition - 1
	for i := 0; i < remainingRounds; i++ {
		gameData.DoGameRound()
	}
	//for i := 0; i < 15; i++ {
	//	gameData.DoGameRound()
	//	fmt.Printf("Say: %d\n", gameData.LastNumber)
	//}

	//fmt.Printf("Game data after 30000000 rounds: %v\n", gameData)
	fmt.Printf("30000000th number: %d\n", gameData.LastNumber)
}
