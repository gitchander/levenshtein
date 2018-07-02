package levenshtein

// Weights
type Costs struct {
	DelCost int // Delete cost
	InsCost int // Insert cost
	SubCost int // Substitution cost
}

var DefaultCosts = Costs{
	InsCost: 1,
	DelCost: 1,
	SubCost: 1,
}

type Params struct {
	Costs
	LenA, LenB int
	Match      func(i, j int) bool
}

func (p *Params) calcSubCost(i, j int) int {
	if p.Match(i-1, j-1) {
		return 0
	}
	return p.SubCost
}

func NewParamsRunes(a, b []rune) *Params {
	return &Params{
		Costs: DefaultCosts,
		LenA:  len(a),
		LenB:  len(b),
		Match: func(i, j int) bool {
			return a[i] == b[j]
		},
	}
}

func NewParamsStrings(a, b string) *Params {
	return NewParamsRunes([]rune(a), []rune(b))
}

func Runes(a, b []rune) int {
	return Distance(NewParamsRunes(a, b))
}

func Strings(a, b string) int {
	return Distance(NewParamsStrings(a, b))
}
