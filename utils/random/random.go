package random

import (
	"math/rand"
)

func RuneByCorpus(r *rand.Rand, corpus []rune) rune {
	return corpus[r.Intn(len(corpus))]
}

func RunesByCorpus(r *rand.Rand, corpus []rune, n int) []rune {
	rs := make([]rune, n)
	for i := range rs {
		rs[i] = RuneByCorpus(r, corpus)
	}
	return rs
}
