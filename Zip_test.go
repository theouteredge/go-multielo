package MultiElo

import "testing"

func Test_Zip_Can_Zip_Two_Equal_Sized_Slices(t *testing.T) {
	left := make([]string, 0)
	left = append(left, "Left 1", "Left 2", "Left 3")

	right := make([]int, 0)
	right = append(right, 1, 2, 4)

	var zipped = Zip(left, right)
	if len(zipped) != 3 {
		t.Errorf("Expected the result of Zip(left, right) to have a length of 3 but got %v", len(zipped))
	}

	for i, z := range zipped {
		if z.Left != left[i] {
			t.Errorf("Expected 'Left %v' at position %v", i, i)
		}
		if z.Right != right[i] {
			t.Errorf("Expected '%v' at position %v", i, i)
		}
	}
}

func Test_Zip_Can_Zip_Two_Different_Sized_Slices_Left_Larger(t *testing.T) {
	left := make([]string, 0)
	left = append(left, "Left 1", "Left 2", "Left 3", "Left 4")

	right := make([]int, 0)
	right = append(right, 1, 2, 3)

	var zipped = Zip(left, right)
	if len(zipped) != 3 {
		t.Errorf("Expected the result of Zip(left, right) to have a length of 3 but got %v", len(zipped))
	}

	for i, z := range zipped {
		if z.Left != left[i] {
			t.Errorf("Expected 'Left %v' at position %v", i, i)
		}
		if z.Right != right[i] {
			t.Errorf("Expected '%v' at position %v", i, i)
		}
	}
}

func Test_Zip_Can_Zip_Two_Different_Sized_Slices_Right_Larger(t *testing.T) {
	left := make([]string, 0)
	left = append(left, "Left 1", "Left 2", "Left 3")

	right := make([]int, 0)
	right = append(right, 1, 2, 3, 4)

	var zipped = Zip(left, right)
	if len(zipped) != 3 {
		t.Errorf("Expected the result of Zip(left, right) to have a length of 3 but got %v", len(zipped))
	}

	for i, z := range zipped {
		if z.Left != left[i] {
			t.Errorf("Expected 'Left %v' at position %v", i, i)
		}
		if z.Right != right[i] {
			t.Errorf("Expected '%v' at position %v", i, i)
		}
	}
}
