package levenshtein

// (alternative names: Weights)
type Costs struct {
	DelCost int // Delete cost
	InsCost int // Insert cost
	SubCost int // Substitution cost
}

// DefaultCosts returns the standard operation costs (1 for all operations).
var DefaultCosts = Costs{
	DelCost: 1,
	InsCost: 1,
	SubCost: 1,
}
