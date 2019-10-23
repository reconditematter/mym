package mym

// FiniteIs -- returns true iff `x` is a finite floating point number.
func FiniteIs(x float64) bool {
	return (x - x) == 0
}

// Finite -- returns true iff `x` and `y` are finite floating point numbers.
func Finite(x, y float64) bool {
	return ((x - x) == 0) && ((y - y) == 0)
}

// Sq -- returns x².
func Sq(x float64) float64 {
	return x * x
}

// Cb -- returns x³.
func Cb(x float64) float64 {
	return x * x * x
}
