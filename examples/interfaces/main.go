package main

import (
	"fmt"

	lev "github.com/gitchander/levenshtein"
)

type Person struct {
	Name string
	Age  int
}

type PersonPair [2][]Person

var _ lev.Interface = PersonPair{}

func (p PersonPair) Lens() (ni, nj int) {
	return len(p[0]), len(p[1])
}

func (p PersonPair) Match(i, j int) bool {
	return p[0][i] == p[1][j]
}

func main() {
	var (
		a = []Person{{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}}
		b = []Person{{"one", 1}, {"two", 2}, {"three", 3}}
	)
	distance := lev.Distance(PersonPair{a, b})
	fmt.Printf("the levenshtein distance = %d\n", distance)
}
