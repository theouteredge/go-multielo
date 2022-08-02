package MultiElo

type Zipped[L, R any] struct {
	Left  L
	Right R
}

func Zip[L any, R any](left []L, right []R) []Zipped[L, R] {
	var shortest = 0
	if len(left) < len(right) {
		shortest = len(left)
	} else {
		shortest = len(right)
	}

	results := make([]Zipped[L, R], shortest)
	for i := 0; i < shortest; i++ {
		results[i] = Zipped[L, R]{Left: left[i], Right: right[i]}
	}

	return results
}
