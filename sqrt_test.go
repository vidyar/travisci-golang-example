package newmath

import "testing"

func TestSqrt(t *testing.T) {
	const in, out = 6, 2
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}
