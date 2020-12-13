package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 01: Report Repair\n=====================")

	earliestDepartTime, busIds := readTimeTable("RawData.txt")
	fmt.Printf("Earliest depart time: %d\n", earliestDepartTime)
	fmt.Printf("Bus IDs             : %v\n", busIds)

	fmt.Println("\nPart 1: Find earliest bus, id * wait minutes\n-----------------------------------------------------")
	solvePart1(earliestDepartTime, busIds)

	fmt.Println("\nPart 2: Three numbers who add up to 2020, their product\n-------------------------------------------------------")
	solvePart2(busIds)
}

func readTimeTable(filename string) (int, []int) {
	// 1st read the file content
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	lines := strings.Split(string(buf), "\r\n")

	// 2nd parse single-value first line, earliest departure time
	earliestDepartTime, err := strconv.Atoi(lines[0])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", lines[0], err)
		os.Exit(1)
	}

	// 3rd parse from comma-separated second line the bus id's
	var busIds []int
	for _, n := range strings.Split(lines[1], ",") {
		if n == "x" {
			busIds = append(busIds, -1)
			continue
		}
		i, err := strconv.Atoi(n)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", n, err)
			continue
		}
		busIds = append(busIds, i)
	}

	return earliestDepartTime, busIds
}

func solvePart1(earliestDepartTime int, busIds []int) {
	busId, departTime := 0, 0
	found := false
	for dt := earliestDepartTime; !found; dt++ {
		for _, id := range busIds {
			if id == -1 {
				continue
			}
			if dt%id == 0 {
				busId = id
				departTime = dt
				found = true
			}
		}
	}

	fmt.Printf("Found bus and depart time: %d at %d\n", busId, departTime)
	minutesToWait := departTime - earliestDepartTime
	fmt.Printf("Minutes to wait: %d\n", minutesToWait)
	fmt.Printf("Result for part 1: %d\n", busId*minutesToWait)
}

// solve part 2 using Chinese Remainder Theorem, like https://de.wikipedia.org/wiki/Chinesischer_Restsatz#Finden_einer_L%C3%B6sung
func solvePart2(busIds []int) {
	M := big.NewInt(productOfAllPositives(busIds))

	sumOfAllXatI := big.NewInt(0)
	for i, id := range busIds {
		if id == -1 {
			continue
		}
		mAtI := big.NewInt(int64(id))
		MwithoutmAtI := new(big.Int).Div(M, mAtI)
		z, rAtI, sAtI, eAtI := big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)
		z.GCD(rAtI, sAtI, mAtI, MwithoutmAtI)
		eAtI.Mul(sAtI, MwithoutmAtI)
		xAtI, aAtI := big.NewInt(0), big.NewInt(int64(-i))
		xAtI.Mul(eAtI, aAtI)
		sumOfAllXatI.Add(sumOfAllXatI, xAtI)
	}

	result := new(big.Int).Mod(sumOfAllXatI, M)
	fmt.Printf("Earliest departure timestamp: %v\n", result)
}

func productOfAllPositives(numbers []int) int64 {
	res := int64(1)
	for _, n := range numbers {
		if n < 1 {
			continue
		}
		res = res * int64(n)
	}

	return res
}
