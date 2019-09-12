package levenshtein

// Levenshtein distance

// https://en.wikipedia.org/wiki/Levenshtein_distance

type Interface interface {
	Len(k int) int // k = [0, 1]
	Match(i, j int) bool
}

func Distance(v Interface) int {
	return DistanceCosts(v, DefaultCosts)
}

func DistanceCosts(v Interface, cs Costs) int {
	if v.Len(0) < v.Len(1) {
		return distanceByLen0(v, cs)
	}
	return distanceByLen1(v, cs)
}

func distanceByLen0(v Interface, cs Costs) int {
	var (
		ni = v.Len(0) + 1
		nj = v.Len(1) + 1
	)
	vis := make([]int, ni)
	for i := 0; i < ni; i++ {
		vis[i] = i * cs.DelCost
	}
	for j := 1; j < nj; j++ {
		vi := vis[0]
		vis[0] = j * cs.InsCost
		for i := 1; i < ni; i++ {
			temp := vis[i]
			vis[i] = minInt3(
				vis[i-1]+cs.DelCost,                     // (i-1, j) - Delete
				vis[i]+cs.InsCost,                       // (i, j-1) - Insert
				vi+calcSubCost(v, i-1, j-1, cs.SubCost), // (i-1, j-1) - Substitution
			)
			vi = temp
		}
	}
	return vis[ni-1]
}

func distanceByLen1(v Interface, cs Costs) int {
	var (
		ni = v.Len(0) + 1
		nj = v.Len(1) + 1
	)
	vjs := make([]int, nj)
	for j := 0; j < nj; j++ {
		vjs[j] = j * cs.InsCost
	}
	for i := 1; i < ni; i++ {
		vj := vjs[0]
		vjs[0] = i * cs.DelCost
		for j := 1; j < nj; j++ {
			temp := vjs[j]
			vjs[j] = minInt3(
				vjs[j]+cs.DelCost,                       // (i-1, j) - Delete
				vjs[j-1]+cs.InsCost,                     // (i, j-1) - Insert
				vj+calcSubCost(v, i-1, j-1, cs.SubCost), // (i-1, j-1) - Substitution
			)
			vj = temp
		}
	}
	return vjs[nj-1]
}

func calcSubCost(v Interface, i, j int, subCost int) int {
	if v.Match(i, j) {
		return 0
	}
	return subCost
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
