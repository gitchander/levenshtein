package main

import (
	"fmt"

	lev "github.com/gitchander/levenshtein"
)

// https://en.wikipedia.org/wiki/Levenshtein_distance

func main() {
	printSamples()
}

type sample struct {
	a, b  []rune
	costs lev.Costs
}

func printSamples() {
	samples := []sample{
		{
			a:     []rune("sitting"),
			b:     []rune("kitten"),
			costs: lev.DefaultCosts,
		},
		{
			a:     []rune("Sunday"),
			b:     []rune("Saturday"),
			costs: lev.DefaultCosts,
		},
		{
			a:     []rune("hello"),
			b:     []rune("world"),
			costs: lev.DefaultCosts,
		},
		{
			a: []rune("AACGCA"),
			b: []rune("GAGCTA"),
			costs: lev.Costs{
				MatchCost:   0,
				ReplaceCost: 2,
				DeleteCost:  1,
				InsertCost:  1,
			},
		},
		{
			a:     []rune("exponential"),
			b:     []rune("polynomial"),
			costs: lev.DefaultCosts,
		},
	}
	for _, sample := range samples {
		fmt.Println(lev.PrintableMatrix(sample.a, sample.b, sample.costs, ""))
	}
}
