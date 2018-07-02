package levenshtein

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandomSamples(t *testing.T) {
	costs := Costs{
		DelCost: 1,
		InsCost: 1,
		SubCost: 1,
	}
	source := []rune("abcdefghijklmnopqrstuvwxyz")

	r := newRandNow()
	//r := newRandSeed(0)

	for i := 0; i < 1000; i++ {
		var (
			a = randRunes(r, source, r.Intn(50))
			b = copyRunes(a)
		)
		var cost int
		if len(a) > 0 {
			jn := r.Intn(len(a))
			for j := 0; j < jn; j++ {
				n := len(b)
				if n == 0 {
					b = insertRune(b, r.Intn(n+1), source[r.Intn(len(source))])
					cost += costs.InsCost
				} else {
					switch k := r.Intn(3); k {
					case 0:
						b = deleteRune(b, r.Intn(n))
						cost += costs.DelCost
					case 1:
						b = insertRune(b, r.Intn(n+1), source[r.Intn(len(source))])
						cost += costs.InsCost
					case 2:
						b = substituteRune(b, r.Intn(n), source[r.Intn(len(source))])
						cost += costs.SubCost
					}
				}
			}
		}
		res := Runes(a, b)
		//		if res != cost {
		//			t.Logf("(%d != %d), ['%s', '%s']", res, cost, string(a), string(b))
		//		}
		if res > cost {
			t.Fatalf("(%d > %d), ['%s', '%s']", res, cost, string(a), string(b))
		}
	}
}

func newRandNow() *rand.Rand {
	return newRandTime(time.Now())
}

func newRandTime(t time.Time) *rand.Rand {
	return newRandSeed(t.UTC().UnixNano())
}

func newRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func copyRunes(a []rune) []rune {
	b := make([]rune, len(a))
	copy(b, a)
	return b
}

func randRunes(r *rand.Rand, source []rune, n int) []rune {
	rs := make([]rune, n)
	for i := range rs {
		rs[i] = source[r.Intn(len(source))]
	}
	return rs
}

func deleteRune(rs []rune, i int) []rune {
	n := len(rs)
	if (0 > i) || (i >= n) {
		panic("invalid index")
	}
	copy(rs[i:], rs[i+1:])
	return rs[:n-1]
}

func insertRune(rs []rune, i int, r rune) []rune {

	n := len(rs)

	if (0 > i) || (i > n) {
		panic("invalid index")
	}

	rs = append(rs, r)

	if i == n {
		return rs
	}

	copy(rs[i+1:n+1], rs[i:n])
	rs[i] = r

	return rs
}

func substituteRune(rs []rune, i int, r rune) []rune {
	n := len(rs)
	if (0 > i) || (i >= n) {
		panic("invalid index")
	}
	rs[i] = r
	return rs
}
