package MultiElo

func Range(start int, end int) []int {
	result := make([]int, end)
	for i := start; i <= end; i++ {
		result[i] = i
	}

	return result
}
