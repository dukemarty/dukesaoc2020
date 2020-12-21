package main

import (
	"fmt"
	"strings"
)

type EnergySource3D struct {
	Planes []ConwayCubePlane
}

func (es EnergySource3D) String() string {
	blocks := make([]string, 0, len(es.Planes))
	for _, p := range es.Planes {
		blocks = append(blocks, p.String())
	}

	return strings.Join(blocks, "\n\n")
}

func (es EnergySource3D) CountActiveCubes() int {
	res := 0

	for _, plane := range es.Planes {
		res = res + plane.CountActiveCubes()
	}

	return res
}

func (es *EnergySource3D) ExecuteCycle() {
	//fmt.Printf("Length of es.Planes: %d\n", len(es.Planes))
	nextPlanes := make([]ConwayCubePlane, 0, len(es.Planes))
	nextPlanes = append(nextPlanes, es.Planes[0])
	for i := 1; i < len(es.Planes)-1; i++ {
		plane := es.ComputeNextCycleState(i)
		nextPlanes = append(nextPlanes, plane)
	}
	nextPlanes = append(nextPlanes, es.Planes[len(es.Planes)-1])
	es.Planes = nextPlanes
}

func (es *EnergySource3D) ComputeNextCycleState(planeIndex int) ConwayCubePlane {
	//fmt.Printf("  ComputeNextCycleState(%d)\n", planeIndex)
	ccp := es.Planes[planeIndex]
	resPlane := MakeEmptyConwayCubePlane(ccp.Height, ccp.Width)
	for r := 1; r < ccp.Height+1; r++ {
		for c := 1; c < ccp.Width+1; c++ {
			switch ccp.Cubes[r][c] {
			case ACTIVE:
				count := es.CountActiveNeighbours(planeIndex, r, c)
				if count == 2 || count == 3 {
					//fmt.Printf("Change cube (1) %d/%d/%d\n", planeIndex, r, c)
					resPlane.Cubes[r][c] = ACTIVE
				} else {
					//fmt.Printf("Change cube (2) %d/%d/%d\n", planeIndex, r, c)
					resPlane.Cubes[r][c] = INACTIVE
				}
			case INACTIVE:
				count := es.CountActiveNeighbours(planeIndex, r, c)
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

func (es EnergySource3D) CountActiveNeighbours(planeIndex int, rowIndex int, columnIndex int) int {
	res := 0

	for _, dp := range [...]int{-1, 0, 1} {
		for _, dr := range [...]int{-1, 0, 1} {
			for _, dc := range [...]int{-1, 0, 1} {
				if dp == 0 && dr == 0 && dc == 0 {
					continue
				}
				if es.Planes[planeIndex+dp].Cubes[rowIndex+dr][columnIndex+dc] == ACTIVE {
					res++
				}
			}
		}
	}

	return res
}

func MakeEnergySource3D(initialState []string, intendedRoundCount int) EnergySource3D {
	resESource := EnergySource3D{
		Planes: make([]ConwayCubePlane, 0, 2*intendedRoundCount+3),
	}
	fmt.Printf("  Length of resESource.Planes: %d\n", len(resESource.Planes))

	for i := 0; i < intendedRoundCount+1; i++ {
		resESource.Planes = append(resESource.Planes, MakeEmptyConwayCubePlane(len(initialState)+2*intendedRoundCount, len(initialState[0])+2*intendedRoundCount))
	}
	fmt.Printf("  Length of resESource.Planes: %d\n", len(resESource.Planes))
	resESource.Planes = append(resESource.Planes, MakeConwayCubePlane(initialState, intendedRoundCount))
	fmt.Printf("  Length of resESource.Planes: %d\n", len(resESource.Planes))
	for i := 0; i < intendedRoundCount+1; i++ {
		resESource.Planes = append(resESource.Planes, MakeEmptyConwayCubePlane(len(initialState)+2*intendedRoundCount, len(initialState[0])+2*intendedRoundCount))
	}
	fmt.Printf("  Length of resESource.Planes: %d\n", len(resESource.Planes))

	return resESource
}
