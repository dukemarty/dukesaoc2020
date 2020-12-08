package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 08: Handheld Halting\n========================")

	programCode := readProgramCode("RawData.txt")
	fmt.Printf("Read code: %q\n", programCode)

	fmt.Println("\nPart 1: Determine ACC value at 1st looping\n------------------------------------------")
	solvePart1(programCode)

	fmt.Println("\nPart 2: Try to fix bootloader\n-------------------------------------")
	solvePart2(programCode)
}

func readProgramCode(filename string) []Instruction {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	var res []Instruction
	for _, line := range strings.Split(string(buf), "\r\n") {
		inst := MakeInstruction(line)
		if inst.Op != INVALID {
			res = append(res, inst)
		}
	}

	return res
}

func solvePart1(program []Instruction) {
	machine := MakeMachine(program)

	endcode := machine.Run()
	fmt.Printf("%s ->%d\n", machine, endcode)
	fmt.Printf("Final value in ACC: %d\n", machine.Regs.ACC)
}

func solvePart2(program []Instruction) {
	machine := MakeMachine(program)

	for i := 0; i < len(program); i++ {
		if !toggleOperation(&program[i]) {
			continue
		}

		machine.Reset()
		endcode := machine.Run()

		if endcode == FINISHED {
			fmt.Printf("Operation to toggle: #%d\n", i)
			break
		}

		toggleOperation(&program[i])
	}
	fmt.Printf("%s\n", machine)

	fmt.Printf("Final ACC value after program termination: %d\n", machine.Regs.ACC)
}

func toggleOperation(inst *Instruction) bool {
	switch inst.Op {
	case JMP:
		inst.Op = NOP
	case NOP:
		inst.Op = JMP
	default:
		return false
	}
	return true
}
