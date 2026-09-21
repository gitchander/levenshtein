package levenshtein

// Interface defines the requirements for sequence comparison.
type Interface interface {
	Lens() (ni, nj int)
	Match(i, j int) bool
}

//------------------------------------------------------------------------------

// RunePair wraps a pair of rune slices for comparison.
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

// StringPair wraps a pair of string slices for comparison.
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

// BoolPair wraps a pair of boolean slices for comparison.
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

type BytePair [2][]byte

var _ Interface = BytePair{}

func (p BytePair) Lens() (ni, nj int) {
	return len(p[0]), len(p[1])
}

func (p BytePair) Match(i, j int) bool {
	return p[0][i] == p[1][j]
}

//------------------------------------------------------------------------------
