package main

import "strings"

const (
	INACTIVE = '.'
	ACTIVE   = '#'
	IGNORED  = 'X'
)

type ConwayCubePlane struct {
	Height int
	Width  int
	Cubes  [][]byte
}

func (ccp ConwayCubePlane) CountActiveCubes() int {
	res := 0
	for _, r := range ccp.Cubes {
		for _, b := range r {
			if b == ACTIVE {
				res++
			}
		}
	}

	return res
}

func (ccp ConwayCubePlane) String() string {
	rows := make([]string, 0, ccp.Height+2)
	for _, r := range ccp.Cubes {
		rows = append(rows, string(r))
	}

	return strings.Join(rows, "\n")
}

func MakeEmptyConwayCubePlane(height int, width int) ConwayCubePlane {
	res := ConwayCubePlane{
		Height: height,
		Width:  width,
		Cubes:  make([][]byte, 0, height+2),
	}

	upperLowerBoundary := make([]byte, width+2)
	for i := range upperLowerBoundary {
		upperLowerBoundary[i] = IGNORED
	}
	res.Cubes = append(res.Cubes, upperLowerBoundary)

	for r := 0; r < res.Height; r++ {
		row := make([]byte, res.Width+2)
		row[0] = IGNORED
		for i := 0; i < res.Width; i++ {
			row[i+1] = INACTIVE
		}
		row[len(row)-1] = IGNORED
		res.Cubes = append(res.Cubes, row)
	}

	res.Cubes = append(res.Cubes, upperLowerBoundary)

	return res
}

func MakeConwayCubePlane(initialState []string, intendedRoundCount int) ConwayCubePlane {
	res := ConwayCubePlane{
		Height: len(initialState) + 2*intendedRoundCount,
		Width:  len(initialState[0]) + 2*intendedRoundCount,
		Cubes:  make([][]byte, 0, len(initialState)+2*intendedRoundCount+2),
	}

	// create upper boundary element
	upperLowerBoundary := make([]byte, res.Width+2)
	for i := range upperLowerBoundary {
		upperLowerBoundary[i] = IGNORED
	}
	// add upper boundary element to plane
	res.Cubes = append(res.Cubes, upperLowerBoundary)

	// add empty lines with left/right boundaries between upper boundary and copy of initial state
	for r := 0; r < intendedRoundCount; r++ {
		row := make([]byte, res.Width+2)
		row[0] = IGNORED
		for i := 0; i < res.Width; i++ {
			row[i+1] = INACTIVE
		}
		row[len(row)-1] = IGNORED
		res.Cubes = append(res.Cubes, row)
	}

	// add rows containing initial state
	for _, r := range initialState {
		row := make([]byte, res.Width+2)
		row[0] = IGNORED
		for i := 1; i < intendedRoundCount+1; i++ {
			row[i] = INACTIVE
		}
		copy(row[1+intendedRoundCount:], r)
		for i := 1 + intendedRoundCount + len(initialState[0]); i < len(row)-1; i++ {
			row[i] = INACTIVE
		}
		row[len(row)-1] = IGNORED
		res.Cubes = append(res.Cubes, row)
	}

	// add empty lines with left/right boundaries between copy of initial state and lower boundary
	for r := 0; r < intendedRoundCount; r++ {
		row := make([]byte, res.Width+2)
		row[0] = IGNORED
		for i := 0; i < res.Width; i++ {
			row[i+1] = INACTIVE
		}
		row[len(row)-1] = IGNORED
		res.Cubes = append(res.Cubes, row)
	}

	res.Cubes = append(res.Cubes, upperLowerBoundary)

	return res
}
