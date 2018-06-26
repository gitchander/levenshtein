package main

import (
	"fmt"

	lev "github.com/gitchander/levenshtein"

	//"github.com/agnivade/levenshtein"
)

// https://en.wikipedia.org/wiki/Levenshtein_distance

func main() {
	sample := samples[3]
	printMatrix(sample.a, sample.b)
	fmt.Println(DistanceMatrix2(sample.a, sample.b))
}

func hw() {
	fmt.Println("Hello, Levenshtein!")
}

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

//func sample1() {
//	s1 := "kitten"
//	s2 := "sitting"
//	distance := levenshtein.ComputeDistance(s1, s2)
//	fmt.Printf("The distance between %s and %s is %d.\n", s1, s2, distance)
//	// Output:
//	// The distance between kitten and sitting is 3.
//}

func Distance(a, b string) int {

	as := []rune(a)
	bs := []rune(b)

	p := lev.Params{
		InsCost: 1,
		DelCost: 1,
		SubCost: 1,
		//		Match: func(i, j int) bool {
		//			return as[i] == bs[j]
		//		},
	}

	return lev.Lev(p, len(as), len(bs))
}

func printMatrix(a, b string) {
	lev.PrintMatrix(lev.DefaultParams, a, b)
}

func DistanceMatrix2(a, b string) int {
	return lev.LevDistance(lev.Strings(a, b))
}
