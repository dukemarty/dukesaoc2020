package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Day 21: Allergen Assessment\n========================")

	foodList := readFoodList("RawData.txt")
	fmt.Printf("Read rules: %q\n", foodList)

	fmt.Println("\nPart 1: Count occurrences of non-allergent ingredients\n----------------------------------")
	solvePart1(foodList)

	fmt.Println("\nPart 2: Allergen-alphabetic ingredients list\n-------------------------------------")
	solvePart2(foodList)
}

func readFoodList(filename string) []Food {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error occurred when trying to read data from file: ", err)
		os.Exit(1)
	}

	var res []Food
	for _, line := range strings.Split(string(buf), "\r\n") {
		rule := MakeFood(line)
		res = append(res, rule)
	}

	return res
}

func solvePart1(foodList []Food) {
	ingredients, allergens := constructContentLists(foodList)

	//fmt.Printf("Ingredients: %v\n", ingredients)
	//fmt.Printf("Allergens: %v\n", allergens)

	res := 0
	for i, c := range ingredients {
		if !isContainerCandidate(i, allergens) {
			res = res + c
		}
	}

	fmt.Printf("Number of occurrences of not-allergen-containing ingredients: %d\n", res)
}

func isContainerCandidate(ingredient string, allergens map[string]Allergen) bool {
	for _, a := range allergens {
		if a.PossibleContainers[ingredient] {
			return true
		}
	}

	return false
}

func solvePart2(foodList []Food) {
	_, allergens := constructContentLists(foodList)
	determineAllergenContainers(allergens)

	fmt.Printf("Allergens: %v\n", allergens)

	allAllergens := make([]string, 0, len(allergens))
	for n := range allergens {
		allAllergens = append(allAllergens, n)
	}
	sort.Strings(allAllergens)

	orderedIngredients := make([]string, 0, len(allAllergens))
	for _, a := range allAllergens {
		orderedIngredients = append(orderedIngredients, allergens[a].Container)
	}

	fmt.Printf("Ordered container elements: %s\n", strings.Join(orderedIngredients, ","))
}

func determineAllergenContainers(allergens map[string]Allergen) {
	changesMade := true
	for changesMade {
		changesMade = false
		for _, a := range allergens {
			if len(a.PossibleContainers) == 1 {
				var container string
				//fmt.Printf("  Found container of %s: %v\n", a.Name, a.PossibleContainers)
				for i := range a.PossibleContainers {
					container = i
				}
				newAllergen := MakeAllergen(a.Name)
				newAllergen.Container = container
				allergens[newAllergen.Name] = newAllergen
				for _, b := range allergens {
					delete(b.PossibleContainers, container)
				}
				changesMade = true
			}
		}
	}
}

func constructContentLists(foodList []Food) (map[string]int, map[string]Allergen) {
	resIngredients := make(map[string]int)
	resAllergens := make(map[string]Allergen)

	for _, f := range foodList {
		ingredientDict := make(map[string]bool)
		for _, i := range f.Ingredients {
			ingredientDict[i] = true
			if _, ok := resIngredients[i]; ok {
				resIngredients[i]++
			} else {
				resIngredients[i] = 1
			}
		}
		for _, a := range f.Allergens {
			if _, ok := resAllergens[a]; ok {
				for k, v := range resAllergens[a].PossibleContainers {
					if v == true && ingredientDict[k] != true {
						resAllergens[a].PossibleContainers[k] = false
						delete(resAllergens[a].PossibleContainers, k)
					}
				}
			} else {
				allergen := MakeAllergen(a)
				for _, i := range f.Ingredients {
					allergen.PossibleContainers[i] = true
				}
				resAllergens[a] = allergen
			}
		}
	}

	return resIngredients, resAllergens
}
