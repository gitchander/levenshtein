package levenshtein

func Recursive(v Interface) int {
	cs := DefaultCosts
	ni, nj := v.Lens()
	return recursiveDistance(v, &cs, ni, nj)
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
