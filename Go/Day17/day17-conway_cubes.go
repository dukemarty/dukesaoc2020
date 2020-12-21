package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 17: Conway Cubes\n===========================")

	initialState := readInitialState("RawData.txt")
	fmt.Printf("Initial seat plan layout:\n%s\n", strings.Join(initialState, "\n"))

	fmt.Println("\nPart 1: Count active cubes\n--------------------------------")
	solvePart1(initialState)

	fmt.Println("\nPart 2: Count active cubes with additional dimension\n-------------------------------------------------------")
	solvePart2(initialState)
}

func readInitialState(filename string) []string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	rows := strings.Split(string(buf), "\r\n")

	return rows
}

func solvePart1(initialState []string) {
	energySource := MakeEnergySource3D(initialState, 8)
	//fmt.Println(energySource)
	//fmt.Printf("Length of es.Planes: %d\n", len(energySource.Planes))

	for i := 0; i < 6; i++ {
		energySource.ExecuteCycle()
		//fmt.Println("=============================================")
		//fmt.Println(energySource)
		//fmt.Printf("Number of active cubes with i=%d: %d\n", i, energySource.CountActiveCubes())
	}

	res := energySource.CountActiveCubes()
	fmt.Printf("Active cubes after 6 rounds: %d\n", res)
}

func solvePart2(initialState []string) {
	energySource := MakeEnergySource4D(initialState, 8)
	fmt.Println(energySource)
	fmt.Printf("Length of es.Planes: %d\n", len(energySource.Planes))

	for i := 0; i < 6; i++ {
		energySource.ExecuteCycle()
		//fmt.Println("=============================================")
		//fmt.Println(energySource)
		//fmt.Printf("Number of active cubes with i=%d: %d\n", i, energySource.CountActiveCubes())
	}

	res := energySource.CountActiveCubes()
	fmt.Printf("Active cubes after 6 rounds: %d\n", res)
}
