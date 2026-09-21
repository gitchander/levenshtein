package levenshtein

// (alternative names: Weights)
type Costs struct {
	MatchCost   int // Match cost (0)
	ReplaceCost int // Substitution cost
	DeleteCost  int // Delete cost
	InsertCost  int // Insert cost
}

// DefaultCosts returns the standard operation costs (1 for all operations).
var DefaultCosts = Costs{
	MatchCost:   0,
	ReplaceCost: 1,
	DeleteCost:  1,
	InsertCost:  1,
}

func getDiagonalCost(cs Costs, match bool) int {
	if match {
		return cs.MatchCost
	}
	return cs.ReplaceCost
}
