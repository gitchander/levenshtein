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
