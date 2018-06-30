package levenshtein

import (
	"bytes"
	"fmt"
)

func Strings(a, b string) int {
	return Distance(NewParamsStr(a, b))
}

func Runes(a, b []rune) int {
	p := &Params{
		Costs: DefaultCosts,
		LenA:  len(a),
		LenB:  len(b),
		Match: func(i, j int) bool {
			return a[i] == b[j]
		},
	}
	return Distance(p)
}

func NewParamsStr(a, b string) *Params {
	var (
		as = []rune(a)
		bs = []rune(b)
	)
	return &Params{
		Costs: DefaultCosts,
		LenA:  len(as),
		LenB:  len(bs),
		Match: func(i, j int) bool {
			return as[i] == bs[j]
		},
	}
}

func PrintMatrix(cs Costs, a, b string, prefix string) string {

	var (
		as = []rune(a)
		bs = []rune(b)
	)
	p := &Params{
		Costs: cs,
		LenA:  len(as),
		LenB:  len(bs),
		Match: func(i, j int) bool {
			return as[i] == bs[j]
		},
	}

	ssd := MakeMatrix(p)

	var (
		ni = p.LenA + 1
		nj = p.LenB + 1
	)

	var cellWidth int
	for _, sd := range ssd {
		for _, d := range sd {
			cellWidth = maxInt2(cellWidth, digitsNumber(d, 10))
		}
	}
	cellWidth++

	var buf bytes.Buffer

	var (
		//		formatRune   = "%[1]*c"
		//		formatNumber = "%[1]*d"

		formatRune   = "%-[1]*c"
		formatNumber = "%-[1]*d"
	)

	empty := fmt.Sprintf(formatRune, cellWidth, '.')

	buf.WriteString(prefix)
	buf.WriteString(empty)
	buf.WriteString(empty)
	for j := 1; j < nj; j++ {
		fmt.Fprintf(&buf, formatRune, cellWidth, bs[j-1])
	}
	buf.WriteByte('\n')

	buf.WriteString(prefix)
	buf.WriteString(empty)
	for j := 0; j < nj; j++ {
		fmt.Fprintf(&buf, formatNumber, cellWidth, ssd[0][j])
	}
	buf.WriteByte('\n')

	for i := 1; i < ni; i++ {
		buf.WriteString(prefix)
		fmt.Fprintf(&buf, formatRune, cellWidth, as[i-1])
		for j := 0; j < nj; j++ {
			fmt.Fprintf(&buf, formatNumber, cellWidth, ssd[i][j])
		}
		buf.WriteByte('\n')
	}

	return string(buf.Bytes())
}

func maxInt2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func digitsNumber(x, base int) int {
	if x == 0 {
		return 1
	}
	var n int
	if x < 0 {
		n++ // sign
		x = -x
	}
	for x > 0 {
		x /= base
		n++
	}
	return n
}
