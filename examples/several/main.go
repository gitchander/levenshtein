package main

import (
	"fmt"
	"strings"

	lev "github.com/gitchander/levenshtein"
)

// Levenshtein distance
// https://en.wikipedia.org/wiki/Levenshtein_distance

func main() {
	exampleSamples()
	exampleRunes()
	exampleDistanceCosts()
	examplePrintMatrix()
	exampleRecursive()
	exampleFields()
	exampleBits()
	exampleTrans()
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

func exampleRunes() {
	fmt.Println("Distance between rune slices:")
	var (
		sample = samples[0]

		a = []rune(sample[0])
		b = []rune(sample[1])
	)
	distance := lev.RuneSlices(a, b)
	fmt.Println(distance)
}

func exampleDistanceCosts() {
	fmt.Println("Example Distance Costs:")
	var (
		sample = samples[0]

		a = []rune(sample[0])
		b = []rune(sample[1])
	)
	v := lev.RunePair{a, b}
	cs := lev.Costs{
		MatchCost:   0,
		ReplaceCost: 1,
		DeleteCost:  1,
		InsertCost:  1,
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
	fmt.Println(lev.PrintableMatrix(a, b, cs, ""))
}

func exampleRecursive() {
	fmt.Println("Example NaiveRecursive:")
	var (
		sample = samples[0]

		a = []rune(sample[0])
		b = []rune(sample[1])
	)
	v := lev.RunePair{a, b}
	fmt.Println(lev.NaiveRecursive(v))
}

func exampleFields() {
	fmt.Println("Example Fields:")
	var (
		a = strings.Fields("Computing the Levenshtein, distance is based on the observation that if we reserve")
		b = strings.Fields("Computing the Levenshtein distance- is based on he observation that if we reserve.")
	)
	fmt.Println(lev.StringSlices(a, b))
}

func exampleBits() {
	fmt.Println("Example Bits:")
	var (
		a = parseBits("100101110101010100010111000111011010001010001111101011101011")
		b = parseBits("100101110101010100010111000111011010001010001111101011101010")
	)
	fmt.Println(lev.BoolSlices(a, b))
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

func exampleTrans() {
	var (
		a = []rune("abcdefllgh")
		b = []rune("hello, world!")
	)
	v := lev.RunePair{a, b} // переконайтеся, що поля структури названі правильно
	ts := lev.GetTransformations(v, lev.DefaultCosts)
	fmt.Println(ts)

	offset := 0
	rs := cloneRunes(a)

	for _, t := range ts {
		// Поточний індекс у нашому змінюваному масиві `rs`
		// (t.SourceIdx показує позицію в `а`, а `offset` коригує її з урахуванням попередніх вставок/видалень)
		currentIdx := t.SourceIdx + offset

		switch t.Op {
		case lev.OpMatch:
			// Нічого не робимо, елементи вже збігаються
		case lev.OpReplace:
			// Замінюємо символ у rs на відповідний символ з b
			rs = substituteRune(rs, currentIdx, b[t.TargetIdx])
		case lev.OpDelete:
			// Видаляємо символ за поточним індексом
			rs = deleteRune(rs, currentIdx)
			offset-- // при видаленні наступні елементи зсуваються вліво
		case lev.OpInsert:
			// Вставляємо символ із рядка b за індексом t.TargetIdx
			rs = insertRune(rs, currentIdx, b[t.TargetIdx])
			offset++ // при вставці наступні елементи зсуваються вправо
		}
	}
	fmt.Println(string(rs)) // Очікувано виведе "ome7"
}
