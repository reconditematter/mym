package mym

import (
	"testing"
)

// TestMT19937 checks if the 10000th consecutive pseudo-random number
// generated by MT19937 is 9981545732273789042.
// Reference: `Standard for Programming Language C++` (§29.6.5).
func TestMT19937(t *testing.T) {
	rng := MT19937()
	const V10000 = uint64(9981545732273789042)
	var v uint64
	for m := 0; m < 10000; m++ {
		v = rng.Uint64()
	}
	if v != V10000 {
		t.Fatalf("mt19937(10000) != %v", V10000)
	}
}
