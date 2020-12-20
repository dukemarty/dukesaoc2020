package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 17: Conway Cubes\n===========================")

	initialState := readInitialState("RawData.txt")
	fmt.Printf("Initial seat plan layout:\n%s\n", strings.Join(initialState, "\n"))

	fmt.Println("\nPart 1: Count active cubes\n--------------------------------")
	solvePart1(initialState)

	//fmt.Println("\nPart 2: Count seats after stabilization with visibility\n-------------------------------------------------------")
	//iterationCount2, result2 := solve(seatPlan, func(plan [][]byte, r int, c int) int { return countRelevantOccupiedSeats(plan, r, c, math.MaxInt32) }, 5)
	//fmt.Printf("Stabilized after %d iterations\n", iterationCount2)
	//fmt.Printf("Number of occupied seats: %d\n", result2)
}

func readInitialState(filename string) []string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	rows := strings.Split(string(buf), "\r\n")

	return rows
}

func solvePart1(initialState []string) {
	energySource := MakeEnergySource(initialState, 8)
	fmt.Println(energySource)
	fmt.Printf("Length of es.Planes: %d\n", len(energySource.Planes))

	for i := 0; i < 6; i++ {
		energySource.ExecuteCycle()
		//fmt.Println("=============================================")
		//fmt.Println(energySource)
		//fmt.Printf("Number of active cubes with i=%d: %d\n", i, energySource.CountActiveCubes())
	}

	res := energySource.CountActiveCubes()
	fmt.Printf("Active cubes after 6 rounds: %d\n", res)
}

//// Returns number of iterations till stabilization and final count of occupied seats.
//func solve(seatPlan []string, countOccupation func([][]byte, int, int) int, emptyOccupiedSeatThreshold int) (int, int) {
//	plans := [2]SeatPlan{MakeSeatPlan(seatPlan), MakeSeatPlan(seatPlan)}
//	activePlan := 0
//
//	changeCount := -1
//	iterationCount := 0
//	for changeCount != 0 {
//		iterationCount++
//		changeCount = generalIteration(&plans[activePlan], &plans[1-activePlan], countOccupation, emptyOccupiedSeatThreshold)
//		activePlan = 1 - activePlan
//	}
//	result := plans[activePlan].CountOccupiedSeats()
//
//	return iterationCount, result
//}
//
//func countRelevantOccupiedSeats(plan [][]byte, row int, col int, maxDist int) int {
//	res := 0
//
//	for _, dx := range []int{-1, 0, 1} {
//		for _, dy := range []int{-1, 0, 1} {
//			if dx == 0 && dy == 0 {
//				continue
//			}
//
//			x, y, reached := row, col, false
//			dist := 0
//			for !reached && dist < maxDist {
//				x, y = x+dx, y+dy
//				switch plan[x][y] {
//				case 'X', 'L':
//					reached = true
//				case '#':
//					res++
//					reached = true
//				}
//				dist++
//			}
//		}
//	}
//
//	return res
//}
//
//func generalIteration(from *SeatPlan, to *SeatPlan, countOccupation func([][]byte, int, int) int, emptyOccupiedSeatThreshold int) int {
//	changeCount := 0
//
//	for r := 1; r < from.Height+1; r++ {
//		for c := 1; c < from.Width+1; c++ {
//			switch from.Plan[r][c] {
//			case 'L':
//				occupiedCount := countOccupation(from.Plan, r, c)
//				if occupiedCount == 0 {
//					to.Plan[r][c] = '#'
//					changeCount++
//				} else {
//					to.Plan[r][c] = from.Plan[r][c]
//				}
//			case '#':
//				occupiedCount := countOccupation(from.Plan, r, c)
//				if occupiedCount >= emptyOccupiedSeatThreshold {
//					to.Plan[r][c] = 'L'
//					changeCount++
//				} else {
//					to.Plan[r][c] = from.Plan[r][c]
//				}
//			default:
//				to.Plan[r][c] = from.Plan[r][c]
//			}
//		}
//	}
//
//	return changeCount
//}
