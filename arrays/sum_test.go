package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	sum := Sum(numbers)
	expected := 15

	if sum != expected {
		t.Errorf("expected %d got %d, %v", expected, sum, numbers)
	}
}
