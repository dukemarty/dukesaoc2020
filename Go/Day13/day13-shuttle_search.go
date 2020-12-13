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
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	lines := strings.Split(string(buf), "\r\n")

	earliestDepartTime, err := strconv.Atoi(lines[0])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", lines[0], err)
		os.Exit(1)
	}

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
	M := int64(1)
	for _, id := range busIds {
		if id == -1 {
			continue
		}
		M = M * int64(id)
	}
	fmt.Printf("M = %v\n", M)
	x := big.NewInt(0)
	for i, id := range busIds {
		if id == -1 {
			continue
		}
		m_i, M_i := big.NewInt(int64(id)), big.NewInt(M/int64(id))
		z, r, s, e_i := big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)
		z.GCD(r, s, m_i, M_i)
		fmt.Printf("%v = %v*%v + %v*%v (z=r*m_i+s*M_i)\n", z, r, m_i, s, M_i)
		e_i.Mul(s, M_i)
		fmt.Printf("  e_%d = s*M_%d = %v*%v = %v\n", i, i, s, M_i, e_i)
		x_i := big.NewInt(0)
		a_i := big.NewInt(int64(id - i))
		x_i.Mul(e_i, a_i)
		fmt.Printf("  x_%d = %v\n", i, x_i)
		x.Add(x, x_i)
	}

	result, bigM := big.NewInt(0), big.NewInt(M)
	result = result.Mod(x, bigM)

	fmt.Printf("Earliest departure timestamp: %v\n", result)
}
