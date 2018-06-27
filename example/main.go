package main

import (
	"fmt"
	"strings"

	lev "github.com/gitchander/levenshtein"
)

// https://en.wikipedia.org/wiki/Levenshtein_distance

func main() {
	exampleStrings()
	exampleWords()
	exampleBits()
}

func exampleStrings() {
	var samples = []struct {
		a, b     string
		distance int
	}{
		{
			a:        ``,
			b:        ``,
			distance: 0,
		},
		{
			a:        "sitting",
			b:        "kitten",
			distance: 3,
		},
		{
			a:        "sunday",
			b:        "saturday",
			distance: 3,
		},
		{
			a:        "exponential",
			b:        "polynomial",
			distance: 6,
		},
	}

	var (
		sample = samples[3]

		a = sample.a
		b = sample.b
	)

	printSample(a, b)

	fmt.Println(lev.Distance(lev.Strings(a, b)))
}

func printSample(a, b string) {
	//pl := lev.DefaultPriceList
	pl := lev.PriceList{
		DelCost: 1,
		InsCost: 1,
		SubCost: 1,
	}
	fmt.Println(lev.PrintMatrix(pl, a, b, ""))

	var (
		as = []rune(a)
		bs = []rune(b)
	)
	p := &lev.Params{
		PriceList: pl,
		LenA:      len(as),
		LenB:      len(bs),
		Match: func(i, j int) bool {
			return as[i] == bs[j]
		},
	}
	fmt.Println(lev.Recursive(p))
	fmt.Println(lev.Distance(p))
}

func exampleWords() {
	var (
		a = strings.Split("Computing the Levenshtein distance is based on the observation that if we reserve", " ")
		b = strings.Split("Computing the Levenshtein distance- is based on he observation that if we reserve.", " ")
	)
	p := &lev.Params{
		PriceList: lev.DefaultPriceList,
		LenA:      len(a),
		LenB:      len(b),
		Match: func(i, j int) bool {
			return a[i] == b[j]
		},
	}
	fmt.Println("Example Words:")
	fmt.Println(lev.Distance(p))
}

func exampleBits() {
	var (
		a = parseBits("100101110101010100010111000111011010001010001111101011101011")
		b = parseBits("100101110101010100010111000111011010001010001111101011101010")
	)
	p := &lev.Params{
		PriceList: lev.DefaultPriceList,
		LenA:      len(a),
		LenB:      len(b),
		Match: func(i, j int) bool {
			return a[i] == b[j]
		},
	}
	fmt.Println("Example Bits:")
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
