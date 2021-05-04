// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

import (
	"math"
)

// FiniteIs -- returns true iff `x` is a finite floating point number.
func FiniteIs(x float64) bool {
	return !(math.IsNaN(x) || math.IsInf(x, 0))
}

// Finite -- returns true iff `x` and `y` are finite floating point numbers.
func Finite(x, y float64) bool {
	return !(math.IsNaN(x) || math.IsInf(x, 0) || math.IsNaN(y) || math.IsInf(y, 0))
}

// Sq -- returns x².
func Sq(x float64) float64 {
	return x * x
}

// Cb -- returns x³.
func Cb(x float64) float64 {
	return x * x * x
}
