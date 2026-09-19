package levenshtein

import (
	"reflect"
	"testing"
)

type testSample struct {
	a, b  string
	costs Costs

	distance int
	matrix   [][]int
}

func TestMatrixSamples(t *testing.T) {
	samples := []testSample{
		{
			a: "sitting",
			b: "kitten",
			costs: Costs{
				DelCost: 1,
				InsCost: 1,
				SubCost: 1,
			},
			matrix: [][]int{
				{0, 1, 2, 3, 4, 5, 6},
				{1, 1, 2, 3, 4, 5, 6},
				{2, 2, 1, 2, 3, 4, 5},
				{3, 3, 2, 1, 2, 3, 4},
				{4, 4, 3, 2, 1, 2, 3},
				{5, 5, 4, 3, 2, 2, 3},
				{6, 6, 5, 4, 3, 3, 2},
				{7, 7, 6, 5, 4, 4, 3},
			},
			distance: 3,
		},
		{
			a: "Sunday",
			b: "Saturday",
			costs: Costs{
				DelCost: 1,
				InsCost: 1,
				SubCost: 1,
			},
			matrix: [][]int{
				{0, 1, 2, 3, 4, 5, 6, 7, 8},
				{1, 0, 1, 2, 3, 4, 5, 6, 7},
				{2, 1, 1, 2, 2, 3, 4, 5, 6},
				{3, 2, 2, 2, 3, 3, 4, 5, 6},
				{4, 3, 3, 3, 3, 4, 3, 4, 5},
				{5, 4, 3, 4, 4, 4, 4, 3, 4},
				{6, 5, 4, 4, 5, 5, 5, 4, 3},
			},
			distance: 3,
		},
		{
			a: "hello",
			b: "world",
			costs: Costs{
				DelCost: 1,
				InsCost: 1,
				SubCost: 1,
			},
			matrix: [][]int{
				{0, 1, 2, 3, 4, 5},
				{1, 1, 2, 3, 4, 5},
				{2, 2, 2, 3, 4, 5},
				{3, 3, 3, 3, 3, 4},
				{4, 4, 4, 4, 3, 4},
				{5, 5, 4, 5, 4, 4},
			},
			distance: 4,
		},
		{
			a: "AACGCA",
			b: "GAGCTA",
			costs: Costs{
				DelCost: 1,
				InsCost: 1,
				SubCost: 2,
			},
			matrix: [][]int{
				{0, 1, 2, 3, 4, 5, 6},
				{1, 2, 1, 2, 3, 4, 5},
				{2, 3, 2, 3, 4, 5, 4},
				{3, 4, 3, 4, 3, 4, 5},
				{4, 3, 4, 3, 4, 5, 6},
				{5, 4, 5, 4, 3, 4, 5},
				{6, 5, 4, 5, 4, 5, 4},
			},
			distance: 4,
		},
		{
			a: "exponential",
			b: "polynomial",
			costs: Costs{
				DelCost: 1,
				InsCost: 1,
				SubCost: 1,
			},
			matrix: [][]int{
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
			distance: 6,
		},
	}
	for i, sample := range samples {

		v := RunePair{
			[]rune(sample.a),
			[]rune(sample.b),
		}

		var (
			haveDistance = DistanceCosts(v, sample.costs)
			wantDistance = sample.distance
		)
		if haveDistance != wantDistance {
			t.Fatalf("sample [%d]: invalid distance: have %d, want %d", i,
				haveDistance, wantDistance)
		}

		var (
			haveMatrix = MakeMatrix(v, sample.costs)
			wantMatrix = sample.matrix
		)
		if !(reflect.DeepEqual(haveMatrix, wantMatrix)) {
			t.Fatalf("sample [%d]: matrixes are not equal", i)
		}
	}
}
