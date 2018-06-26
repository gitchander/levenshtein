package levenshtein

import (
	"bytes"
	"fmt"
)

type Params struct {
	DelCost int // Delete. Стоимость удаления
	InsCost int // Insert. Стоимость вставки.
	SubCost int // Substitution. Стоимость замены/замещения.

	Match func(i, j int) bool
}

var DefaultParams = Params{
	InsCost: 1,
	DelCost: 1,
	SubCost: 1,
}

// Recursive
func Lev(p Params, i, j int) int {

	if i == 0 {
		return j
	}
	if j == 0 {
		return i
	}

	var subCost int
	if !p.Match(i-1, j-1) {
		subCost = p.SubCost
	}

	return minInt3(
		Lev(p, i-1, j)+p.DelCost, // (i-1, j) - Delete
		Lev(p, i, j-1)+p.InsCost, // (i, j-1) - Insert
		Lev(p, i-1, j-1)+subCost, // (i-1, j-1) - Substitution
	)
}

type Values interface {
	Size() (ni, nj int)
	//BothLen() (ni, nj int)

	Match(i, j int) bool
}

func Strings(a, b string) Values {
	return &valStrings{
		a: []rune(a),
		b: []rune(b),
	}
}

type valStrings struct {
	a, b []rune
}

func (p *valStrings) Size() (ni, nj int)  { return len(p.a), len(p.b) }
func (p *valStrings) Match(i, j int) bool { return p.a[i] == p.b[j] }

func PrintMatrix(p Params, a, b string) {

	var (
		as = []rune(a)
		bs = []rune(b)
	)

	var (
		ni = len(as) + 1
		nj = len(bs) + 1
	)

	var ssd = make([][]int, ni)
	for i := 0; i < ni; i++ {
		ssd[i] = make([]int, nj)
	}

	for i := 0; i < ni; i++ {
		ssd[i][0] = i * p.DelCost
	}
	for j := 0; j < nj; j++ {
		ssd[0][j] = j * p.InsCost
	}

	for i := 1; i < ni; i++ {
		for j := 1; j < nj; j++ {
			var subCost int
			if as[i-1] != bs[j-1] {
				subCost = p.SubCost
			}
			ssd[i][j] = minInt3(
				ssd[i-1][j]+p.DelCost, // (i-1, j) - Delete
				ssd[i][j-1]+p.InsCost, // (i, j-1) - Insert
				ssd[i-1][j-1]+subCost, // (i-1, j-1) - Substitution
			)
		}
	}

	prefix := ""
	var buf bytes.Buffer

	d := digitsNumber(maxInt2(ni*p.DelCost, nj*p.InsCost), 10) + 1

	//empty := string(repeatByte(' ', d))
	empty := fmt.Sprintf("%-[1]*c", d, '.')

	buf.WriteString(prefix)
	buf.WriteString(empty)
	buf.WriteString(empty)
	for j := 1; j < nj; j++ {
		fmt.Fprintf(&buf, "%-[1]*c", d, bs[j-1])
	}
	buf.WriteByte('\n')

	buf.WriteString(prefix)
	buf.WriteString(empty)
	for j := 0; j < nj; j++ {
		fmt.Fprintf(&buf, "%-[1]*d", d, ssd[0][j])
	}
	buf.WriteByte('\n')

	for i := 1; i < ni; i++ {
		buf.WriteString(prefix)
		fmt.Fprintf(&buf, "%-[1]*c", d, as[i-1])
		for j := 0; j < nj; j++ {
			fmt.Fprintf(&buf, "%-[1]*d", d, ssd[i][j])
		}
		buf.WriteByte('\n')
	}

	fmt.Println(buf.String())
}

func LevDistance(h Values) int {

	var (
		lenA, lenB = h.Size()

		ni = lenA + 1
		nj = lenB + 1
	)

	return levDistance_1(DefaultParams, h, ni, nj)
}

func levDistance_1(p Params, h Values, ni, nj int) int {

	vs := make([]int, nj)
	for j := 0; j < nj; j++ {
		vs[j] = j * p.InsCost
	}
	printVS(vs)

	for i := 1; i < ni; i++ {
		vj := vs[0]
		vs[0] = i * p.DelCost
		for j := 1; j < nj; j++ {
			var subCost int
			if !h.Match(i-1, j-1) {
				subCost = p.SubCost
			}
			temp := vs[j]
			vs[j] = minInt3(
				vs[j]+p.DelCost,   // (i-1, j) - Delete
				vs[j-1]+p.InsCost, // (i, j-1) - Insert
				vj+subCost,        // (i-1, j-1) - Substitution
			)
			vj = temp
		}
		printVS(vs)
	}

	return vs[nj-1]
}

func levDistance_2(p Params, h Values, ni, nj int) int {

	vs := make([]int, ni)
	for i := 0; i < ni; i++ {
		vs[i] = i * p.DelCost
	}
	printVS(vs)

	for j := 1; j < nj; j++ {
		vi := vs[0]
		vs[0] = j * p.InsCost
		for i := 1; i < ni; i++ {
			var subCost int
			if !h.Match(i-1, j-1) {
				subCost = p.SubCost
			}
			temp := vs[i]
			vs[i] = minInt3(
				vs[i-1]+p.DelCost, // (i-1, j) - Delete
				vs[i]+p.InsCost,   // (i, j-1) - Insert
				vi+subCost,        // (i-1, j-1) - Substitution
			)
			vi = temp
		}
		printVS(vs)
	}

	return vs[ni-1]
}

func printVS(vs []int) {
	prefix := "   "
	fmt.Print(prefix)
	for _, v := range vs {
		fmt.Printf("%-3d", v)
	}
	fmt.Println()
}

// Distance
func Distance(v Values) int {

	return 0
}
