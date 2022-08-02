package MultiElo

type Settings struct {
	k, d, base, log float32
	scoring         func(int) []float32
}

func (s *Settings) Initalise(k, d, base, log float32, customScoring func(int) []float32) {
	s.k = k
	s.d = d
	s.base = base
	s.log = log

	if customScoring == nil {
		s.scoring = Create(s.base)
	} else {
		s.scoring = customScoring
	}
}

func (s *Settings) CalculateRating(ratings []float32, order []int) []float32 {
	var scores = s.scoring(len(ratings))

	//todo: replace with actual results
	return scores
}
