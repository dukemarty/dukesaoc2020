package main

import "fmt"

type Registers struct {
	IP  int
	ACC int
}

func (r Registers) String() string {
	return fmt.Sprintf("IP:%d, ACC:%d", r.IP, r.ACC)
}

type ProgramEndCode int

const (
	UNKNOWN      ProgramEndCode = 0
	FINISHED     ProgramEndCode = 1
	INFINITELOOP ProgramEndCode = 2
)

type Machine struct {
	Regs    Registers
	Program []Instruction
}

func (mach Machine) String() string {
	return fmt.Sprintf("Machine(Registers: %s)", mach.Regs)
}

func (mach Machine) CurrentInstruction() Instruction {
	return mach.Program[mach.Regs.IP]
}

func (mach *Machine) Reset() {
	mach.Regs.ACC = 0
	mach.Regs.IP = 0
}

func (mach *Machine) Step() int {
	ci := mach.CurrentInstruction()
	switch ci.Op {
	case NOP:
		mach.Regs.IP++
	case ACC:
		mach.Regs.ACC = mach.Regs.ACC + ci.Arg
		mach.Regs.IP++
	case JMP:
		mach.Regs.IP = mach.Regs.IP + ci.Arg
	}

	return mach.Regs.IP
}

func (mach *Machine) Run() ProgramEndCode {
	alreadyRunSet := make(map[int]bool)
	for mach.Regs.IP < len(mach.Program) && !alreadyRunSet[mach.Regs.IP] {
		alreadyRunSet[mach.Regs.IP] = true
		mach.Step()
	}

	var res ProgramEndCode
	switch {
	case mach.Regs.IP >= len(mach.Program):
		res = FINISHED
	case alreadyRunSet[mach.Regs.IP]:
		res = INFINITELOOP
	default:
		res = UNKNOWN
	}

	return res
}

func MakeMachine(program []Instruction) Machine {
	return Machine{
		Regs:    Registers{},
		Program: program,
	}
}
