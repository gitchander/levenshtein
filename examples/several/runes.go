package main

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
