package go-multielo

import (
	"math"
	"sort"

	u "github.com/rjNemo/underscore"
)

type Elo struct {
	k, d, base, log float32
	scoring         func(int) []float32
}

func (s *Elo) initalise() {
	if s.scoring == nil {
		s.scoring = Create(s.base)
	}
}

func (s *Elo) CalculateRating(ratings []float32, positions []int) []float32 {
	s.initalise()
	var n = len(ratings)

	// if we din't receive finishing positions for our rating then its already in order 1..n
	if positions == nil {
		positions = Range(1, n)
	}

	var expectedScores = s.CalculateExpectedScores(ratings)
	var actualScores = s.CalculateActualScores(n, positions)
	var scaleFactor = s.k * (float32(n) - 1)

	var adjustments = u.Map(
		Zip(actualScores, expectedScores),
		func(t Tuple[float32, float32]) float32 {
			return scaleFactor * (t.Left - t.Right)
		})

	return u.Map(
		Zip(ratings, adjustments),
		func(t Tuple[float32, float32]) float32 {
			return t.Left + t.Right
		})
}

/*
return ratings
	.Zip(adjustments)
	.Select(x => x.First + x.Second);
*/

func (s *Elo) CalculateActualScores(n int, positions []int) []float32 {
	order := make([]int, n)
	copy(order, positions)
	sort.Ints(order)

	// we need to detrmin if their where any ties, if so we will need to sum
	// the tied scores togther and then distribte them out evenly to the tied players
	// i.e. positions [2,2,1,3] for scores [0.166667, 0.333333, 0.5, 0.0]
	//                            would be [0.25, 0.25, 0.5, 0.0]

	scores := s.scoring(n)
	var joined = Zip(order, scores)
	var s1 = u.GroupBy(joined, func(x Tuple[int, float32]) int {
		return x.Left
	})

	var s2 = make([]Tuple[int, float32], 0)
	for k, v := range s1 {
		var ties = Sum(v, func(x Tuple[int, float32]) float32 { return x.Right })
		s2 = append(s2, Tuple[int, float32]{Left: k, Right: ties / float32(len(v))})
	}

	return JoinProject(positions, s2,
		func(x int) int { return x },
		func(x Tuple[int, float32]) int { return x.Left },
		func(x Tuple[int, []Tuple[int, float32]]) float32 { return x.Right[0].Right })
}

func (s *Elo) CalculateExpectedScores(ratings []float32) []float32 {
	var diffMX = u.Map(ratings, func(x float32) []float32 {
		return u.Map(ratings, func(y float32) float32 {
			return y - x
		})
	})

	var logMX = u.Map(diffMX, func(diffs []float32) []float32 {
		return u.Map(diffs, func(x float32) float32 {
			return float32(1 / (1 + math.Pow(float64(s.log), float64(x)/float64(s.d))))
		})
	})

	logMX = DiagonalFill(logMX, func(_ float32) float32 { return 0 })
	var expected = u.Map(logMX, func(x []float32) float32 { return u.Sum(x) })

	var n = float32(len(expected))
	var denom = n * (n - 1) / 2

	return u.Map(expected, func(x float32) float32 { return x / denom })
}
