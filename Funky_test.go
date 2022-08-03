package MultiElo

import (
	"reflect"
	"testing"
)

func Test_Zip_Can_Zip_Two_Equal_Sized_Slices(t *testing.T) {
	left := []string{"Left 1", "Left 2", "Left 3"}
	right := []int{1, 2, 3}

	var zipped = Zip(left, right)

	want := []Tuple[string, int]{
		{Left: "Left 1", Right: 1},
		{Left: "Left 2", Right: 2},
		{Left: "Left 3", Right: 3},
	}

	if !reflect.DeepEqual(zipped, want) {
		t.Errorf("Expected the result to be %v but we got %v", want, zipped)
	}
}

func Test_Zip_Can_Zip_Two_Different_Sized_Slices_Left_Larger(t *testing.T) {
	left := []string{"Left 1", "Left 2", "Left 3", "Left 4"}
	right := []int{1, 2, 3}

	var zipped = Zip(left, right)
	if len(zipped) != 3 {
		t.Errorf("Expected the result of Zip(left, right) to have a length of 3 but got %v", len(zipped))
	}

	want := []Tuple[string, int]{
		{Left: "Left 1", Right: 1},
		{Left: "Left 2", Right: 2},
		{Left: "Left 3", Right: 3},
	}

	if !reflect.DeepEqual(zipped, want) {
		t.Errorf("Expected the result to be %v but we got %v", want, zipped)
	}
}

func Test_Zip_Can_Zip_Two_Different_Sized_Slices_Right_Larger(t *testing.T) {
	left := []string{"Left 1", "Left 2", "Left 3"}
	right := []int{1, 2, 3, 4}

	var zipped = Zip(left, right)
	if len(zipped) != 3 {
		t.Errorf("Expected the result of Zip(left, right) to have a length of 3 but got %v", len(zipped))
	}

	want := []Tuple[string, int]{
		{Left: "Left 1", Right: 1},
		{Left: "Left 2", Right: 2},
		{Left: "Left 3", Right: 3},
	}

	if !reflect.DeepEqual(zipped, want) {
		t.Errorf("Expected the result to be %v but we got %v", want, zipped)
	}
}

func Test_Join_Can_Join_Two_Slices_Together(t *testing.T) {
	one := Tuple[int, string]{Left: 1, Right: "One"}
	two := Tuple[int, string]{Left: 2, Right: "Two"}
	three := Tuple[int, string]{Left: 3, Right: "Three"}

	var left = []Tuple[int, string]{one, two, three}
	var right = []Tuple[int, string]{one, three, two, three, two, three}

	selector := func(x Tuple[int, string]) int { return x.Left }

	var joined = Join(left, right, selector, selector)
	var want = []Tuple[Tuple[int, string], []Tuple[int, string]]{
		{Left: one, Right: []Tuple[int, string]{one}},
		{Left: two, Right: []Tuple[int, string]{two, two}},
		{Left: three, Right: []Tuple[int, string]{three, three, three}},
	}

	if !reflect.DeepEqual(joined, want) {
		t.Errorf("Expected to get %v but we got %v instead", want, joined)
	}
}
