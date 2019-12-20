package main

import (
	"fmt"

	lev "github.com/gitchander/levenshtein"
)

func main() {
	var (
		a = []Person{{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}}
		b = []Person{{"one", 1}, {"two", 2}, {"three", 3}}
	)
	distance := lev.Distance(PersonSlices{a, b})
	fmt.Printf("the levenshtein distance = %d\n", distance)
}

type Person struct {
	Name string
	Age  int
}

type PersonSlices [2][]Person

var _ lev.Interface = PersonSlices{}

func (p PersonSlices) Len(k int) int {
	return len(p[k])
}

func (p PersonSlices) Match(i, j int) bool {
	return p[0][i] == p[1][j]
}
