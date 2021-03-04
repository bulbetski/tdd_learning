package integers

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(1,2)
	expected := 3

	if sum != expected {
		t.Errorf("expected %d got %d", expected, sum)
	}
}

