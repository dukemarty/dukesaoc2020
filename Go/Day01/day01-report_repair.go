package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 01: Report Repair\n=====================")

	expenseNumbers := readExpenseReport("RawData.txt")
	fmt.Printf("Read numbers: %q\n", expenseNumbers)

	fmt.Println("\nPart 1: Two numbers who add up to 2020, their product\n-----------------------------------------------------")
	solvePart1(expenseNumbers)

	fmt.Println("\nPart 2: Three numbers who add up to 2020, their product\n-------------------------------------------------------")
	solvePart2(expenseNumbers)
}

func readExpenseReport(filename string) []int {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	numberStrings := strings.Split(string(buf), "\r\n")
	//fmt.Printf("%q\n", numberStrings)
	//fmt.Println(strings.Join(numberStrings, ","))

	var res []int
	for _, n := range numberStrings {
		i, err := strconv.Atoi(n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", n, err)
			continue
		}
		res = append(res, i)
	}

	return res
}

func solvePart1(numbers []int) {
	for i, n := range numbers {
		for _, m := range numbers[i+1:] {
			if m+n == 2020 {
				fmt.Printf("Found both numbers: %d + %d = %d\n", m, n, m+n)
				res := m * n
				fmt.Printf("Multiplication result: %d\n", res)
			}
		}
	}
}

func solvePart2(numbers []int) {
	for i, n := range numbers {
		for j, m := range numbers[i+1:] {
			for _, o := range numbers[i+j:] {
				if m+n+o == 2020 {
					fmt.Printf("Found a numbers: %d + %d + %d = %d\n", m, n, o, m+n+o)
					res := m * n * o
					fmt.Printf("Multiplication result: %d\n", res)
				}
			}
		}
	}
}
