package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 01: Docking Data\n=====================")

	program := readInitializationProgram("RawData.txt")
	fmt.Printf("Program: %v\n", program)

	fmt.Println("\nPart 1: Run as V1 protocol and sum all values\n----------------------------------")
	solvePart1(program)

	fmt.Println("\nPart 2: Run as V2 protocol and sum all values\n-------------------------------------------------------")
	solvePart2(program)
}

func readInitializationProgram(filename string) []Instruction {
	// 1st read the file content
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	var program []Instruction
	for _, line := range strings.Split(string(buf), "\r\n") {
		fields := strings.Fields(line)
		if fields[0] == "mask" {
			inst := MakeMaskInstruction(fields[2])
			program = append(program, inst)
		} else {
			inst, err := MakeMemInstruction(fields[0], fields[2])
			if err != nil {
				continue
			}
			program = append(program, inst)
		}
	}

	return program
}

func solvePart1(program []Instruction) {
	interpreter := MakeInterpreter(program)

	interpreter.RunWithVersion1()
	fmt.Printf("Interpreter: %v\n", interpreter)

	res := interpreter.SumMemory()
	fmt.Printf("Result for part 1: %d\n", res)
}

func solvePart2(program []Instruction) {
	interpreter := MakeInterpreter(program)

	interpreter.RunWithVersion2()
	// fmt.Printf("Interpreter: %v\n", interpreter)

	res := interpreter.SumMemory()
	fmt.Printf("Result for part 2: %d\n", res)
}
