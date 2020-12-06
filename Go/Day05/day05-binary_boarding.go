package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Day 05: Binary Boarding\n=======================")

	boardingPasses := readBoardingPasses("RawData.txt")
	fmt.Printf("Read boarding passes: %q\n", boardingPasses)

	fmt.Println("\nPart 1: Highest Seat ID\n-----------------------")
	solvePart1(boardingPasses)

	fmt.Println("\nPart 2: My free seat\n--------------------")
	solvePart2(boardingPasses)
}

func readBoardingPasses(filename string) []BoardingPass {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	lines := strings.Split(string(buf), "\r\n")
	var res []BoardingPass
	for _, line := range lines {
		res = append(res, MakeBoardingPass(line))
	}

	return res
}

func solvePart1(boardingPasses []BoardingPass) {

	maxSeatId := 0
	for _, bp := range boardingPasses {
		if bp.SeatId > maxSeatId {
			maxSeatId = bp.SeatId
		}
	}

	fmt.Printf("Highest Seat ID: %d\n", maxSeatId)
}

func solvePart2(boardingPasses []BoardingPass) {
	sort.Slice(boardingPasses, func(l, r int) bool {
		return boardingPasses[l].SeatId < boardingPasses[r].SeatId
	})

	var freeSeat int
	for i, bp := range boardingPasses[:len(boardingPasses)-1] {
		if bp.SeatId+2 == boardingPasses[i+1].SeatId {
			freeSeat = bp.SeatId + 1
			break
		}
	}

	fmt.Printf("Free Seat ID: %d\n", freeSeat)
}
