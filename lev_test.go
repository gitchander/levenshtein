package levenshtein

import (
	"math/rand"
	"testing"

	"github.com/gitchander/levenshtein/utils/random"
)

func TestRandomSamples(t *testing.T) {

	cs := Costs{
		DelCost: 1,
		InsCost: 1,
		SubCost: 1,
	}
	corpus := []rune("abcdefghijklmnopqrstuvwxyz")

	r := random.NewRandNow()
	//r := random.NewRandSeed(0)

	for i := 0; i < 1000; i++ {
		var (
			a = random.RunesByCorpus(r, corpus, r.Intn(50))
			b = cloneRunes(a)
		)
		var cost int
		if len(a) > 0 {
			jn := r.Intn(len(a))
			for j := 0; j < jn; j++ {
				n := len(b)
				if n == 0 {
					nr := random.RuneByCorpus(r, corpus)
					b = insertRune(b, r.Intn(n+1), nr)
					cost += cs.InsCost
				} else {
					switch k := r.Intn(3); k {
					case 0:
						{
							b = deleteRune(b, r.Intn(n))
							cost += cs.DelCost
						}
					case 1:
						{
							nr := random.RuneByCorpus(r, corpus)
							b = insertRune(b, r.Intn(n+1), nr)
							cost += cs.InsCost
						}
					case 2:
						{
							nr := random.RuneByCorpus(r, corpus)
							b = substituteRune(b, r.Intn(n), nr)
							cost += cs.SubCost
						}
					}
				}
			}
		}
		res := Runes(a, b)
		// if res != cost {
		// 	t.Logf("(%d != %d), ['%s', '%s']", res, cost, string(a), string(b))
		// }
		if res > cost {
			t.Fatalf("(%d > %d), ['%s', '%s']", res, cost, string(a), string(b))
		}
	}
}

func cloneRunes(a []rune) []rune {
	b := make([]rune, len(a))
	copy(b, a)
	return b
}

func deleteRune(rs []rune, i int) []rune {
	n := len(rs)
	if (i < 0) || (n <= i) {
		panic("invalid index")
	}
	copy(rs[i:], rs[i+1:])
	return rs[:n-1]
}

func insertRune(as []rune, i int, r rune) []rune {

	n := len(as)
	if (i < 0) || (n < i) {
		panic("invalid index")
	}

	bs := make([]rune, n+1)

	copy(bs, as[:i])
	bs[i] = r
	copy(bs[i+1:], as[i:])

	return bs
}

func substituteRune(rs []rune, i int, r rune) []rune {
	n := len(rs)
	if (i < 0) || (n <= i) {
		panic("invalid index")
	}
	rs[i] = r
	return rs
}

func TestLens(t *testing.T) {

	r := random.NewRandNow()

	cs := Costs{
		DelCost: 1,
		InsCost: 1,
		SubCost: 1,
	}
	corpus := []rune("abcdefghijklmnopqrstuvwxyz")

	m := &runesMutator{
		r:      r,
		cs:     cs,
		corpus: corpus,
	}

	for i := 0; i < 1000; i++ {
		var (
			a = random.RunesByCorpus(r, corpus, r.Intn(50))
			b = cloneRunes(a)
		)

		jn := r.Intn(30)
		for j := 0; j < jn; j++ {
			b = m.Mutate(b)
		}

		rs := RuneSlices{a, b}

		var (
			d0 = distanceByLen0(rs, cs)
			d1 = distanceByLen1(rs, cs)
		)

		if d0 != d1 {
			t.Fatalf("distances: (%d != %d), values: (%q, %q)",
				d0, d1, string(a), string(b))
		}
	}
}

type runesMutator struct {
	r      *rand.Rand
	cs     Costs
	corpus []rune

	cost int
}

func (m *runesMutator) Reset() {
	m.cost = 0
}

func (m *runesMutator) Cost() int {
	return m.cost
}

func (m *runesMutator) randRune() rune {
	return random.RuneByCorpus(m.r, m.corpus)
}

func (m *runesMutator) Mutate(rs []rune) []rune {

	n := len(rs)
	if n == 0 {
		nr := m.randRune()
		rs = insertRune(rs, m.r.Intn(n+1), nr)
		m.cost += m.cs.InsCost
		return rs
	}

	switch k := m.r.Intn(3); k {
	case 0:
		{
			rs = deleteRune(rs, m.r.Intn(n))
			m.cost += m.cs.DelCost
		}
	case 1:
		{
			nr := m.randRune()
			rs = insertRune(rs, m.r.Intn(n+1), nr)
			m.cost += m.cs.InsCost
		}
	case 2:
		{
			nr := m.randRune()
			rs = substituteRune(rs, m.r.Intn(n), nr)
			m.cost += m.cs.SubCost
		}
	}

	return rs
}
