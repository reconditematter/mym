// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

import (
	"math"
	"testing"
)

// TestComplete12 checks if the Legendre's relation for m=0.00001(0.00001)0.99999
// is satisfied with an accuracy given by `tol`.
func TestComplete12(t *testing.T) {
	const tol = 1.0e-14

	for m := 1; m < 100000; m++ {
		mf := float64(m) / 100000
		mp := 1 - mf
		K, E := Complete12(mf)
		Kp, Ep := Complete12(mp)
		legendre := math.Abs(E*Kp + Ep*K - K*Kp - math.Pi/2)
		if legendre > tol {
			t.Fatalf("legendre > tol: m=%v, legendre=%v, tol=%v", mf, legendre, tol)
		}
	}
}
