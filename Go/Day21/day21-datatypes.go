package main

import (
	"fmt"
	"strings"
)

type Food struct {
	Ingredients []string
	Allergens   []string
}

func (f Food) String() string {
	return fmt.Sprintf("Food: %s -> [%s]", strings.Join(f.Ingredients, ","), strings.Join(f.Allergens, ","))
}

func MakeFood(raw string) Food {

	mainParts := strings.Split(raw, " (contains ")
	ingredients := strings.Fields(mainParts[0])
	allergens := strings.Split(strings.TrimRight(mainParts[1], ")"), ", ")

	res := Food{
		Ingredients: ingredients,
		Allergens:   allergens,
	}

	return res
}

type Allergen struct {
	Name               string
	PossibleContainers map[string]bool
	Container          string
}

func (a Allergen) String() string {
	return fmt.Sprintf("%s <- %s<=[%v]", a.Name, a.Container, a.PossibleContainers)
}

func MakeAllergen(name string) Allergen {
	res := Allergen{
		Name:               name,
		PossibleContainers: make(map[string]bool),
	}

	return res
}
