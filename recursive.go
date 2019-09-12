package levenshtein

func Recursive(v Interface) int {
	cs := DefaultCosts
	return recursiveDistance(v, &cs, v.Len(0), v.Len(1))
}

func recursiveDistance(v Interface, cs *Costs, i, j int) int {
	if j == 0 {
		return i * cs.DelCost
	}
	if i == 0 {
		return j * cs.InsCost
	}
	return minInt3(
		recursiveDistance(v, cs, i-1, j)+cs.DelCost,                             // (i-1, j) - Delete
		recursiveDistance(v, cs, i, j-1)+cs.InsCost,                             // (i, j-1) - Insert
		recursiveDistance(v, cs, i-1, j-1)+calcSubCost(v, i-1, j-1, cs.SubCost), // (i-1, j-1) - Substitution
	)
}
