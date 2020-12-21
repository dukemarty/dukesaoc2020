package main

import (
	"fmt"
	"strings"
)

type EnergySource4D struct {
	Planes [][]ConwayCubePlane
}

func (es EnergySource4D) String() string {
	blocks := make([]string, 0, len(es.Planes))
	for hpi, hp := range es.Planes {
		for pi, p := range hp {
			blocks = append(blocks, fmt.Sprintf("Plane at %d/%d", hpi, pi))
			blocks = append(blocks, p.String())
		}
	}

	return strings.Join(blocks, "\n\n")
}

func (es EnergySource4D) CountActiveCubes() int {
	res := 0

	for _, hyperplane := range es.Planes {
		for _, plane := range hyperplane {
			res = res + plane.CountActiveCubes()
		}
	}

	return res
}

func (es *EnergySource4D) ExecuteCycle() {
	fmt.Printf("Length of es.Planes: %d\n", len(es.Planes))
	nextPlanes := make([][]ConwayCubePlane, 0, len(es.Planes))
	nextPlanes = append(nextPlanes, es.Planes[0])
	for hpi := 1; hpi < len(es.Planes)-1; hpi++ {
		hyperplane := make([]ConwayCubePlane, 0, len(es.Planes))
		hyperplane = append(hyperplane, es.Planes[hpi][0])
		for pi := 1; pi < len(es.Planes[hpi])-1; pi++ {
			plane := es.ComputeNextCycleState(hpi, pi)
			hyperplane = append(hyperplane, plane)
		}
		hyperplane = append(hyperplane, es.Planes[hpi][len(es.Planes[hpi])-1])
		nextPlanes = append(nextPlanes, hyperplane)
	}
	nextPlanes = append(nextPlanes, es.Planes[len(es.Planes)-1])
	es.Planes = nextPlanes
}

func (es *EnergySource4D) ComputeNextCycleState(hyperplaneIndex int, planeIndex int) ConwayCubePlane {
	//fmt.Printf("  ComputeNextCycleState(%d/%d)\n", hyperplaneIndex, planeIndex)
	ccp := es.Planes[hyperplaneIndex][planeIndex]
	resPlane := MakeEmptyConwayCubePlane(ccp.Height, ccp.Width)
	for r := 1; r < ccp.Height+1; r++ {
		for c := 1; c < ccp.Width+1; c++ {
			switch ccp.Cubes[r][c] {
			case ACTIVE:
				count := es.CountActiveNeighbours(hyperplaneIndex, planeIndex, r, c)
				if count == 2 || count == 3 {
					//fmt.Printf("Change cube (1) %d/%d/%d\n", planeIndex, r, c)
					resPlane.Cubes[r][c] = ACTIVE
				} else {
					//fmt.Printf("Change cube (2) %d/%d/%d\n", planeIndex, r, c)
					resPlane.Cubes[r][c] = INACTIVE
				}
			case INACTIVE:
				count := es.CountActiveNeighbours(hyperplaneIndex, planeIndex, r, c)
				if count == 3 {
					//fmt.Printf("Change cube (3) %d/%d/%d\n", planeIndex, r, c)
					resPlane.Cubes[r][c] = ACTIVE
				} else {
					//fmt.Printf("Change cube (4) %d/%d/%d\n", planeIndex, r, c)
					resPlane.Cubes[r][c] = INACTIVE
				}
			default:
				resPlane.Cubes[r][c] = ccp.Cubes[r][c]
			}
		}
	}

	return resPlane
}

func (es EnergySource4D) CountActiveNeighbours(hyperplaneIndex int, planeIndex int, rowIndex int, columnIndex int) int {
	res := 0

	for _, dh := range [...]int{-1, 0, 1} {
		for _, dp := range [...]int{-1, 0, 1} {
			for _, dr := range [...]int{-1, 0, 1} {
				for _, dc := range [...]int{-1, 0, 1} {
					if dh == 0 && dp == 0 && dr == 0 && dc == 0 {
						continue
					}
					if es.Planes[hyperplaneIndex+dh][planeIndex+dp].Cubes[rowIndex+dr][columnIndex+dc] == ACTIVE {
						res++
					}
				}
			}
		}
	}

	return res
}

func MakeEnergySource4D(initialState []string, intendedRoundCount int) EnergySource4D {
	resESource := EnergySource4D{
		Planes: make([][]ConwayCubePlane, 0, 2*intendedRoundCount+3),
	}
	fmt.Printf("  Length of resESource.Planes: %d\n", len(resESource.Planes))

	// 1st create hyperplane slices
	for i := 0; i < 2*intendedRoundCount+3; i++ {
		hp := make([]ConwayCubePlane, 0, 2*intendedRoundCount+3)
		resESource.Planes = append(resESource.Planes, hp)
	}

	// 2nd fill half of the hyperplanes with empty planes
	for i := 0; i < intendedRoundCount+1; i++ {
		for j := 0; j < 2*intendedRoundCount+3; j++ {
			resESource.Planes[i] = append(resESource.Planes[i], MakeEmptyConwayCubePlane(len(initialState)+2*intendedRoundCount, len(initialState[0])+2*intendedRoundCount))
		}
	}

	// 3rd middle hyperplane has the initial state in the middle
	for i := 0; i < intendedRoundCount+1; i++ {
		resESource.Planes[intendedRoundCount+1] = append(resESource.Planes[intendedRoundCount+1], MakeEmptyConwayCubePlane(len(initialState)+2*intendedRoundCount, len(initialState[0])+2*intendedRoundCount))
	}
	fmt.Printf("  Length of resESource.Planes: %d\n", len(resESource.Planes))
	resESource.Planes[intendedRoundCount+1] = append(resESource.Planes[intendedRoundCount+1], MakeConwayCubePlane(initialState, intendedRoundCount))
	fmt.Printf("  Length of resESource.Planes: %d\n", len(resESource.Planes))
	for i := 0; i < intendedRoundCount+1; i++ {
		resESource.Planes[intendedRoundCount+1] = append(resESource.Planes[intendedRoundCount+1], MakeEmptyConwayCubePlane(len(initialState)+2*intendedRoundCount, len(initialState[0])+2*intendedRoundCount))
	}
	fmt.Printf("  Length of resESource.Planes: %d\n", len(resESource.Planes))

	// 4th fill the second half of the hyperplanes with empty planes
	for i := intendedRoundCount + 2; i < 2*intendedRoundCount+3; i++ {
		for j := 0; j < 2*intendedRoundCount+3; j++ {
			resESource.Planes[i] = append(resESource.Planes[i], MakeEmptyConwayCubePlane(len(initialState)+2*intendedRoundCount, len(initialState[0])+2*intendedRoundCount))
		}
	}

	return resESource
}
