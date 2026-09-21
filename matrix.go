package levenshtein

// MakeMatrix generates the full dynamic programming matrix for testing and analysis.
func MakeMatrix(v Interface, cs Costs) [][]int {

	ni, nj := v.Lens()

	matrix := make([][]int, (ni + 1))
	for i := 0; i <= ni; i++ {
		matrix[i] = make([]int, (nj + 1))
	}

	for i := 0; i <= ni; i++ {
		matrix[i][0] = i * cs.DeleteCost
	}
	for j := 0; j <= nj; j++ {
		matrix[0][j] = j * cs.InsertCost
	}

	for i := 1; i <= ni; i++ {
		for j := 1; j <= nj; j++ {

			diagCost := getDiagonalCost(cs, v.Match(i-1, j-1))

			matrix[i][j] = minInt3(
				(matrix[i-1][j-1] + diagCost),    // (i-1, j-1) - Match or Replace
				(matrix[i-1][j] + cs.DeleteCost), // (i-1, j) - Delete
				(matrix[i][j-1] + cs.InsertCost), // (i, j-1) - Insert
			)
		}
	}

	return matrix
}
