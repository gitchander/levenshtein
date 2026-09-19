package levenshtein

// +------------------+------------------+
// | (i-1, j-1) - Sub | (i-1, j) - Del   |
// +------------------+------------------+
// | (i, j-1) - Ins   | (i, j)           |
// +------------------+------------------+

// MakeMatrix generates the full dynamic programming matrix for testing and analysis.
func MakeMatrix(v Interface, cs Costs) [][]int {

	ni, nj := v.Lens()

	m := make([][]int, (ni + 1))
	for i := 0; i <= ni; i++ {
		m[i] = make([]int, (nj + 1))
	}

	for i := 0; i <= ni; i++ {
		m[i][0] = i * cs.DelCost
	}
	for j := 0; j <= nj; j++ {
		m[0][j] = j * cs.InsCost
	}

	for i := 1; i <= ni; i++ {
		for j := 1; j <= nj; j++ {

			// (i-1, j) - Delete
			delCost := m[i-1][j] + cs.DelCost

			// (i, j-1) - Insert
			insCost := m[i][j-1] + cs.InsCost

			// (i-1, j-1) - Substitution
			subCost := m[i-1][j-1]
			if !(v.Match(i-1, j-1)) {
				subCost += cs.SubCost
			}

			m[i][j] = minInt3(delCost, insCost, subCost)
		}
	}

	return m
}
