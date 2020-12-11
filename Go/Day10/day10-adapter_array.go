package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 10: Adapter Array\n=====================")

	adapterJoltages := readAdapterJoltages("RawData.txt")
	fmt.Printf("Read %d joltages: %q\n", len(adapterJoltages), adapterJoltages)

	fmt.Println("\nPart 1: Multipy 1-jolt and 3-jolt differences\n-----------------------------------------------------")
	solvePart1(adapterJoltages)

	fmt.Println("\nPart 2: Number of possible adapter combinations\n-------------------------------------------------------")
	solvePart2(adapterJoltages)
}

func readAdapterJoltages(filename string) []int {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	numberStrings := strings.Split(string(buf), "\r\n")

	var res []int
	for _, n := range numberStrings {
		i, err := strconv.Atoi(n)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", n, err)
			continue
		}
		res = append(res, i)
	}

	return res
}

func solvePart1(adapters []int) {
	joltages := make([]int, len(adapters))
	copy(joltages, adapters)
	joltages = append(joltages, 0)
	sort.Ints(joltages)
	joltages = append(joltages, joltages[len(joltages)-1]+3)
	fmt.Printf("Joltages: %q\n", joltages)

	count1 := 0
	count3 := 0
	for i := 1; i < len(joltages); i++ {
		switch joltages[i] - joltages[i-1] {
		case 1:
			count1++
		case 3:
			count3++
		default:
			fmt.Printf("Compared %d and %d -> %d", joltages[i-1], joltages[i], joltages[i]-joltages[i-1])
		}
	}

	fmt.Printf("Counts: 1-jolts=%d, 3-jolts=%d\n", count1, count3)
	fmt.Printf("Multiplied result: %d\n", count1*count3)
}

type CombState struct {
	Adapter           int
	CombCountUpToHere int
	ReachedVia        []int
}

func (cs CombState) String() string {
	return fmt.Sprintf("[%d Joltage <= SUM=%d  (via %v)]", cs.Adapter, cs.CombCountUpToHere, cs.ReachedVia)
}

func solvePart2(adapters []int) {
	joltages := make([]int, len(adapters))
	copy(joltages, adapters)
	joltages = append(joltages, 0)
	sort.Ints(joltages)
	joltages = append(joltages, joltages[len(joltages)-1]+3)
	fmt.Printf("Joltages: %v\n", joltages)

	combinations := make([]CombState, len(joltages))
	for i := range combinations {
		combinations[i].Adapter = joltages[i]
	}
	combinations[0].CombCountUpToHere = 1

	for i := 0; i < len(joltages); i++ {
		if len(combinations[i].ReachedVia) > 0 {
			for _, rv := range combinations[i].ReachedVia {
				combinations[i].CombCountUpToHere += combinations[rv].CombCountUpToHere
			}
		}
		for d := 1; d < 4; d++ {
			if i+d < len(joltages) && joltages[i+d]-joltages[i] <= 3 {
				combinations[i+d].ReachedVia = append(combinations[i+d].ReachedVia, i)
			}
		}
	}

	for i, c := range combinations {
		fmt.Printf("%d: %v\n", i, c)
	}

	fmt.Printf("Possible combination for reaching adapter:: %d\n", combinations[len(combinations)-1].CombCountUpToHere)
}
