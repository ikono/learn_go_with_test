package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5
	if sum != expected {
		t.Errorf("sum is %d, expected is %d", sum, expected)
	}
}
