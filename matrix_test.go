package levenshtein

import (
	"reflect"
	"testing"
)

type matrixSample struct {
	a, b  []rune
	costs Costs

	result [][]int
}

func TestMatrixSamples(t *testing.T) {
	samples := []matrixSample{
		{
			a: []rune("sitting"),
			b: []rune("kitten"),
			costs: Costs{
				InsCost: 1,
				DelCost: 1,
				SubCost: 1,
			},
			result: [][]int{
				{0, 1, 2, 3, 4, 5, 6},
				{1, 1, 2, 3, 4, 5, 6},
				{2, 2, 1, 2, 3, 4, 5},
				{3, 3, 2, 1, 2, 3, 4},
				{4, 4, 3, 2, 1, 2, 3},
				{5, 5, 4, 3, 2, 2, 3},
				{6, 6, 5, 4, 3, 3, 2},
				{7, 7, 6, 5, 4, 4, 3},
			},
		},
		{
			a: []rune("Sunday"),
			b: []rune("Saturday"),
			costs: Costs{
				InsCost: 1,
				DelCost: 1,
				SubCost: 1,
			},
			result: [][]int{
				{0, 1, 2, 3, 4, 5, 6, 7, 8},
				{1, 0, 1, 2, 3, 4, 5, 6, 7},
				{2, 1, 1, 2, 2, 3, 4, 5, 6},
				{3, 2, 2, 2, 3, 3, 4, 5, 6},
				{4, 3, 3, 3, 3, 4, 3, 4, 5},
				{5, 4, 3, 4, 4, 4, 4, 3, 4},
				{6, 5, 4, 4, 5, 5, 5, 4, 3},
			},
		},
		{
			a: []rune("hello"),
			b: []rune("world"),
			costs: Costs{
				InsCost: 1,
				DelCost: 1,
				SubCost: 1,
			},
			result: [][]int{
				{0, 1, 2, 3, 4, 5},
				{1, 1, 2, 3, 4, 5},
				{2, 2, 2, 3, 4, 5},
				{3, 3, 3, 3, 3, 4},
				{4, 4, 4, 4, 3, 4},
				{5, 5, 4, 5, 4, 4},
			},
		},
		{
			a: []rune("AACGCA"),
			b: []rune("GAGCTA"),
			costs: Costs{
				InsCost: 1,
				DelCost: 1,
				SubCost: 2,
			},
			result: [][]int{
				{0, 1, 2, 3, 4, 5, 6},
				{1, 2, 1, 2, 3, 4, 5},
				{2, 3, 2, 3, 4, 5, 4},
				{3, 4, 3, 4, 3, 4, 5},
				{4, 3, 4, 3, 4, 5, 6},
				{5, 4, 5, 4, 3, 4, 5},
				{6, 5, 4, 5, 4, 5, 4},
			},
		},
		{
			a: []rune("exponential"),
			b: []rune("polynomial"),
			costs: Costs{
				InsCost: 1,
				DelCost: 1,
				SubCost: 1,
			},
			result: [][]int{
				{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{3, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10},
				{4, 3, 2, 3, 4, 5, 5, 6, 7, 8, 9},
				{5, 4, 3, 3, 4, 4, 5, 6, 7, 8, 9},
				{6, 5, 4, 4, 4, 5, 5, 6, 7, 8, 9},
				{7, 6, 5, 5, 5, 4, 5, 6, 7, 8, 9},
				{8, 7, 6, 6, 6, 5, 5, 6, 7, 8, 9},
				{9, 8, 7, 7, 7, 6, 6, 6, 6, 7, 8},
				{10, 9, 8, 8, 8, 7, 7, 7, 7, 6, 7},
				{11, 10, 9, 8, 9, 8, 8, 8, 8, 7, 6},
			},
		},
	}
	for i, sample := range samples {
		var (
			v = RuneSlices{sample.a, sample.b}

			haveMatrix = MakeMatrix(v, sample.costs)
			wantMatrix = sample.result
		)
		if !(reflect.DeepEqual(haveMatrix, wantMatrix)) {
			t.Fatalf("sample [%d]: matrixes are not equal", i)
		}
	}
}
