package MultiElo

import (
	"fmt"
	"math"

	us "github.com/rjNemo/underscore"
)

func Create(base float32) func(int) []float32 {
	if base < 1 {
		panic(fmt.Sprintf("The value of base must be 1 or greater, but we recieved a value of %v", 2))
	}

	if base == 1 {
		return func(n int) []float32 {
			return Liner(n)
		}
	}

	return func(n int) []float32 {
		return Exponential(n, base)
	}
}

func Liner(n int) (scores []float32) {
	var nf = float32(n)
	for p := 1; p <= n; p++ {
		pf := float32(p)
		scores = append(scores, (nf-pf)/(nf*(nf-1)/2))
	}

	return scores
}

func Exponential(n int, base float32) []float32 {
	var nf = float64(n)
	var output = make([]float32, 0)

	for p := 1; p <= n; p++ {
		pf := float64(p)
		output = append(output, float32(math.Pow(float64(base), nf-pf)-1))
	}

	length := float32(len(output))
	return us.Map(output, func(x float32) float32 {
		if x == 0 {
			return x
		}

		return x / length
	})
}
