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
	fmt.Println("Day 01: Operation Order\n=====================")

	expressions := readExpressions("RawData.txt")
	fmt.Printf("Read numbers: %v\n", expressions)

	fmt.Println("\nPart 1: Sum of evaluated expressions with left-to-right precedence\n-----------------------------------------------")
	solvePart1(expressions)

	fmt.Println("\nPart 2: Sum of evaluated expressions with add highest precedence\n----------------------------------------------")
	solvePart2(expressions)
}

func readExpressions(filename string) []string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	expressionStrings := strings.Split(string(buf), "\r\n")

	return expressionStrings
}

func solvePart1(expressions []string) {
	res := 0
	for _, e := range expressions {
		exprRes := evaluateExpression(e, evaluateSimpleExpressionLeftToRight)
		fmt.Printf("  Result of '%s': %d\n", e, exprRes)
		res = res + exprRes
	}

	fmt.Printf("Sum of all expressions: %d\n", res)
}

func solvePart2(expressions []string) {
	res := 0
	for _, e := range expressions {
		exprRes := evaluateExpression(e, evaluateSimpleExpressionPlusOverDot)
		fmt.Printf("  Result of '%s': %d\n", e, exprRes)
		res = res + exprRes
	}

	fmt.Printf("Sum of all expressions: %d\n", res)
}

func evaluateExpression(expr string, evaluateSimpleExpression func(string) string) int {
	currExpr := expr
	parenthesesRegexFormat := `\([0-9+* ]+\)`
	parenRegex, err := regexp.Compile(parenthesesRegexFormat)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error trying to find compile parentheses expression in '%s': %q", parenthesesRegexFormat, err)
		return -1
	}
	for strings.Contains(currExpr, "(") {
		fmt.Printf("    currExpr = %s\n", currExpr)
		match := parenRegex.FindString(currExpr)
		//fmt.Printf("    Found match: %s\n", match)
		matchValue := evaluateSimpleExpression(match[1 : len(match)-1])
		currExpr = strings.Replace(currExpr, match, matchValue, 1)
	}
	fmt.Printf("    currExpr = %s\n", currExpr)

	resString := evaluateSimpleExpression(currExpr)
	res, err := strconv.Atoi(resString)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", resString, err)
		res = -1
	}

	return res
}

func evaluateSimpleExpressionLeftToRight(expr string) string {
	fields := strings.Fields(expr)
	fmt.Println(fields)

	res, err := strconv.Atoi(fields[0])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", fields[0], err)
		return "-1"
	}

	for i := 1; i < len(fields); {
		op := fields[i]
		i++
		rval, err := strconv.Atoi(fields[i])
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", fields[i], err)
			return "-1"
		}
		i++
		switch op {
		case "+":
			res = res + rval
		case "*":
			res = res * rval
		}
	}

	return fmt.Sprintf("%d", res)
}

func evaluateSimpleExpressionPlusOverDot(expr string) string {
	parts := strings.Split(expr, " * ")
	fmt.Println(parts)

	res, err := strconv.Atoi(evaluateSimpleExpressionLeftToRight(parts[0]))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", parts[0], err)
		return "-1"
	}

	for i := 1; i < len(parts); {
		rval, err := strconv.Atoi(evaluateSimpleExpressionLeftToRight(parts[i]))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", parts[i], err)
			return "-1"
		}
		i++
		res = res * rval
	}

	return fmt.Sprintf("%d", res)
}
