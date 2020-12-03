package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 03: Toboggan Trajectory\n===========================")

	localGeology := readGeology("RawData.txt")
	fmt.Printf("Local geology:\n%s\n", strings.Join(localGeology, "\n"))

	fmt.Println("\nPart 1: Count trees on the way\n----------------------------------------")
	solvePart1(localGeology)

	fmt.Println("\nPart 2: Count trees on different slopes, their product\n-------------------------------------------------------")
	solvePart2(localGeology)
}

func readGeology(filename string) []string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	rows := strings.Split(string(buf), "\r\n")

	return rows
}

type Slope struct {
	X int
	Y int
}

func (slope Slope) String() string {
	return fmt.Sprintf("Slope(%d/%d)", slope.X, slope.Y)
}

type WanderingPosition struct {
	Step  Slope
	XWrap int
	X     int
	Y     int
}

func (wp *WanderingPosition) step() {
	wp.X = (wp.X + wp.Step.X) % wp.XWrap
	wp.Y += wp.Step.Y
}

func solvePart1(geology []string) {
	treeCount := checkSlope(geology, Slope{3, 1})
	fmt.Printf("Counted trees: %d\n", treeCount)
}

func solvePart2(geology []string) {
	var slopes = []Slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	res := 1
	for _, s := range slopes {
		treeCount := checkSlope(geology, s)
		fmt.Printf("    Tree count for slope %s: %d\n", s, treeCount)
		res *= treeCount
	}

	fmt.Printf("Multiplied tree counts: %d\n", res)
}

// returns tree count
func checkSlope(geology []string, step Slope) int {
	pos := WanderingPosition{
		Step:  step,
		XWrap: len(geology[0]),
		X:     0,
		Y:     0}

	treeCount := 0
	for ; pos.Y < len(geology); pos.step() {
		if geology[pos.Y][pos.X] == '#' {
			treeCount++
		}
	}

	return treeCount
}
