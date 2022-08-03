package MultiElo

import (
	"sort"

	u "github.com/rjNemo/underscore"
)

type Elo struct {
	k, d, base, log float32
	scoring         func(int) []float32
}

func (s *Elo) Initalise() {
	if s.scoring == nil {
		s.scoring = Create(s.base)
	}
}

func (s *Elo) CalculateRating(ratings []float32, order []int) []float32 {
	var scores = s.scoring(len(ratings))

	//todo: replace with actual results
	return scores
}

func (s *Elo) CalculateActualScores(n int, positions []int) []float32 {
	// we need to preserve the finishing positions, so we can return the
	// scores back in the same order as we got them
	var standings = make([]int, len(positions))
	if positions == nil {
		standings = Range(1, n)
	} else {
		copy(standings, positions)
		sort.Ints(standings)
	}

	scores := s.scoring(n)

	// we need to detrmin if their where any ties, if so we will need to sum
	// the tied scores togther and then distribte them out evenly to the tied
	// players

	var joined = Zip(standings, scores)
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
