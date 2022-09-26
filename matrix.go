package levenshtein

func MakeMatrix(v Interface, cs Costs) [][]int {

	ni, nj := v.Lens()
	ni, nj = ni+1, nj+1

	var ssd = make([][]int, ni)
	for i := 0; i < ni; i++ {
		ssd[i] = make([]int, nj)
	}

	for i := 0; i < ni; i++ {
		ssd[i][0] = i * cs.DelCost
	}
	for j := 0; j < nj; j++ {
		ssd[0][j] = j * cs.InsCost
	}

	for i := 1; i < ni; i++ {
		for j := 1; j < nj; j++ {

			// (i-1, j) - Delete
			delCost := ssd[i-1][j] + cs.DelCost

			// (i, j-1) - Insert
			insCost := ssd[i][j-1] + cs.InsCost

			// (i-1, j-1) - Substitution
			subCost := ssd[i-1][j-1]
			if !(v.Match(i-1, j-1)) {
				subCost += cs.SubCost
			}

			ssd[i][j] = minInt3(delCost, insCost, subCost)
		}
	}

	return ssd
}
