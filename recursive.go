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

	// (i-1, j) - Delete
	delCost := recursiveDistance(v, cs, i-1, j) + cs.DelCost

	// (i, j-1) - Insert
	insCost := recursiveDistance(v, cs, i, j-1) + cs.InsCost

	// (i-1, j-1) - Substitution
	subCost := recursiveDistance(v, cs, i-1, j-1)
	if !(v.Match(i-1, j-1)) {
		subCost += cs.SubCost
	}

	return minInt3(delCost, insCost, subCost)
}
