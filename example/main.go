package main

import (
	"fmt"
	"strings"

	lev "github.com/gitchander/levenshtein"
)

// https://en.wikipedia.org/wiki/Levenshtein_distance

func main() {
	exampleSamples()
	exampleDistanceCosts()
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

func exampleDistanceCosts() {
	fmt.Println("Example Distance Costs:")
	var (
		sample = samples[0]

		a = []rune(sample[0])
		b = []rune(sample[1])
	)
	v := lev.RuneSlices{a, b}
	cs := lev.Costs{
		DelCost: 1,
		InsCost: 1,
		SubCost: 1,
	}
	fmt.Println(lev.DistanceCosts(v, cs))
}

func examplePrintMatrix() {
	fmt.Println("Example Print Matrix:")
	cs := lev.DefaultCosts
	var (
		sample = samples[0]

		a = []rune(sample[0])
		b = []rune(sample[1])
	)
	fmt.Println(lev.PrintableMatrix(cs, a, b, ""))
}

func exampleRecursive() {
	fmt.Println("Example Recursive:")
	var (
		sample = samples[0]

		a = []rune(sample[0])
		b = []rune(sample[1])
	)
	v := lev.RuneSlices{a, b}
	fmt.Println(lev.Recursive(v))
}

func exampleFields() {
	fmt.Println("Example Fields:")
	var (
		a = strings.Fields("Computing the Levenshtein, distance is based on the observation that if we reserve")
		b = strings.Fields("Computing the Levenshtein distance- is based on he observation that if we reserve.")
	)
	v := lev.StringSlices{a, b}
	fmt.Println(lev.Distance(v))
}

func exampleBits() {
	fmt.Println("Example Bits:")
	var (
		a = parseBits("100101110101010100010111000111011010001010001111101011101011")
		b = parseBits("100101110101010100010111000111011010001010001111101011101010")
	)
	v := lev.BoolSlices{a, b}
	fmt.Println(lev.Distance(v))
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
