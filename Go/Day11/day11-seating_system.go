package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 11: Seating System\n===========================")

	seatPlan := readSeatPlan("RawData.txt")
	fmt.Printf("Initial seat plan layout:\n%s\n", strings.Join(seatPlan, "\n"))

	fmt.Println("\nPart 1: Count seats after stabilization\n----------------------------------------")
	solvePart1(seatPlan)

	fmt.Println("\nPart 2: Count seats after stabilization with visibility\n-------------------------------------------------------")
	solvePart2(seatPlan)
}

func readSeatPlan(filename string) []string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	rows := strings.Split(string(buf), "\r\n")

	return rows
}

func solvePart1(seatPlan []string) {
	plans := [2]SeatPlan{MakeSeatPlan(seatPlan), MakeSeatPlan(seatPlan)}
	activePlan := 0

	changeCount := -1
	iterationCount := 0
	for changeCount != 0 {
		iterationCount++
		changeCount = generalIteration(&plans[activePlan], &plans[1-activePlan], countOccupiedNeighbors, 4) // stepIteration(&plans[activePlan], &plans[1-activePlan])
		activePlan = 1 - activePlan
		fmt.Printf("  ChangeCount: %d\n", changeCount)
		fmt.Printf("Active:\n%v\n", plans[activePlan])
		fmt.Printf("Inactive:\n%v\n", plans[1-activePlan])
	}
	result := plans[activePlan].CountOccupiedSeats()

	fmt.Printf("Stabilized after %d iterations\n", iterationCount)
	fmt.Printf("Number of occupied seats: %d\n", result)
}

func countOccupiedNeighbors(plan [][]byte, row int, col int) int {
	res := 0

	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			if i != 0 || j != 0 {
				if plan[row+i][col+j] == '#' {
					res++
				}
			}
		}
	}

	return res
}

func countOccupiedVisibleSeats(plan [][]byte, row int, col int) int {
	res := 0

	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}
			x := row
			y := col
			reached := false
			for !reached {
				x = x + i
				y = y + j
				switch plan[x][y] {
				case 'X':
					reached = true
				case 'L':
					reached = true
				case '#':
					res++
					reached = true
				}
			}
		}
	}

	return res
}

func solvePart2(seatPlan []string) {
	plans := [2]SeatPlan{MakeSeatPlan(seatPlan), MakeSeatPlan(seatPlan)}
	activePlan := 0

	changeCount := -1
	iterationCount := 0
	for changeCount != 0 {
		iterationCount++
		changeCount = generalIteration(&plans[activePlan], &plans[1-activePlan], countOccupiedVisibleSeats, 5) // stepIteration(&plans[activePlan], &plans[1-activePlan])
		activePlan = 1 - activePlan
		fmt.Printf("  ChangeCount: %d\n", changeCount)
		fmt.Printf("Active:\n%v\n", plans[activePlan])
		fmt.Printf("Inactive:\n%v\n", plans[1-activePlan])
	}
	result := plans[activePlan].CountOccupiedSeats()

	fmt.Printf("Stabilized after %d iterations\n", iterationCount)
	fmt.Printf("Number of occupied seats: %d\n", result)
}

func generalIteration(from *SeatPlan, to *SeatPlan, countOccupation func([][]byte, int, int) int, emptyOccupiedSeatThreshold int) int {
	changeCount := 0

	for r := 1; r < from.Height+1; r++ {
		for c := 1; c < from.Width+1; c++ {
			switch from.Plan[r][c] {
			case 'L':
				occupiedCount := countOccupation(from.Plan, r, c)
				if occupiedCount == 0 {
					to.Plan[r][c] = '#'
					changeCount++
				} else {
					to.Plan[r][c] = from.Plan[r][c]
				}
			case '#':
				occupiedCount := countOccupation(from.Plan, r, c)
				if occupiedCount >= emptyOccupiedSeatThreshold {
					to.Plan[r][c] = 'L'
					changeCount++
				} else {
					to.Plan[r][c] = from.Plan[r][c]
				}
			default:
				to.Plan[r][c] = from.Plan[r][c]
			}
		}
	}

	return changeCount
}
