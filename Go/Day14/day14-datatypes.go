package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type InstructionType string

const (
	INVALID InstructionType = "INVALID"
	MASK    InstructionType = "MASK"
	MEM     InstructionType = "MEM"
)

type Instruction interface {
	GetType() InstructionType
	ApplyAsVersion1(m *Machine)
	ApplyAsVersion2(m *Machine)
}

type Mask struct {
	AndMask      uint64
	OrMask       uint64
	PositionsOfX []int
}

func (mask Mask) Apply(value uint64) uint64 {
	res := (value & mask.AndMask) | mask.OrMask

	return res
}

type MaskInstruction struct {
	RawMask [36]byte
	Mask    Mask
}

func (mi MaskInstruction) GetType() InstructionType {
	return MASK
}

func (mi MaskInstruction) ApplyAsVersion1(m *Machine) {
	m.Mask = mi.Mask
}

func (mi MaskInstruction) ApplyAsVersion2(m *Machine) {
	m.Mask = mi.Mask
}

func (mi MaskInstruction) String() string {
	bs := mi.RawMask[:]
	return fmt.Sprintf("Mask <%s>: AND=%b, OR=%b, Xs at %v", string(bs), mi.Mask.AndMask, mi.Mask.OrMask, mi.Mask.PositionsOfX)
}

func MakeMaskInstruction(rawMask string) MaskInstruction {
	mask := []byte(rawMask)

	res := MaskInstruction{}

	for i := 0; i < 36; i++ {
		res.RawMask[i] = mask[i]
		res.Mask.AndMask = res.Mask.AndMask << 1
		res.Mask.OrMask = res.Mask.OrMask << 1
		switch mask[i] {
		case '0':
			// res.AndMask least bit is kept as 0
			// res.OrMask least bit is kept as 0
		case '1':
			res.Mask.AndMask = res.Mask.AndMask | 1
			res.Mask.OrMask = res.Mask.OrMask | 1
		case 'X':
			res.Mask.AndMask = res.Mask.AndMask | 1
			// res.OrMask least bit is kept as 0
			res.Mask.PositionsOfX = append(res.Mask.PositionsOfX, 35-i)
		}
	}

	return res
}

type MemInstruction struct {
	Address uint64
	Value   uint64
}

func (mi MemInstruction) ApplyAsVersion1(m *Machine) {
	m.Memory[mi.Address] = m.Mask.Apply(mi.Value)
}

func (mi MemInstruction) ApplyAsVersion2(m *Machine) {
	fmt.Printf("Basic address for writing: %d = %b\n", mi.Address, mi.Address)
	fmt.Printf("  Mask to apply: %v\n", m.Mask)
	for i := 0; i < 1<<len(m.Mask.PositionsOfX); i++ {
		oneMask := uint64(0)
		zeroMask := uint64(0)
		x := i
		for j := 0; j < len(m.Mask.PositionsOfX); j++ {
			if x%2 == 0 {
				zeroMask = zeroMask | 1<<m.Mask.PositionsOfX[j]
			} else {
				oneMask = oneMask | 1<<m.Mask.PositionsOfX[j]
			}
			x = x >> 1
		}
		address := ((mi.Address | m.Mask.OrMask) | oneMask) &^ zeroMask
		fmt.Printf("    Address %d: %d = %b\n", i, address, address)
		m.Memory[address] = mi.Value
	}
}

func (mi MemInstruction) GetType() InstructionType {
	return MEM
}

func (mi MemInstruction) String() string {
	return fmt.Sprintf("Mem: [%d] <- %d", mi.Address, mi.Value)
}

func MakeMemInstruction(rawAddressPart string, value string) (MemInstruction, error) {
	res := MemInstruction{}

	re, _ := regexp.Compile(`^mem\[(?P<address>\d+)\]$`)
	match := re.FindStringSubmatch(rawAddressPart)
	address, err := strconv.ParseUint(match[1], 10, 64)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", match[1], err)
		return res, errors.New("Invalid address for MemInstruction.")
	}
	res.Address = address

	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", value, err)
		return res, errors.New("Invalid value for MemInstruction.")
	}
	res.Value = i

	return res, nil
}
