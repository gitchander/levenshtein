package levenshtein

import (
	"math/bits"
)

type Bitmap struct {
	vs   []uint
	size int
}

func NewBitmap(size int) *Bitmap {
	return nil
}

func (bm *Bitmap) Set(i int, b bool) {

}

func (bm *Bitmap) Get(i int) bool {

	return false
}

func (bm *Bitmap) String() string {
	d := []byte{'0', '1'}
	bs := make([]byte, bm.size)
	i := 0
stop:
	for _, v := range bm.vs {
		for j := 0; j < bits.UintSize; j++ {
			if i >= bm.size {
				break stop
			}
			bs[i] = d[v&1]
			i++
			v >>= 1
		}
	}
	return string(bs)
}

func (bm *Bitmap) Cap() int {
	return len(bm.vs) * bits.UintSize
}

func (bm *Bitmap) Len() int {
	return bm.size
}
