package chow

import "testing"

func TestDummy(t *testing.T) {
	total := 5 + 5

	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
