package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Color   string
	Content map[string]int
}

func (r Rule) String() string {
	return fmt.Sprintf("Rule: %s -> [%q]", r.Color, r.Content)
}

func MakeRule(raw string) Rule {
	res := Rule{
		Content: make(map[string]int),
	}

	mainParts := strings.Split(raw, " contain ")
	res.Color = strings.TrimSuffix(strings.TrimSuffix(mainParts[0], " bag"), " bags")

	if mainParts[1] == "no other bags." {
		return res
	}

	containedBags := strings.Split(mainParts[1], ", ")
	re, _ := regexp.Compile(`^(?P<count>\d+) (?P<bag>[^.]*) bags?\.?$`)
	for _, cb := range containedBags {
		match := re.FindStringSubmatch(cb)
		cbCount, err := strconv.Atoi(match[1])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error when trying to parse bag count: %s", err)
			continue
		}
		res.Content[match[2]] = cbCount
	}

	return res
}

type InnerBag struct {
	Color string
	Count int
}
