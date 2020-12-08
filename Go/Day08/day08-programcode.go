package main

import (
	"fmt"
	"strconv"
	"strings"
)

type HHCOperation string

const (
	INVALID HHCOperation = ""
	NOP     HHCOperation = "nop"
	ACC     HHCOperation = "acc"
	JMP     HHCOperation = "jmp"
)

type Instruction struct {
	Op  HHCOperation
	Arg int
}

func (inst Instruction) String() string {
	return fmt.Sprintf("[%s %d]", strings.ToUpper(string(inst.Op)), inst.Arg)
}

func MakeInstruction(line string) Instruction {
	tokens := strings.Fields(line)

	op := HHCOperation(tokens[0])
	argVal, err := strconv.Atoi(tokens[1])
	if err != nil {
		op = INVALID
		argVal = 0
	}

	res := Instruction{
		Op:  op,
		Arg: argVal,
	}

	return res
}
