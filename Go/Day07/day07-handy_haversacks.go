package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 07: Handy Haversacks\n========================")

	baggageRules := readBaggageRules("RawData.txt")
	fmt.Printf("Read rules: %q\n", baggageRules)

	fmt.Println("\nPart 1: Count possible bags for gold\n----------------------------------")
	solvePart1(baggageRules)

	fmt.Println("\nPart 2: Count number of required bags\n-------------------------------------")
	solvePart2(baggageRules)
}

func readBaggageRules(filename string) []Rule {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	var res []Rule
	for _, line := range strings.Split(string(buf), "\r\n") {
		rule := MakeRule(line)
		res = append(res, rule)
	}

	return res
}

func solvePart1(rules []Rule) {
	isContainedIn := make(map[string][]string)
	for _, r := range rules {
		for k := range r.Content {
			if _, ok := isContainedIn[k]; !ok {
				isContainedIn[k] = []string{}
			}
			isContainedIn[k] = append(isContainedIn[k], r.Color)
		}
	}

	unprocessed := make([]string, len(isContainedIn["shiny gold"]))
	copy(unprocessed, isContainedIn["shiny gold"])
	resultSet := make(map[string]bool)
	for len(unprocessed) > 0 {
		next := unprocessed[0]
		unprocessed = unprocessed[1:]
		if !resultSet[next] {
			resultSet[next] = true
			unprocessed = append(unprocessed, isContainedIn[next]...)
		}
	}

	fmt.Println(isContainedIn)

	fmt.Printf("Number of bags who contain a 'shiny gold' bag: %d\n", len(resultSet))
}

func solvePart2(rules []Rule) {
	bagCatalog := make(map[string][]InnerBag)
	for _, r := range rules {
		bagCatalog[r.Color] = transformContentToInnerBags(r)
	}

	res := 0
	marked := make([]InnerBag, len(bagCatalog["shiny gold"]))
	copy(marked, bagCatalog["shiny gold"])
	for len(marked) > 0 {
		next := marked[0]
		marked = marked[1:]
		res += next.Count
		for _, ib := range bagCatalog[next.Color] {
			marked = append(marked, InnerBag{Color: ib.Color, Count: next.Count * ib.Count})
		}
	}

	fmt.Printf("Number of bags in a 'shiny gold' bag: %d\n", res)
}

func transformContentToInnerBags(rule Rule) []InnerBag {
	res := make([]InnerBag, len(rule.Content))

	for color, count := range rule.Content {
		ib := InnerBag{
			Color: color,
			Count: count,
		}
		res = append(res, ib)
	}

	return res
}
