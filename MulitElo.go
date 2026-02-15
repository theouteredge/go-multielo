package MultiElo

import (
	"math"
	"sort"

	u "github.com/rjNemo/underscore"
	"golang.org/x/exp/constraints"
)

var k, d, base, log float32 = 32, 400, 10, 10
var scoring func(int) []float32

type Rating[T any] struct {
	Position int
	Rating   float32
	Adjusted float32
	Value    T
}

func Initalise(_k, _d, _base, _log float32) {
	k, d, base, log = _k, _d, _base, _log
	if scoring == nil {
		scoring = Create(base)
	}
}

func CalculateRating[T any](ratings []*Rating[T]) {
	var n = len(ratings)

	positions := u.Map(ratings, func(x *Rating[T]) int { return x.Position })
	expected := CalculateExpectedScores(ratings)
	scores := CalculateActualScores(n, positions)
	scaleFactor := k * (float32(n) - 1)

	adjustments := u.Map(u.Zip(scores, expected),
		func(t u.Tuple[float32, float32]) float32 {
			return scaleFactor * (t.Left - t.Right)
		})

	zipped := u.Zip(ratings, adjustments)
	for _, t := range zipped {
		t.Left.Adjusted = t.Right
		t.Left.Rating = t.Left.Rating + t.Left.Adjusted
	}
}

func CalculateActualScores(n int, positions []int) []float32 {
	order := make([]int, n)
	copy(order, positions)
	sort.Ints(order)

	// we need to detrmin if their where any ties, if so we will need to sum
	// the tied scores togther and then distribte them out evenly to the tied players
	// i.e. positions [2,2,1,3] for scores [0.166667, 0.333333, 0.5, 0.0]
	//                            would be [0.25, 0.25, 0.5, 0.0]

	scores := scoring(n)
	joined := u.Zip(order, scores)
	s1 := u.GroupBy(joined, func(x u.Tuple[int, float32]) int {
		return x.Left
	})

	var s2 = make([]u.Tuple[int, float32], 0)
	for k, v := range s1 {
		ties := u.SumMap(v, func(x u.Tuple[int, float32]) float32 { return x.Right })
		s2 = append(s2, u.Tuple[int, float32]{Left: k, Right: ties / float32(len(v))})
	}

	return u.JoinProject(positions, s2,
		func(x int) int { return x },
		func(x u.Tuple[int, float32]) int { return x.Left },
		func(x u.Tuple[int, []u.Tuple[int, float32]]) float32 { return x.Right[0].Right })
}

func CalculateExpectedScores[T any](ratings []*Rating[T]) []float32 {
	r := u.Map(ratings, func(x *Rating[T]) float32 { return x.Rating })

	diffMX := u.Map(r, func(x float32) []float32 {
		return u.Map(r, func(y float32) float32 {
			return y - x
		})
	})

	logMX := u.Map(diffMX, func(diffs []float32) []float32 {
		return u.Map(diffs, func(x float32) float32 {
			return float32(1 / (1 + math.Pow(float64(log), float64(x)/float64(d))))
		})
	})

	logMX = diagonalFill(logMX, func(_ float32) float32 { return 0 })
	expected := u.Map(logMX, func(x []float32) float32 { return u.Sum(x) })

	n := float32(len(expected))
	denom := n * (n - 1) / 2

	normalized := u.Map(expected, func(x float32) float32 { return x / denom })
	return normalized
}

func diagonalFill[T constraints.Ordered](list [][]T, f func(T) T) [][]T {
	for i := 0; i < len(list); i++ {
		list[i][i] = f(list[i][i])
	}

	return list
}
