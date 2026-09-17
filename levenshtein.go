package levenshtein

// Levenshtein distance

// https://en.wikipedia.org/wiki/Levenshtein_distance

type Interface interface {
	Lens() (ni, nj int)
	Match(i, j int) bool
}

// Weights
type Costs struct {
	InsCost int // Insert cost
	DelCost int // Delete cost
	SubCost int // Substitution cost
}

var DefaultCosts = Costs{
	InsCost: 1,
	DelCost: 1,
	SubCost: 1,
}

func Distance(v Interface) int {
	return DistanceCosts(v, DefaultCosts)
}

func DistanceCosts(v Interface, cs Costs) int {
	ni, nj := v.Lens()
	if ni < nj {
		return distanceByLen0(v, cs)
	}
	return distanceByLen1(v, cs)
}

func distanceByLen0(v Interface, cs Costs) int {

	ni, nj := v.Lens()

	row := make([]int, (ni + 1))
	for i := 0; i <= ni; i++ {
		row[i] = i * cs.DelCost
	}

	for j := 1; j <= nj; j++ {
		prevDiag := row[0]
		row[0] = j * cs.InsCost
		for i := 1; i <= ni; i++ {

			// (i-1, j) - Delete
			delCost := row[i-1] + cs.DelCost

			// (i, j-1) - Insert
			insCost := row[i] + cs.InsCost

			// (i-1, j-1) - Substitution
			subCost := prevDiag
			if !(v.Match(i-1, j-1)) {
				subCost += cs.SubCost
			}

			prevDiag = row[i]
			row[i] = minInt3(delCost, insCost, subCost)
		}
	}
	return row[ni]
}

func distanceByLen1(v Interface, cs Costs) int {

	ni, nj := v.Lens()

	row := make([]int, (nj + 1))
	for j := 0; j <= nj; j++ {
		row[j] = j * cs.InsCost
	}

	for i := 1; i <= ni; i++ {
		prevDiag := row[0]
		row[0] = i * cs.DelCost
		for j := 1; j <= nj; j++ {

			// (i-1, j) - Delete
			delCost := row[j] + cs.DelCost

			// (i, j-1) - Insert
			insCost := row[j-1] + cs.InsCost

			// (i-1, j-1) - Substitution
			subCost := prevDiag
			if !(v.Match(i-1, j-1)) {
				subCost += cs.SubCost
			}

			prevDiag = row[j]
			row[j] = minInt3(delCost, insCost, subCost)
		}
	}
	return row[nj]
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
