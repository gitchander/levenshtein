package main

import (
	"fmt"
	"strings"

	lev "github.com/gitchander/levenshtein"
)

func main() {
	var (
		line1 = "one two three four"
		line2 = "one two three"
	)
	var (
		a = strings.Split(line1, " ")
		b = strings.Split(line2, " ")
	)
	distance := lev.Distance(lev.StringSlices{a, b})
	fmt.Printf("the levenshtein distance between %q and %q = %d\n", a, b, distance)
}
