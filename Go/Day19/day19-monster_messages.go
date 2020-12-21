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

	fmt.Println("\nPart 2: Count completely matched messages with updated rules\n-------------------------------------------------------")
	solvePart2(puzzleInput)
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

func solvePart2(puzzleInput PuzzleInput) {
	regexFormat := reduceRulesetToRegex(puzzleInput.RuleSet)
	fmt.Println(regexFormat)
	puzzleInput.RuleSet[8] = Rule{
		Id:     8,
		Syntax: "42 +",
	}
	regexFormat = reduceRulesetToRegex(puzzleInput.RuleSet)
	fmt.Println(regexFormat)
	puzzleInput.RuleSet[11] = Rule{
		Id:     11,
		Syntax: "42 31 | 42 42 31 31 | 42 42 42 31 31 31 | 42 42 42 42 31 31 31 31 | 42 42 42 42 42 31 31 31 31 31",
	}
	regexFormat = reduceRulesetToRegex(puzzleInput.RuleSet)
	fmt.Println(regexFormat)

	validMessageCount := 0
	for _, m := range puzzleInput.Messages {
		matched, _ := regexp.MatchString(regexFormat, m)
		if matched {
			validMessageCount++
		}
	}

	fmt.Printf("Result for part 2: %d\n", validMessageCount)
}

func reduceRulesetToRegex(rules map[int]Rule) string {
	res := `^ 0 $`

	changed := true
	for changed {
		//fmt.Println(res)
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
