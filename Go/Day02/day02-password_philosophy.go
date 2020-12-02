package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 02: Password Philosophy\n=====================")

	checkCases := readPasswords("RawData.txt")
	fmt.Printf("Read passwords: %q\n", checkCases)

	fmt.Println("\nPart 1: Number of valid passwords old variant\n---------------------------------------------")
	solvePart1(checkCases)

	fmt.Println("\nPart 2: Number of valid passwords new variant\n---------------------------------------------")
	solvePart2(checkCases)
}

type Policy struct {
	LeftDigit  int
	RightDigit int
	Letter     string
}

func (p Policy) String() string {
	return fmt.Sprintf("Policy(%d/%d -> '%s')", p.LeftDigit, p.RightDigit, p.Letter)
}

type CheckCase struct {
	Policy   Policy
	Password string
}

func readPasswords(filename string) []CheckCase {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	var res []CheckCase
	for _, line := range strings.Split(string(buf), "\r\n") {
		// 1-7 l: zlmsmlxpvvlzv
		parts := strings.Split(line, ":")
		// "1-7 l", " zlmsmlxpvvlzv"
		policyParts := strings.Fields(parts[0])
		// "1-7", "l" / " zlmsmlxpvvlzv"
		occBounds := strings.Split(policyParts[0], "-")
		// "1", "7" / "l" // " zlmsmlxpvvlzv"
		minOcc, err := strconv.Atoi(occBounds[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", minOcc, err)
			continue
		}
		maxOcc, err := strconv.Atoi(occBounds[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing '%s' as integer: %q", maxOcc, err)
			continue
		}
		entry := CheckCase{
			Policy: Policy{
				Letter:     policyParts[1],
				LeftDigit:  minOcc,
				RightDigit: maxOcc,
			},
			Password: strings.TrimSpace(parts[1]),
		}
		res = append(res, entry)
	}

	return res
}

func solvePart1(cases []CheckCase) {
	res := 0

	for _, check := range cases {
		if isPasswordValidNumberOfLetter(check.Password, check.Policy) {
			res += 1
		}
	}

	fmt.Printf("Number of valid passwords: %d\n", res)
}

func isPasswordValidNumberOfLetter(password string, policy Policy) bool {
	count := strings.Count(password, policy.Letter)

	return policy.LeftDigit <= count && count <= policy.RightDigit
}

func solvePart2(cases []CheckCase) {
	res := 0

	for _, check := range cases {
		if isPasswordValidLetterPosition(check.Password, check.Policy) {
			res += 1
		} else {
			fmt.Println(check)
		}
	}

	fmt.Printf("Number of valid passwords: %d\n", res)
}

func isPasswordValidLetterPosition(password string, policy Policy) bool {
	fmt.Printf("    Elements: %d , %d , %d\n", password[policy.LeftDigit-1], password[policy.RightDigit-1], policy.Letter[0])
	fmt.Printf("    Elements: >%s< , >%s< , >%s<\n", string(password[policy.LeftDigit-1]), string(password[policy.RightDigit-1]), string(policy.Letter[0]))
	firstOcc := password[policy.LeftDigit-1] == policy.Letter[0]
	secondOcc := password[policy.RightDigit-1] == policy.Letter[0]
	fmt.Printf("    1st/2nd: %t/%t\n", firstOcc, secondOcc)
	return firstOcc != secondOcc
}
