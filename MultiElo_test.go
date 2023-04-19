package MultiElo

import (
	"reflect"
	"testing"
)

func Test_MultiElo_Can_Calculate_Basic_Two_Player_Liner(t *testing.T) {
	elo := Elo{K: 32, D: 400, Base: 1, Log: 10, Scoring: nil}
	var result = elo.CalculateRating([]float32{1000, 1000}, nil)
	var want = []float32{1016, 984}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Expected to get %v but we got %v instead", want, result)
	}
}

func Test_MultiElo_Can_Calculate_Basic_Two_Player_Exponential(t *testing.T) {
	elo := Elo{K: 32, D: 400, Base: 1.5, Log: 10, Scoring: nil}
	var result = elo.CalculateRating([]float32{1200, 1000}, nil)
	var want = []float32{1207.6881, 992.3119}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Expected to get %v but we got %v instead", want, result)
	}
}

func Test_MultiElo_CalculateActualScores_Works(t *testing.T) {
	elo := Elo{K: 32, D: 400, Base: 1, Log: 10, Scoring: nil}
	elo.initalise()

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
	elo := Elo{K: 32, D: 400, Base: 1, Log: 10, Scoring: nil}
	elo.initalise()

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

func Test_MultiElo_CalculateActualScores_Works_With_Exponential_Scoring(t *testing.T) {
	elo := Elo{K: 32, D: 400, Base: 1.5, Log: 10, Scoring: nil}
	elo.initalise()

	positions := []int{2, 3, 1, 2}

	var scores = elo.CalculateActualScores(4, positions)
	if len(scores) != 4 {
		t.Errorf("Expected to get 4 scores but we got %v", len(scores))
	}

	want := []float32{0.21212121212121213, 0, 0.5757575757575758, 0.21212121212121213}
	if !reflect.DeepEqual(scores, want) {
		t.Errorf("Expected to recieve %v but we got %v instead", want, scores)
	}
}

func Test_MultiElo_CalculateExpectedScores_Works(t *testing.T) {
	elo := Elo{K: 32, D: 400, Base: 1, Log: 10, Scoring: nil}
	elo.initalise()

	ratings := []float32{1200, 1000, 900}
	var scores = elo.CalculateExpectedScores(ratings)

	want := []float32{0.5362558, 0.29343936, 0.17030485}
	if !reflect.DeepEqual(scores, want) {
		t.Errorf("\nExpected to get %v but we got %v instead", want, scores)
	}
}
