package levenshtein

// NaiveRecursive calculates the Levenshtein distance using a naive recursive approach.
//
// WARNING: This function has an exponential time complexity O(3^(m+n)) and does not
// use memoization. It will block on strings longer than 10-15 characters.
// Use Distance instead for production code.
func NaiveRecursive(v Interface) int {
	cs := DefaultCosts
	ni, nj := v.Lens()
	return recursiveDistance(v, cs, ni, nj)
}

func recursiveDistance(v Interface, cs Costs, i, j int) int {

	if j == 0 {
		return i * cs.DeleteCost
	}

	if i == 0 {
		return j * cs.InsertCost
	}

	diagCost := getDiagonalCost(cs, v.Match(i-1, j-1))

	return minInt3(
		(recursiveDistance(v, cs, i-1, j-1) + diagCost),    // (i-1, j-1) - Match or Replace
		(recursiveDistance(v, cs, i-1, j) + cs.DeleteCost), // (i-1, j) - Delete
		(recursiveDistance(v, cs, i, j-1) + cs.InsertCost), // (i, j-1) - Insert
	)
}
