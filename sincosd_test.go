// Copyright (c) 2019-2020 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

import (
	"math"
	"testing"
)

// TestSinCosD checks if the maximum absolute difference between SinCosD and math.Sincos
// for m=-360°(1")+360° is less than a given tolerance `tol`.
func TestSinCosD(t *testing.T) {
	const tol = 1.0e-15

	for m := -(360 * 60 * 60); m <= +(360 * 60 * 60); m++ {
		mf := float64(m) / (360 * 60 * 60)
		sd, cd := SinCosD(mf)
		s, c := math.Sincos(mf * (math.Pi / 180))
		err := math.Max(math.Abs(sd-s), math.Abs(cd-c))
		if err > tol {
			t.Fatalf("sincosd err > tol: m=%v, err=%v, tol=%v", mf, err, tol)
		}
	}
}
