package levenshtein

//------------------------------------------------------------------------------

// RunePair (alternative names: RuneSlices, RunesPair, RuneSlicePair)

type RunePair [2][]rune

var _ Interface = RunePair{}

func (p RunePair) Lens() (ni, nj int) {
	return len(p[0]), len(p[1])
}

func (p RunePair) Match(i, j int) bool {
	return p[0][i] == p[1][j]
}

//------------------------------------------------------------------------------

// StringPair (alternative names: StringSlices, StringsPair, StringSlicePair)

type StringPair [2][]string

var _ Interface = StringPair{}

func (p StringPair) Lens() (ni, nj int) {
	return len(p[0]), len(p[1])
}

func (p StringPair) Match(i, j int) bool {
	return p[0][i] == p[1][j]
}

//------------------------------------------------------------------------------

// BoolPair (alternative names: BoolSlices, BoolsPair, BoolSlicePair)

type BoolPair [2][]bool

var _ Interface = BoolPair{}

func (p BoolPair) Lens() (ni, nj int) {
	return len(p[0]), len(p[1])
}

func (p BoolPair) Match(i, j int) bool {
	return p[0][i] == p[1][j]
}

//------------------------------------------------------------------------------

func RuneSlices(a, b []rune) int {
	return Distance(RunePair{a, b})
}

func StringSlices(a, b []string) int {
	return Distance(StringPair{a, b})
}

func BoolSlices(a, b []bool) int {
	return Distance(BoolPair{a, b})
}

//------------------------------------------------------------------------------

func Strings(a, b string) int {
	return RuneSlices([]rune(a), []rune(b))
}

//------------------------------------------------------------------------------
