package levenshtein

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

func maxInt2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func digitsNumber(x, base int) int {
	var n int
	for x > 0 {
		x /= base
		n++
	}
	if n == 0 {
		n = 1
	}
	return n
}

func repeatByte(b byte, n int) []byte {
	bs := make([]byte, n)
	for i := range bs {
		bs[i] = b
	}
	return bs
}
