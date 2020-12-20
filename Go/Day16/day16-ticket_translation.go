package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 16: Docking Data\n=====================")

	puzzleInput := readPuzzleInput("RawData.txt")
	fmt.Printf("Program: %v\n", puzzleInput)

	fmt.Println("\nPart 1: Count invalid tickets\n----------------------------------")
	solvePart1(puzzleInput)

	fmt.Println("\nPart 2: Multiply all departure values on my ticket\n-------------------------------------------------------")
	solvePart2(puzzleInput)
}

func readPuzzleInput(filename string) PuzzleInput {
	// 1st read the file content
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	res := PuzzleInput{}
	inputSection := 0
	for _, line := range strings.Split(string(buf), "\r\n") {
		if len(line) == 0 {
			inputSection++
			continue
		}
		switch inputSection {
		case 0:
			field, err := MakeField(line)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error occurred when trying to parse ticket data: %v", err)
				os.Exit(1)
			}
			res.TicketFields = append(res.TicketFields, field)
		case 1:
			if line == "your ticket:" {
				continue
			}
			ticket, err := MakeTicket(line)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error occurred when trying to parse ticket data: %v", err)
				os.Exit(1)
			}
			res.MyTicket = ticket
		case 2:
			if line == "nearby tickets:" {
				continue
			}
			ticket, err := MakeTicket(line)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error occurred when trying to parse ticket data: %v", err)
				os.Exit(1)
			}
			res.OtherTickets = append(res.OtherTickets, ticket)
		}
	}

	return res
}

func solvePart1(puzzleInput PuzzleInput) {
	ticketScanningErrorRate := 0

	for _, ticket := range puzzleInput.OtherTickets {
		for _, value := range ticket.Values {
			foundSatisfiableConstraint := false
			for _, ticketField := range puzzleInput.TicketFields {
				if ticketField.IsConstraintSatisfied(value) {
					foundSatisfiableConstraint = true
					break
				}
			}
			if !foundSatisfiableConstraint {
				ticketScanningErrorRate = ticketScanningErrorRate + value
			}
		}
	}

	fmt.Printf("Result for part 1: %d\n", ticketScanningErrorRate)
}

func solvePart2(puzzleInput PuzzleInput) {
	validTickets := DetermineValidTickets(puzzleInput)
	//fmt.Printf("  Valid tickets: %v\n", validTickets)
	possibleFieldPositions := ComputePossibleFieldPositions(puzzleInput, validTickets)
	fmt.Printf("  Possible positions: %v\n", possibleFieldPositions)

	prevLen := -1
	for len(possibleFieldPositions) > 0 {
		if prevLen == len(possibleFieldPositions) {
			fmt.Println("Could not go further?!?!")
			break
		}
		prevLen = len(possibleFieldPositions)
		fieldsDone := make([]string, 0)
		for name, positions := range possibleFieldPositions {
			if len(positions) == 1 {
				SetTicketFieldPosition(&puzzleInput, name, positions[0])
				//fmt.Printf("  Determined %s is at position %d\n", name, positions[0])
				for oName, oPositions := range possibleFieldPositions {
					if oName == name {
						continue
					}
					delPos := -1
					for i, n := range oPositions {
						if n == positions[0] {
							delPos = i
						}
					}
					if delPos > -1 {
						oPositions[delPos] = oPositions[len(oPositions)-1]
						possibleFieldPositions[oName] = possibleFieldPositions[oName][:len(possibleFieldPositions[oName])-1]
					}
				}
				fieldsDone = append(fieldsDone, name)
			}
		}
		for _, name := range fieldsDone {
			delete(possibleFieldPositions, name)
		}
		//fmt.Printf("  Ticked fields: %v\n", puzzleInput.TicketFields)
		fmt.Printf("  Possible positions: %v\n", possibleFieldPositions)
	}
	fmt.Printf("Ticked fields: %v\n", puzzleInput.TicketFields)

	res := 1
	for _, field := range puzzleInput.TicketFields {
		if strings.HasPrefix(field.Name, "departure") {
			res = res * puzzleInput.MyTicket.Values[field.Position]
		}
	}
	fmt.Printf("Result for part 2: %d\n", res)
}

func SetTicketFieldPosition(puzzleInput *PuzzleInput, name string, position int) {
	for i, tf := range puzzleInput.TicketFields {
		if tf.Name == name {
			puzzleInput.TicketFields[i].Position = position
			break
		}
	}
}

func DetermineValidTickets(puzzleInput PuzzleInput) []Ticket {
	validTickets := []Ticket{puzzleInput.MyTicket}
	for _, ticket := range puzzleInput.OtherTickets {
		anyInvalidValueFound := false
		for _, value := range ticket.Values {
			foundSatisfiableConstraint := false
			for _, ticketField := range puzzleInput.TicketFields {
				if ticketField.IsConstraintSatisfied(value) {
					foundSatisfiableConstraint = true
					break
				}
			}
			if !foundSatisfiableConstraint {
				anyInvalidValueFound = true
				break
			}
		}
		if !anyInvalidValueFound {
			validTickets = append(validTickets, ticket)
		}
	}

	return validTickets
}

func ComputePossibleFieldPositions(puzzleInput PuzzleInput, validTickets []Ticket) map[string][]int {
	possibleFieldPositions := make(map[string][]int)
	for _, field := range puzzleInput.TicketFields {
		possibleFieldPositions[field.Name] = make([]int, 0)
	}

	for _, field := range puzzleInput.TicketFields {
		for i := 0; i < len(puzzleInput.TicketFields); i++ {
			isPossiblePosition := true
			for _, vt := range validTickets {
				if !field.IsConstraintSatisfied(vt.Values[i]) {
					//fmt.Printf("Constraint for field %v at pos %d is broken by %v.\n", field, i, vt)
					isPossiblePosition = false
					break
				}
			}
			if isPossiblePosition {
				possibleFieldPositions[field.Name] = append(possibleFieldPositions[field.Name], i)
			}
		}
	}

	return possibleFieldPositions
}
