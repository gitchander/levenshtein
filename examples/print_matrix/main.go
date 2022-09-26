package main

import (
	"fmt"

	lev "github.com/gitchander/levenshtein"
)

// https://en.wikipedia.org/wiki/Levenshtein_distance

func main() {
	test1()
	test2()
}

func test1() {
	var (
		a = []rune("sitting")
		b = []rune("kitten")
	)
	costs := lev.DefaultCosts
	fmt.Println(lev.PrintableMatrix(costs, a, b, ""))
}

func test2() {
	var (
		a = []rune("Sunday")
		b = []rune("Saturday")
	)
	costs := lev.DefaultCosts
	fmt.Println(lev.PrintableMatrix(costs, a, b, ""))
}
