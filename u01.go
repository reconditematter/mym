package mym

// U01 -- returns a uniform pseudo-random number x∈]0,1[.
// More precisely, x∈[ε,1-ε], ε=1/2⁵³. This function is safe
// for concurrent use by multiple goroutines.
func U01() float64 {
	n := prisrc.Uint64() & ((1 << 53) - 1)
	for n == 0 {
		n = prisrc.Uint64()
	}
	return float64(n) / (1 << 53)
}
