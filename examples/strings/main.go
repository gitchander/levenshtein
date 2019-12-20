package main

import (
	"fmt"

	lev "github.com/gitchander/levenshtein"
)

func main() {
	var (
		a = "exponential"
		b = "polynomial"
	)
	distance := lev.Strings(a, b)
	fmt.Printf("the levenshtein distance between %q and %q = %d\n", a, b, distance)
}
