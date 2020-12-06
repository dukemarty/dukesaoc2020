package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 06: Custom Customs\n======================")

	groupAnswers := readGroupAnswers("RawData.txt")
	fmt.Printf("Read answers: %q\n", groupAnswers)

	fmt.Println("\nPart 1: Count distinct yes-answers\n----------------------------------")
	solvePart1(groupAnswers)

	fmt.Println("\nPart 2: Count agreed yes-answers\n--------------------------------")
	solvePart2(groupAnswers)
}

func readGroupAnswers(filename string) []GroupAnswers {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	var res []GroupAnswers
	answers := GroupAnswers{}
	for _, row := range strings.Split(string(buf), "\r\n") {
		if len(row) == 0 {
			res = append(res, answers)
			answers = GroupAnswers{}
			continue
		}

		answers.Raw = append(answers.Raw, row)
	}

	if len(answers.Raw) > 0 {
		res = append(res, answers)
	}

	return res
}

func solvePart1(groupAnswers []GroupAnswers) {
	res := 0
	for _, ga := range groupAnswers {
		res = res + ga.CountAny()
	}

	fmt.Printf("Number of distinct answers with 'yes': %d\n", res)
}

func solvePart2(groupAnswers []GroupAnswers) {
	res := 0
	for _, ga := range groupAnswers {
		res = res + ga.CountAll()
	}

	fmt.Printf("Number of agreed answers with 'yes': %d\n", res)
}
