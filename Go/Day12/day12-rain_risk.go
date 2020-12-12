package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 12: Rain Risk\n=====================")

	navigationInstructions := readNavInstructions("RawData.txt")
	fmt.Printf("Read instructions: %v\n", navigationInstructions)

	fmt.Println("\nPart 1: Manhattan distance to target\n-----------------------------------------------------")
	solvePart1(navigationInstructions)

	fmt.Println("\nPart 2: Manhattan distance to target via way point\n-------------------------------------------------------")
	solvePart2(navigationInstructions)
}

func readNavInstructions(filename string) []NavInstruction {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	instructions := strings.Split(string(buf), "\r\n")
	//fmt.Printf("%q\n", numberStrings)
	//fmt.Println(strings.Join(numberStrings, ","))

	var res []NavInstruction
	for _, n := range instructions {
		command := n[0]
		i, err := strconv.Atoi(n[1:])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", n, err)
			continue
		}
		res = append(res, NavInstruction{Command: command, Value: i})
	}

	return res
}

func solvePart1(navInstructions []NavInstruction) {
	x, y, dir := 0, 0, 0

	for _, inst := range navInstructions {
		switch inst.Command {
		case 'N':
			y = y - inst.Value
		case 'E':
			x = x + inst.Value
		case 'S':
			y = y + inst.Value
		case 'W':
			x = x - inst.Value
		case 'F':
			switch dir {
			case 0:
				x = x + inst.Value
			case 1:
				y = y + inst.Value
			case 2:
				x = x - inst.Value
			case 3:
				y = y - inst.Value
			}
		case 'R':
			dir = (dir + (inst.Value / 90) + 4) % 4
		case 'L':
			dir = (dir - (inst.Value / 90) + 4) % 4
		}
	}

	fmt.Printf("Final position: %d/%d, facing %d\n", x, y, dir)
	fmt.Printf("Distance to pos: %d\n", absInt(x)+absInt(y))
}

func solvePart2(navInstructions []NavInstruction) {
	x, y := 0, 0
	wayPoint := WayPoint{X: 10, Y: 1}

	for _, inst := range navInstructions {
		switch inst.Command {
		case 'N', 'E', 'S', 'W', 'R', 'L':
			wayPoint.Move(inst)
		case 'F':
			x = x + wayPoint.X*inst.Value
			y = y + wayPoint.Y*inst.Value
		}
	}

	fmt.Printf("Final position: %d/%d (waypoint at %v)\n", x, y, wayPoint)
	fmt.Printf("Distance to pos: %d\n", absInt(x)+absInt(y))
}

func absInt(number int) int {
	if number > 0 {
		return number
	} else {
		return -number
	}
}
