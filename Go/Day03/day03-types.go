package main

import "fmt"

type Slope struct {
	X int
	Y int
}

func (slope Slope) String() string {
	return fmt.Sprintf("Slope(%d/%d)", slope.X, slope.Y)
}

type WanderingPosition struct {
	Step  Slope
	XWrap int
	X     int
	Y     int
}

func (wp *WanderingPosition) step() {
	wp.X = (wp.X + wp.Step.X) % wp.XWrap
	wp.Y += wp.Step.Y
}
