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
	fmt.Println("Day 04: Passport Processing\n===========================")

	passports := readPassports("RawData.txt")
	fmt.Printf("Passports:\n%q\n", passports)

	fmt.Println("\nPart 1: Check required fields\n----------------------------------------")
	solvePart1(passports)

	fmt.Println("\nPart 2: Validness of fields\n---------------------------------------")
	solvePart2(passports)
}

// readPassports parses Passport objects from text file given via filename. It returns an array of Passport's.
func readPassports(filename string) []Passport {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	var res []Passport
	pass := MakePassport()
	for _, row := range strings.Split(string(buf), "\r\n") {
		if len(row) == 0 {
			res = append(res, pass)
			pass = MakePassport()
			continue
		}

		for _, field := range strings.Fields(row) {
			parts := strings.Split(field, ":")
			pass.Fields[parts[0]] = parts[1]
		}
	}

	if len(pass.Fields) > 0 {
		res = append(res, pass)
	}

	return res
}

func solvePart1(passports []Passport) {
	validCount := 0
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, p := range passports {
		fieldMissing := false
		for _, rf := range requiredFields {
			_, ok := p.Fields[rf]
			if !ok {
				fieldMissing = true
				break
			}
		}
		if !fieldMissing {
			validCount++
		}
	}

	fmt.Printf("Number of valid passports: %d\n", validCount)
}

func solvePart2(passports []Passport) {
	validCount := 0

	for _, p := range passports {
		if !isBasicFormValid(p) {
			continue
		}

		if areAllValueRangesValid(p) {
			validCount++
		}
	}

	fmt.Printf("Number of valid passports: %d\n", validCount)
}

// isBasicFormValid checks simple conditions for the validness of a Passport object pass.
// It returns true if
//  a) all required fields exists
//  b) the basic format of those fields values is correct (e.g. check if it is a number,
//     the range is not checked here)
func isBasicFormValid(pass Passport) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	requiredForms := map[string]string{
		"byr": `^\d{4}$`,
		"iyr": `^\d{4}$`,
		"eyr": `^\d{4}$`,
		"hgt": `^\d+(cm)|(in)$`,
		"hcl": `^#[0-9a-f]{6}$`,
		"ecl": `^(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)$`,
		"pid": `^\d{9}$`,
	}

	fieldInvalidOrMissing := false
	for _, rf := range requiredFields {
		elem, ok := pass.Fields[rf]
		if !ok {
			fieldInvalidOrMissing = true
			break
		}
		matched, _ := regexp.MatchString(requiredForms[rf], elem)
		if !matched {
			fieldInvalidOrMissing = true
			break
		}
	}

	return !fieldInvalidOrMissing
}

// areAllValueRangesValid checks Passport object pass for correct values ranges of different field.
// It returns true if all of the fields "byr", "iyr", "eyr" and "hgt" have values in a given range.
func areAllValueRangesValid(pass Passport) bool {
	allValuesValid := true

	allValuesValid = allValuesValid && checkIntValueRange(pass.Fields["byr"], 1920, 2002)
	allValuesValid = allValuesValid && checkIntValueRange(pass.Fields["iyr"], 2010, 2020)
	allValuesValid = allValuesValid && checkIntValueRange(pass.Fields["eyr"], 2020, 2030)
	hgt := pass.Fields["hgt"]
	if hgt[len(hgt)-1] == 'm' {
		allValuesValid = allValuesValid && checkIntValueRange(hgt[0:len(hgt)-2], 150, 193)
	} else {
		allValuesValid = allValuesValid && checkIntValueRange(hgt[0:len(hgt)-2], 59, 76)
	}

	return allValuesValid
}

func checkIntValueRange(value string, rangeMin int, rangeMax int) bool {
	i, _ := strconv.Atoi(value) // previous check "proved" that value is a number

	return (rangeMin <= i) && (i <= rangeMax)
}
