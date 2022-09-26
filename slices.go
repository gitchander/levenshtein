package levenshtein

//------------------------------------------------------------------------------
type RuneSlices [2][]rune

var _ Interface = RuneSlices{}

func (p RuneSlices) Lens() (ni, nj int) {
	return len(p[0]), len(p[1])
}

func (p RuneSlices) Match(i, j int) bool {
	return p[0][i] == p[1][j]
}

//------------------------------------------------------------------------------
type StringSlices [2][]string

var _ Interface = StringSlices{}

func (p StringSlices) Lens() (ni, nj int) {
	return len(p[0]), len(p[1])
}

func (p StringSlices) Match(i, j int) bool {
	return p[0][i] == p[1][j]
}

//------------------------------------------------------------------------------
type BoolSlices [2][]bool

var _ Interface = BoolSlices{}

func (p BoolSlices) Lens() (ni, nj int) {
	return len(p[0]), len(p[1])
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
