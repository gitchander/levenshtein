package levenshtein

// It works too slowly
func Recursive(p *Params) int {
	return recursiveDistance(p, p.LenA, p.LenB)
}

func recursiveDistance(p *Params, i, j int) int {
	if j == 0 {
		return i * p.DelCost
	}
	if i == 0 {
		return j * p.InsCost
	}
	return minInt3(
		recursiveDistance(p, i-1, j)+p.DelCost,              // (i-1, j) - Delete
		recursiveDistance(p, i, j-1)+p.InsCost,              // (i, j-1) - Insert
		recursiveDistance(p, i-1, j-1)+calcSubCost(p, i, j), // (i-1, j-1) - Substitution
	)
}

func MakeMatrix(p *Params) [][]int {

	var (
		ni = p.LenA + 1
		nj = p.LenB + 1
	)

	var ssd = make([][]int, ni)
	for i := 0; i < ni; i++ {
		ssd[i] = make([]int, nj)
	}

	for i := 0; i < ni; i++ {
		ssd[i][0] = i * p.DelCost
	}
	for j := 0; j < nj; j++ {
		ssd[0][j] = j * p.InsCost
	}

	for i := 1; i < ni; i++ {
		for j := 1; j < nj; j++ {
			ssd[i][j] = minInt3(
				ssd[i-1][j]+p.DelCost,              // (i-1, j) - Delete
				ssd[i][j-1]+p.InsCost,              // (i, j-1) - Insert
				ssd[i-1][j-1]+calcSubCost(p, i, j), // (i-1, j-1) - Substitution
			)
		}
	}

	return ssd
}

// It is main function
func Distance(p *Params) int {
	if p.LenA < p.LenB {
		return distanceLenA(p)
	}
	return distanceLenB(p)
}

func distanceLenA(p *Params) int {
	var (
		ni = p.LenA + 1
		nj = p.LenB + 1
	)
	vis := make([]int, ni)
	for i := 0; i < ni; i++ {
		vis[i] = i * p.DelCost
	}
	for j := 1; j < nj; j++ {
		vi := vis[0]
		vis[0] = j * p.InsCost
		for i := 1; i < ni; i++ {
			temp := vis[i]
			vis[i] = minInt3(
				vis[i-1]+p.DelCost,      // (i-1, j) - Delete
				vis[i]+p.InsCost,        // (i, j-1) - Insert
				vi+calcSubCost(p, i, j), // (i-1, j-1) - Substitution
			)
			vi = temp
		}
	}
	return vis[ni-1]
}

func distanceLenB(p *Params) int {
	var (
		ni = p.LenA + 1
		nj = p.LenB + 1
	)
	vjs := make([]int, nj)
	for j := 0; j < nj; j++ {
		vjs[j] = j * p.InsCost
	}
	for i := 1; i < ni; i++ {
		vj := vjs[0]
		vjs[0] = i * p.DelCost
		for j := 1; j < nj; j++ {
			temp := vjs[j]
			vjs[j] = minInt3(
				vjs[j]+p.DelCost,        // (i-1, j) - Delete
				vjs[j-1]+p.InsCost,      // (i, j-1) - Insert
				vj+calcSubCost(p, i, j), // (i-1, j-1) - Substitution
			)
			vj = temp
		}
	}
	return vjs[nj-1]
}

func calcSubCost(p *Params, i, j int) int {
	if p.Match(i-1, j-1) {
		return 0
	}
	return p.SubCost
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
