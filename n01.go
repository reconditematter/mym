// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

import (
	"math"
)

// N01 -- returns a normal (Gaussian) pseudo-random number x
// (μ(x)=0, σ(x)=1). This function is safe for concurrent
// use by multiple goroutines.
func N01() float64 {
	// Knuth, Seminumerical Algorithms, 3rd ed, pp 131-132 (1998).
	const C1 = 1.71552776992141359296037928255754495624159721550514 // √(8/e)
	const C2 = 5.13610166675096593629368227224974583334512346112585 // 4*e^(1/4)
	for {
		u, v := U01(), U01()
		x := C1 * (v - 0.5) / u
		x2 := x * x
		if x2 <= 5-C2*u {
			return x
		}
		if x2 <= -4*math.Log(u) {
			return x
		}
	}
}
