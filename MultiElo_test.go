package MultiElo

import (
	"reflect"
	"testing"
)

func Test_MultiElo_Can_Calculate_Basic_Two_Player_Liner(t *testing.T) {
	Initalise(32, 400, 10, 10)

	ratings := []*Rating[string]{
		{Rating: 1000, Position: 1, Value: "Player 1"},
		{Rating: 1000, Position: 2, Value: "Player 2"},
	}

	CalculateRating(ratings)
	want := []float32{1016, 984}
	actual := []float32{ratings[0].Rating, ratings[1].Rating}

	if !reflect.DeepEqual(actual, want) {
		t.Errorf("Expected to get %v but we got %v instead", want, actual)
	}
}

func Test_MultiElo_Can_Calculate_Basic_Two_Player_Exponential(t *testing.T) {
	Initalise(32, 400, 1.5, 10)

	ratings := []*Rating[string]{
		{Rating: 1200, Position: 1, Value: "Player 1"},
		{Rating: 1000, Position: 2, Value: "Player 2"},
	}

	CalculateRating(ratings)
	want := []float32{1207.6881, 992.3119}
	actual := []float32{ratings[0].Rating, ratings[1].Rating}

	if !reflect.DeepEqual(actual, want) {
		t.Errorf("Expected to get %v but we got %v instead", want, actual)
	}
}

func Test_MultiElo_CalculateActualScores_Works(t *testing.T) {
	Initalise(32, 400, 1.5, 10)

	positions := []int{2, 4, 1, 3}
	scores := CalculateActualScores(4, positions)
	if len(scores) != 4 {
		t.Errorf("Expected to get 4 scores but we got %v", len(scores))
	}

	want := []float32{0.33333334, 0.0, 0.5, 0.16666667}
	if !reflect.DeepEqual(scores, want) {
		t.Errorf("Expected to recieve %v but we got %v instead", want, scores)
	}
}

func Test_MultiElo_CalculateActualScores_Works_With_Ties(t *testing.T) {
	Initalise(32, 400, 1.5, 10)

	positions := []int{2, 3, 1, 2}
	scores := CalculateActualScores(4, positions)
	if len(scores) != 4 {
		t.Errorf("Expected to get 4 scores but we got %v", len(scores))
	}

	want := []float32{0.25, 0.0, 0.5, 0.25}
	if !reflect.DeepEqual(scores, want) {
		t.Errorf("Expected to recieve %v but we got %v instead", want, scores)
	}
}

func Test_MultiElo_CalculateActualScores_Works_With_Exponential_Scoring(t *testing.T) {
	Initalise(32, 400, 1.5, 10)
	positions := []int{2, 3, 1, 2}

	scores := CalculateActualScores(4, positions)
	if len(scores) != 4 {
		t.Errorf("Expected to get 4 scores but we got %v", len(scores))
	}

	want := []float32{0.21212121212121213, 0, 0.5757575757575758, 0.21212121212121213}
	if !reflect.DeepEqual(scores, want) {
		t.Errorf("Expected to recieve %v but we got %v instead", want, scores)
	}
}

func Test_MultiElo_CalculateExpectedScores_Works(t *testing.T) {
	Initalise(32, 400, 1.5, 10)

	ratings := []*Rating[string]{
		{Rating: 1200, Value: "Player 1"},
		{Rating: 1000, Value: "Player 2"},
		{Rating: 900, Value: "Player 3"},
	}
	scores := CalculateExpectedScores(ratings)

	want := []float32{0.5362558, 0.29343936, 0.17030485}
	if !reflect.DeepEqual(scores, want) {
		t.Errorf("\nExpected to get %v but we got %v instead", want, scores)
	}
}
