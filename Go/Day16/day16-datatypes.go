package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type FieldConstraint struct {
	LowerBound int
	UpperBound int
}

type Field struct {
	Name        string
	Constraints []FieldConstraint
	Position    int
}

func (f Field) IsConstraintSatisfied(value int) bool {
	for _, c := range f.Constraints {
		if c.LowerBound <= value && value <= c.UpperBound {
			return true
		}
	}

	return false
}

func MakeField(line string) (Field, error) {
	mainParts := strings.Split(line, ": ")
	res := Field{Name: mainParts[0], Position: -1}

	constraints := strings.Split(mainParts[1], " or ")
	for _, c := range constraints {
		bounds := strings.Split(c, "-")
		lb, lerr := strconv.Atoi(bounds[0])
		ub, rerr := strconv.Atoi(bounds[1])
		if lerr != nil || rerr != nil {
			return res, errors.New(fmt.Sprintf("Could not parse constraints: %s", c))
		}
		res.Constraints = append(res.Constraints, FieldConstraint{LowerBound: lb, UpperBound: ub})
	}

	return res, nil
}

type Ticket struct {
	Values []int
}

func MakeTicket(line string) (Ticket, error) {
	token := strings.Split(line, ",")

	resTicket := Ticket{
		Values: make([]int, 0, len(token)),
	}
	for _, t := range token {
		v, err := strconv.Atoi(t)
		if err != nil {
			return resTicket, errors.New(fmt.Sprintf("Value '%s' on ticket could not be parsed as int", t))
		}
		resTicket.Values = append(resTicket.Values, v)
	}

	return resTicket, nil
}

type PuzzleInput struct {
	TicketFields []Field
	MyTicket     Ticket
	OtherTickets []Ticket
}
