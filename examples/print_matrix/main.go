package main

import (
	"fmt"

	lev "github.com/gitchander/levenshtein"
)

func main() {
	var (
		a = []rune("sitting")
		b = []rune("kitten")
	)
	costs := lev.DefaultCosts
	fmt.Print(lev.PrintableMatrix(costs, a, b, ""))
}
