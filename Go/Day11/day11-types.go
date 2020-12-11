package main

import "strings"

type SeatPlan struct {
	Height int
	Width  int
	Plan   [][]byte
}

func (sp SeatPlan) String() string {
	rows := make([]string, 0, sp.Height+2)
	for _, r := range sp.Plan {
		rows = append(rows, string(r))
	}

	return strings.Join(rows, "\n")
}

func (sp SeatPlan) CountOccupiedSeats() int {
	res := 0

	for _, row := range sp.Plan {
		for _, c := range row {
			if c == '#' {
				res++
			}
		}
	}

	return res
}

func MakeSeatPlan(rawPlan []string) SeatPlan {
	resPlan := SeatPlan{
		Height: len(rawPlan),
		Width:  len(rawPlan[0]),
	}

	upperLowerBoundary := make([]byte, resPlan.Width+2)
	for i := range upperLowerBoundary {
		upperLowerBoundary[i] = 'X'
	}
	resPlan.Plan = append(resPlan.Plan, upperLowerBoundary)
	for _, r := range rawPlan {
		row := make([]byte, len(rawPlan[0])+2)
		row[0] = 'X'
		copy(row[1:], r)
		row[len(row)-1] = 'X'
		resPlan.Plan = append(resPlan.Plan, row)
	}
	resPlan.Plan = append(resPlan.Plan, upperLowerBoundary)

	return resPlan
}
