package main

import "fmt"

type NavInstruction struct {
	Command byte
	Value   int
}

func (ni NavInstruction) String() string {
	return fmt.Sprintf("(%c:%d)", ni.Command, ni.Value)
}

type WayPoint struct {
	X int
	Y int
}

func (wp *WayPoint) Move(inst NavInstruction) {
	rotLeftX1, rotLeftX2 := [4]int{1, 0, -1, 0}, [4]int{0, -1, 0, 1} // cos, -sin
	rotLeftY1, rotLeftY2 := [4]int{0, 1, 0, -1}, [4]int{1, 0, -1, 0} // sin, cos

	switch inst.Command {
	case 'N':
		wp.Y = wp.Y + inst.Value
	case 'E':
		wp.X = wp.X + inst.Value
	case 'S':
		wp.Y = wp.Y - inst.Value
	case 'W':
		wp.X = wp.X - inst.Value
	case 'R':
		rotAsL := 360 - inst.Value
		sinCosIndex := rotAsL / 90
		fmt.Printf("  Waypoint before applying %v with sinCosIndex=%d: %v\n", inst, sinCosIndex, *wp)
		newX := wp.X*rotLeftX1[sinCosIndex] + wp.Y*rotLeftX2[sinCosIndex]
		newY := wp.X*rotLeftY1[sinCosIndex] + wp.Y*rotLeftY2[sinCosIndex]
		wp.X, wp.Y = newX, newY
		fmt.Printf("  Waypoint after: %v\n", *wp)
	case 'L':
		sinCosIndex := inst.Value / 90
		fmt.Printf("  Waypoint before applying %v with sinCosIndex=%d: %v\n", inst, sinCosIndex, *wp)
		newX := wp.X*rotLeftX1[sinCosIndex] + wp.Y*rotLeftX2[sinCosIndex]
		newY := wp.X*rotLeftY1[sinCosIndex] + wp.Y*rotLeftY2[sinCosIndex]
		wp.X, wp.Y = newX, newY
		fmt.Printf("  Waypoint after: %v\n", *wp)
	}

}

// We work under the (reasonable) assumption that only multiples of 90° are used, so we will work around
// an explicit conversion from degree to rad.
func (wp WayPoint) rotate(inst NavInstruction) {

}
