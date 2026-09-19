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
		a = strings.Fields(line1)
		b = strings.Fields(line2)
	)
	distance := lev.StringSlices(a, b)
	fmt.Printf("the levenshtein distance = %d\n", distance)
}
