package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 09: Encoding Error\n=====================")

	xmasStream := readXmasStream("RawData.txt")
	fmt.Printf("Read numbers: %q\n", xmasStream)

	fmt.Println("\nPart 1: Find first invalid number\n-----------------------------------------------------")
	weaknessNumber := solvePart1(xmasStream)

	fmt.Println("\nPart 2: Find elements to build weakness number\n-------------------------------------------------------")
	solvePart2(weaknessNumber, xmasStream)
}

func readXmasStream(filename string) []int {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	numberStrings := strings.Split(string(buf), "\r\n")
	//fmt.Printf("%q\n", numberStrings)
	//fmt.Println(strings.Join(numberStrings, ","))

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

func solvePart1(numbers []int) int {
	const windowSize = 25
	var res int
	for i, n := range numbers[windowSize:] {
		if !isSumOfTwo(n, numbers[i:windowSize+i]) {
			res = n
			break
		}
	}

	fmt.Printf("Found number which can not be combined from to others: %d", res)
	return res
}

func isSumOfTwo(target int, candidates []int) bool {
	for i, m := range candidates {
		if m < target {
			for _, n := range candidates[i+1:] {
				if m+n == target {
					return true
				}
			}
		}
	}

	return false
}

func solvePart2(weaknessNumber int, numbers []int) {
	found := false
	var begin int
	var end int
	for i := range numbers {
		sum := 0
		for j, n := range numbers[i:] {
			sum = sum + n
			if sum == weaknessNumber {
				begin = i
				end = i + j
				found = true
				break
			} else if sum > weaknessNumber {
				break
			}
		}
		if found {
			break
		}
	}

	resultSlice := numbers[begin : end+1]
	fmt.Printf("Found range to sum to weakness number: %q\n", resultSlice)
	summedSlice := 0
	for _, n := range resultSlice {
		summedSlice = summedSlice + n
	}
	fmt.Printf("    Sum: %d\n", summedSlice)
	smallest := minInSlice(resultSlice)
	largest := maxInSlice(resultSlice)
	fmt.Printf("Found smallest/largest: %d / %d\n", smallest, largest)
	fmt.Printf("Resulting encryption weakness: %d\n", smallest+largest)
}

func minInSlice(slice []int) int {
	var res int
	for i, n := range slice {
		if i == 0 || n < res {
			res = n
		}
	}

	return res
}

func maxInSlice(slice []int) int {
	var res int
	for i, n := range slice {
		if i == 0 || n > res {
			res = n
		}
	}

	return res
}
