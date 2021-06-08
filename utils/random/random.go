package random

import (
	"math/rand"
	"time"
)

func NewRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func NewRandTime(t time.Time) *rand.Rand {
	return NewRandSeed(t.UnixNano())
}

func NewRandNow() *rand.Rand {
	return NewRandTime(time.Now())
}

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
