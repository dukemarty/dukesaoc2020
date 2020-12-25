package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 25: Combo Breaker\n=====================")

	pkCard, pkDoor := readPublicKeys("RawData.txt")
	fmt.Printf("Read public keys: %d  /  %d\n", pkCard, pkDoor)

	fmt.Println("\nPart 1: Compute encryption key\n-----------------------------------------------------")
	solvePart1(pkCard, pkDoor)

	//fmt.Println("\nPart 2: Manhattan distance to target via way point\n-------------------------------------------------------")
	//solvePart2(navigationInstructions)
}

func readPublicKeys(filename string) (int, int) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	lines := strings.Split(string(buf), "\r\n")

	pkCard, err := strconv.Atoi(lines[0])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", lines[0], err)
		return 0, 0
	}

	pkDoor, err := strconv.Atoi(lines[1])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", lines[1], err)
		return 0, 0
	}

	return pkCard, pkDoor
}

func solvePart1(pkCard int, pkDoor int) {

	loopSizeCard := findLoopSize(7, pkCard)
	loopSizeDoor := findLoopSize(7, pkDoor)

	fmt.Printf("Determined card loop size: %d\n", loopSizeCard)
	fmt.Printf("Determined door loop size: %d\n", loopSizeDoor)

	encryptionKeyViaCard := transformSubjectNumber(pkDoor, loopSizeCard)
	encryptionKeyViaDoor := transformSubjectNumber(pkCard, loopSizeDoor)

	fmt.Printf("Encryption key via card: %d\n", encryptionKeyViaCard)
	fmt.Printf("Encryption key via door: %d\n", encryptionKeyViaDoor)
}

func findLoopSize(subjectNumber int, publicKey int) int {
	value := 1

	loopSize := 0
	for value != publicKey {
		loopSize++
		value = value * subjectNumber
		value = value % 20201227
	}

	return loopSize
}

func transformSubjectNumber(subjectNumber int, loopSize int) int {
	value := 1

	for i := 0; i < loopSize; i++ {
		value = value * subjectNumber
		value = value % 20201227
	}

	return value
}

//func solvePart2(navInstructions []NavInstruction) {
//	x, y := 0, 0
//	wayPoint := WayPoint{X: 10, Y: 1}
//
//	for _, inst := range navInstructions {
//		switch inst.Command {
//		case 'N', 'E', 'S', 'W', 'R', 'L':
//			wayPoint.Move(inst)
//		case 'F':
//			x = x + wayPoint.X*inst.Value
//			y = y + wayPoint.Y*inst.Value
//		}
//	}
//
//	fmt.Printf("Final position: %d/%d (waypoint at %v)\n", x, y, wayPoint)
//	fmt.Printf("Distance to pos: %d\n", absInt(x)+absInt(y))
//}
