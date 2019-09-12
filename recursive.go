package levenshtein

func Recursive(v Interface) int {
	return recursiveDistance(v, v.Len(0), v.Len(1))
}

func recursiveDistance(v Interface, i, j int) int {
	cs := DefaultCosts
	if j == 0 {
		return i * cs.DelCost
	}
	if i == 0 {
		return j * cs.InsCost
	}
	return minInt3(
		recursiveDistance(v, i-1, j)+cs.DelCost,                 // (i-1, j) - Delete
		recursiveDistance(v, i, j-1)+cs.InsCost,                 // (i, j-1) - Insert
		recursiveDistance(v, i-1, j-1)+calcSubCost(v, i, j, cs), // (i-1, j-1) - Substitution
	)
}
