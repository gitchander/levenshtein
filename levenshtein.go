package levenshtein

// Levenshtein distance

// https://en.wikipedia.org/wiki/Levenshtein_distance

// Distance calculates the Levenshtein distance using default costs.
func Distance(v Interface) int {
	return DistanceCosts(v, DefaultCosts)
}

// DistanceCosts calculates the Levenshtein distance using custom costs.
func DistanceCosts(v Interface, cs Costs) int {
	ni, nj := v.Lens()
	if ni < nj {
		return distanceByLen0(v, cs)
	}
	return distanceByLen1(v, cs)
}

func distanceByLen0(v Interface, cs Costs) int {

	ni, nj := v.Lens()

	buf := make([]int, (ni + 1))
	for i := 0; i <= ni; i++ {
		buf[i] = i * cs.DeleteCost
	}

	for j := 1; j <= nj; j++ {
		prevDiag := buf[0]
		buf[0] = j * cs.InsertCost
		for i := 1; i <= ni; i++ {

			diagCost := getDiagonalCost(cs, v.Match(i-1, j-1))

			newValue := minInt3(
				(prevDiag + diagCost),      // (i-1, j-1) - Match or Replace
				(buf[i-1] + cs.DeleteCost), // (i-1, j) - Delete
				(buf[i] + cs.InsertCost),   // (i, j-1) - Insert
			)

			prevDiag = buf[i]
			buf[i] = newValue
		}
	}
	return buf[ni]
}

func distanceByLen1(v Interface, cs Costs) int {

	ni, nj := v.Lens()

	buf := make([]int, (nj + 1))
	for j := 0; j <= nj; j++ {
		buf[j] = j * cs.InsertCost
	}

	for i := 1; i <= ni; i++ {
		prevDiag := buf[0]
		buf[0] = i * cs.DeleteCost
		for j := 1; j <= nj; j++ {

			diagCost := getDiagonalCost(cs, v.Match(i-1, j-1))

			newValue := minInt3(
				(prevDiag + diagCost),      // (i-1, j-1) - Match or Replace
				(buf[j] + cs.DeleteCost),   // (i-1, j) - Delete
				(buf[j-1] + cs.InsertCost), // (i, j-1) - Insert
			)

			prevDiag = buf[j]
			buf[j] = newValue
		}
	}
	return buf[nj]
}

func minInt3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
