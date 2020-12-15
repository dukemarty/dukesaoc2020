package main

type Machine struct {
	Mask   Mask
	Memory map[uint64]uint64
}

func MakeMachine() Machine {
	res := Machine{
		Memory: make(map[uint64]uint64),
	}

	return res
}

type Interpreter struct {
	Machine Machine
	Program []Instruction
}

func (interpreter *Interpreter) RunWithVersion1() {
	for _, inst := range interpreter.Program {
		switch inst.GetType() {
		case MASK:
			inst.ApplyAsVersion1(&interpreter.Machine)
		case MEM:
			inst.ApplyAsVersion1(&interpreter.Machine)
		}
	}
}

func (interpreter *Interpreter) RunWithVersion2() {
	for _, inst := range interpreter.Program {
		switch inst.GetType() {
		case MASK:
			inst.ApplyAsVersion2(&interpreter.Machine)
		case MEM:
			inst.ApplyAsVersion2(&interpreter.Machine)
		}
	}
}

func (i Interpreter) SumMemory() uint64 {
	res := uint64(0)

	for _, value := range i.Machine.Memory {
		res = res + value
	}

	return res
}

func MakeInterpreter(program []Instruction) Interpreter {
	res := Interpreter{
		Machine: MakeMachine(),
		Program: program,
	}

	return res
}
