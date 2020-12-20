package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 19: Monster Messages\n=====================")

	puzzleInput := readPuzzleInput("RawData.txt")
	fmt.Printf("Puzzle input:\n%v\n", puzzleInput)

	fmt.Println("\nPart 1: Count completely matched messages\n----------------------------------")
	solvePart1(puzzleInput)

	//fmt.Println("\nPart 2: Multiply all departure values on my ticket\n-------------------------------------------------------")
	//solvePart2(puzzleInput)
}

func readPuzzleInput(filename string) PuzzleInput {
	// 1st read the file content
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	res := MakePuzzleInput()
	inputSection := 0
	for _, line := range strings.Split(string(buf), "\r\n") {
		if len(line) == 0 {
			inputSection++
			continue
		}
		switch inputSection {
		case 0: // Rules
			rule, err := MakeRule(line)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error occurred when trying to parse rule data from '%s': %v", line, err)
				os.Exit(1)
			}
			res.RuleSet[rule.Id] = rule
		case 1: // Messages
			res.Messages = append(res.Messages, line)
		}
	}

	return res
}

func solvePart1(puzzleInput PuzzleInput) {
	regexFormat := reduceRulesetToRegex(puzzleInput.RuleSet)

	validMessageCount := 0
	for _, m := range puzzleInput.Messages {
		matched, _ := regexp.MatchString(regexFormat, m)
		if matched {
			validMessageCount++
		}
	}

	fmt.Printf("Result for part 1: %d\n", validMessageCount)
}

func reduceRulesetToRegex(rules map[int]Rule) string {
	res := `^ 0 $`

	changed := true
	for changed {
		fmt.Println(res)
		changed = false
		fields := strings.Fields(res)
		var sb strings.Builder
		for _, f := range fields {
			val, err := strconv.Atoi(f)
			if err == nil {
				sb.WriteString(fmt.Sprintf("( %s )", rules[val].Syntax))
				changed = true
			} else {
				sb.WriteString(f)
			}
		}
		res = sb.String()
	}
	fmt.Println(res)

	return res
}

//func solvePart2(puzzleInput PuzzleInput) {
//	validTickets := DetermineValidTickets(puzzleInput)
//	//fmt.Printf("  Valid tickets: %v\n", validTickets)
//	possibleFieldPositions := ComputePossibleFieldPositions(puzzleInput, validTickets)
//	fmt.Printf("  Possible positions: %v\n", possibleFieldPositions)
//
//	prevLen := -1
//	for len(possibleFieldPositions) > 0 {
//		if prevLen == len(possibleFieldPositions) {
//			fmt.Println("Could not go further?!?!")
//			break
//		}
//		prevLen = len(possibleFieldPositions)
//		fieldsDone := make([]string, 0)
//		for name, positions := range possibleFieldPositions {
//			if len(positions) == 1 {
//				SetTicketFieldPosition(&puzzleInput, name, positions[0])
//				//fmt.Printf("  Determined %s is at position %d\n", name, positions[0])
//				for oName, oPositions := range possibleFieldPositions {
//					if oName == name {
//						continue
//					}
//					delPos := -1
//					for i, n := range oPositions {
//						if n == positions[0] {
//							delPos = i
//						}
//					}
//					if delPos > -1 {
//						oPositions[delPos] = oPositions[len(oPositions)-1]
//						possibleFieldPositions[oName] = possibleFieldPositions[oName][:len(possibleFieldPositions[oName])-1]
//					}
//				}
//				fieldsDone = append(fieldsDone, name)
//			}
//		}
//		for _, name := range fieldsDone {
//			delete(possibleFieldPositions, name)
//		}
//		//fmt.Printf("  Ticked fields: %v\n", puzzleInput.TicketFields)
//		fmt.Printf("  Possible positions: %v\n", possibleFieldPositions)
//	}
//	fmt.Printf("Ticked fields: %v\n", puzzleInput.TicketFields)
//
//	res := 1
//	for _, field := range puzzleInput.TicketFields {
//		if strings.HasPrefix(field.Name, "departure") {
//			res = res * puzzleInput.MyTicket.Values[field.Position]
//		}
//	}
//	fmt.Printf("Result for part 2: %d\n", res)
//}
