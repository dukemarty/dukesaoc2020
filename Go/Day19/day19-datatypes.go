package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	Id     int
	Syntax string
}

func (r Rule) String() string {
	return fmt.Sprintf("(%d -> %s)", r.Id, r.Syntax)
}

func MakeRule(line string) (Rule, error) {
	mainParts := strings.Split(line, ": ")

	id, err := strconv.Atoi(mainParts[0])
	if err != nil {
		return Rule{}, errors.New(fmt.Sprintf("Could not parse id from %s"))
	}

	res := Rule{
		Id:     id,
		Syntax: strings.Trim(strings.TrimSpace(mainParts[1]), "\""),
	}

	return res, nil
}

type PuzzleInput struct {
	RuleSet  map[int]Rule
	Messages []string
}

func (pi PuzzleInput) String() string {
	return fmt.Sprintf("%d Rules: %v\n%d Messages: [%s]\n", len(pi.RuleSet), pi.RuleSet, len(pi.Messages), strings.Join(pi.Messages, ", "))
}

func MakePuzzleInput() PuzzleInput {
	res := PuzzleInput{
		RuleSet: make(map[int]Rule),
	}

	return res
}
