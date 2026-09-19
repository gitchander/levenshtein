package levenshtein

// Levenshtein distance

// https://en.wikipedia.org/wiki/Levenshtein_distance

type Interface interface {
	Lens() (ni, nj int)
	Match(i, j int) bool
}

// Weights
type Costs struct {
	DelCost int // Delete cost
	InsCost int // Insert cost
	SubCost int // Substitution cost
}

var DefaultCosts = Costs{
	DelCost: 1,
	InsCost: 1,
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

	buf := make([]int, (ni + 1))
	for i := 0; i <= ni; i++ {
		buf[i] = i * cs.DelCost
	}

	for j := 1; j <= nj; j++ {
		prevDiag := buf[0]
		buf[0] = j * cs.InsCost
		for i := 1; i <= ni; i++ {

			// (i-1, j) - Delete
			delCost := buf[i-1] + cs.DelCost

			// (i, j-1) - Insert
			insCost := buf[i] + cs.InsCost

			// (i-1, j-1) - Substitution
			subCost := prevDiag
			if !(v.Match(i-1, j-1)) {
				subCost += cs.SubCost
			}

			prevDiag = buf[i]
			buf[i] = minInt3(delCost, insCost, subCost)
		}
	}
	return buf[ni]
}

func distanceByLen1(v Interface, cs Costs) int {

	ni, nj := v.Lens()

	buf := make([]int, (nj + 1))
	for j := 0; j <= nj; j++ {
		buf[j] = j * cs.InsCost
	}

	for i := 1; i <= ni; i++ {
		prevDiag := buf[0]
		buf[0] = i * cs.DelCost
		for j := 1; j <= nj; j++ {

			// (i-1, j) - Delete
			delCost := buf[j] + cs.DelCost

			// (i, j-1) - Insert
			insCost := buf[j-1] + cs.InsCost

			// (i-1, j-1) - Substitution
			subCost := prevDiag
			if !(v.Match(i-1, j-1)) {
				subCost += cs.SubCost
			}

			prevDiag = buf[j]
			buf[j] = minInt3(delCost, insCost, subCost)
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
