package MultiElo

import (
	"reflect"
	"testing"
)

func Test_MultiElo_CalculateActualScores_Works(t *testing.T) {
	elo := Elo{k: 32, d: 400, base: 1, log: 10, scoring: nil}
	elo.Initalise()

	positions := []int{2, 4, 1, 3}

	var scores = elo.CalculateActualScores(4, positions)
	if len(scores) != 4 {
		t.Errorf("Expected to get 4 scores but we got %v", len(scores))
	}

	want := []float32{0.33333334, 0.0, 0.5, 0.16666667}
	if !reflect.DeepEqual(scores, want) {
		t.Errorf("Expected to recieve %v but we got %v instead", want, scores)
	}
}

func Test_MultiElo_CalculateActualScores_Works_With_Ties(t *testing.T) {
	elo := Elo{k: 32, d: 400, base: 1, log: 10, scoring: nil}
	elo.Initalise()

	positions := []int{2, 3, 1, 2}

	var scores = elo.CalculateActualScores(4, positions)
	if len(scores) != 4 {
		t.Errorf("Expected to get 4 scores but we got %v", len(scores))
	}

	want := []float32{0.25, 0.0, 0.5, 0.25}
	if !reflect.DeepEqual(scores, want) {
		t.Errorf("Expected to recieve %v but we got %v instead", want, scores)
	}
}

func Test_MultiElo_CalculateExpectedScores_Works(t *testing.T) {
	elo := Elo{k: 32, d: 400, base: 1, log: 10, scoring: nil}
	elo.Initalise()

	ratings := []float32{1200, 1000, 900}
	var scores = elo.CalculateExpectedScores(ratings)

	want := []float32{0.5362558, 0.29343936, 0.17030485}
	if !reflect.DeepEqual(scores, want) {
		t.Errorf("\nExpected to get %v but we got %v instead", want, scores)
	}
}
