package MultiElo

import (
	u "github.com/rjNemo/underscore"
	"golang.org/x/exp/constraints"
)

type Tuple[L, R any] struct {
	Left  L
	Right R
}

func Range(start int, end int) []int {
	result := make([]int, end)
	for i := start; i <= end; i++ {
		result[i] = i
	}

	return result
}

func Zip[L any, R any](left []L, right []R) []Tuple[L, R] {
	var shortest = 0
	if len(left) < len(right) {
		shortest = len(left)
	} else {
		shortest = len(right)
	}

	results := make([]Tuple[L, R], shortest)
	for i := 0; i < shortest; i++ {
		results[i] = Tuple[L, R]{Left: left[i], Right: right[i]}
	}

	return results
}

func Join[T, P any, S comparable](
	left []T,
	right []P,
	leftSelector func(T) S,
	rightSelector func(P) S) []Tuple[T, []P] {

	var results = make([]Tuple[T, []P], 0, len(left))
	for _, l := range left {
		var matches = u.Filter(right, func(r P) bool { return leftSelector(l) == rightSelector(r) })
		var tuple = Tuple[T, []P]{Left: l, Right: matches}
		results = append(results, tuple)
	}

	return results
}

func JoinProject[T, P, R any, S comparable](
	left []T,
	right []P,
	leftSelector func(T) S,
	rightSelector func(P) S,
	projection func(Tuple[T, []P]) R) (results []R) {

	for _, x := range Join(left, right, leftSelector, rightSelector) {
		results = append(results, projection(x))
	}

	return results
}

func Sum[T any, R constraints.Ordered](list []T, selector func(T) R) (sum R) {
	for _, v := range list {
		sum += selector(v)
	}

	return sum
}

func DiagonalFill[T constraints.Ordered](list [][]T, f func(T) T) [][]T {
	for i := 0; i < len(list); i++ {
		list[i][i] = f(list[i][i])
	}

	return list
}
