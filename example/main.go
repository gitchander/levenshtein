package main

import (
	"fmt"
	"strings"

	lev "github.com/gitchander/levenshtein"
)

// https://en.wikipedia.org/wiki/Levenshtein_distance

func main() {
	exampleSamples()
	exampleParams()
	examplePrintMatrix()
	exampleRecursive()
	exampleFields()
	exampleBits()
}

var samples = [][2]string{
	{
		"sitting",
		"kitten",
	},
	{
		"sunday",
		"saturday",
	},
	{
		"exponential",
		"polynomial",
	},
	{
		"Population",
		"Education",
	},
	{
		"industry",
		"interest",
	},
	{
		"soylent green is people",
		"people soiled our green",
	},
}

func exampleSamples() {
	fmt.Println("Example Samples:")
	for _, sample := range samples {
		var (
			a = sample[0]
			b = sample[1]
		)
		fmt.Println(lev.Strings(a, b))
	}
}

func exampleParams() {
	fmt.Println("Example Params:")
	var (
		sample = samples[0]

		a = sample[0]
		b = sample[1]
	)

	costs := lev.Costs{
		DelCost: 1,
		InsCost: 1,
		SubCost: 1,
	}

	var (
		as = []rune(a)
		bs = []rune(b)
	)
	p := &lev.Params{
		Costs: costs,
		LenA:  len(as),
		LenB:  len(bs),
		Match: func(i, j int) bool {
			return as[i] == bs[j]
		},
	}

	fmt.Println(lev.Distance(p))
}

func examplePrintMatrix() {
	fmt.Println("Example Print Matrix:")
	costs := lev.Costs{
		DelCost: 1,
		InsCost: 1,
		SubCost: 1,
	}
	var (
		sample = samples[2]

		a = []rune(sample[0])
		b = []rune(sample[1])
	)
	fmt.Println(lev.PrintableMatrix(costs, a, b, ""))
}

func exampleRecursive() {
	fmt.Println("Example Recursive:")
	var (
		sample = samples[0]

		a = sample[0]
		b = sample[1]
	)
	p := lev.NewParamsStrings(a, b)
	fmt.Println(lev.Recursive(p))
}

func exampleFields() {
	fmt.Println("Example Fields:")
	var (
		a = strings.Fields("Computing the Levenshtein, distance is based on the observation that if we reserve")
		b = strings.Fields("Computing the Levenshtein distance- is based on he observation that if we reserve.")
	)
	p := &lev.Params{
		Costs: lev.DefaultCosts,
		LenA:  len(a),
		LenB:  len(b),
		Match: func(i, j int) bool {
			return a[i] == b[j]
		},
	}
	fmt.Println(lev.Distance(p))
}

func exampleBits() {
	fmt.Println("Example Bits:")
	var (
		a = parseBits("100101110101010100010111000111011010001010001111101011101011")
		b = parseBits("100101110101010100010111000111011010001010001111101011101010")
	)
	p := &lev.Params{
		Costs: lev.DefaultCosts,
		LenA:  len(a),
		LenB:  len(b),
		Match: func(i, j int) bool {
			return a[i] == b[j]
		},
	}
	fmt.Println(lev.Distance(p))
}

func parseBits(s string) []bool {
	data := []byte(s)
	bs := make([]bool, len(data))
	for i, b := range data {
		switch b {
		case '0':
			bs[i] = false
		case '1':
			bs[i] = true
		default:
			panic("invalid bits")
		}
	}
	return bs
}
