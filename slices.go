package levenshtein

//------------------------------------------------------------------------------
type RuneSlices [2][]rune

var _ Interface = RuneSlices{}

func (p RuneSlices) Len(k int) int {
	return len(p[k])
}

func (p RuneSlices) Match(i, j int) bool {
	return p[0][i] == p[1][j]
}

//------------------------------------------------------------------------------
type StringSlices [2][]string

var _ Interface = StringSlices{}

func (p StringSlices) Len(k int) int {
	return len(p[k])
}

func (p StringSlices) Match(i, j int) bool {
	return p[0][i] == p[1][j]
}

//------------------------------------------------------------------------------
type BoolSlices [2][]bool

var _ Interface = BoolSlices{}

func (p BoolSlices) Len(k int) int {
	return len(p[k])
}

func (p BoolSlices) Match(i, j int) bool {
	return p[0][i] == p[1][j]
}

//------------------------------------------------------------------------------
func Runes(a, b []rune) int {
	return Distance(RuneSlices{a, b})
}

func Strings(a, b string) int {
	return Distance(RuneSlices{[]rune(a), []rune(b)})
}

func Bools(a, b []bool) int {
	return Distance(BoolSlices{a, b})
}
