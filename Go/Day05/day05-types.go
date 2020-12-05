package main

import "fmt"

type BoardingPass struct {
	Address string
	SeatId  int
}

func (bp BoardingPass) String() string {
	return fmt.Sprintf("BoardingPass(%s -> %d)", bp.Address, bp.SeatId)
}

func MakeBoardingPass(address string) BoardingPass {
	res := BoardingPass{
		Address: address,
		SeatId:  calcSeatId(address),
	}

	return res
}

func calcSeatId(boardingPass string) int {
	f, b := 0, 127
	for i := 0; i < 7; i++ {
		m := (f + b) / 2
		if boardingPass[i] == 'F' {
			b = m
		} else {
			f = m + 1
		}
	}
	row := f

	f, b = 0, 8
	for i := 7; i < 10; i++ {
		m := (f + b) / 2
		if boardingPass[i] == 'L' {
			b = m
		} else {
			f = m
		}
	}
	col := f

	return row*8 + col
}
