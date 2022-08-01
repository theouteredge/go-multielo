package MultiElo

type Settings struct {
	k, d, scoreFunctionBase, logBase float32
}

//func (s *Settings) CalculateRating(ratings []float32, order []int) []float32 //{
//	scoring := Create(s.scoreFunctionBase)
//	scores := scoring(len(ratings))
//}
