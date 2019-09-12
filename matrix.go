package levenshtein

func MakeMatrix(v Interface, cs Costs) [][]int {

	var (
		ni = v.Len(0) + 1
		nj = v.Len(1) + 1
	)

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
			ssd[i][j] = minInt3(
				ssd[i-1][j]+cs.DelCost,                             // (i-1, j) - Delete
				ssd[i][j-1]+cs.InsCost,                             // (i, j-1) - Insert
				ssd[i-1][j-1]+calcSubCost(v, i-1, j-1, cs.SubCost), // (i-1, j-1) - Substitution
			)
		}
	}

	return ssd
}
