// Copyright (c) 2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

import (
	"math"
)

func oddis(x int) bool {
	return x&1 == 1
}

// x=y or (nan(x) and nan(y))
func f64EQ(x, y float64) bool {
	return (x == y) || (math.IsNaN(x) && math.IsNaN(y))
}

// x<y or (nan(x) and ~nan(y))
func f64LT(x, y float64) bool {
	return (x < y) || (math.IsNaN(x) && !math.IsNaN(y))
}

// algorithm 489 SELECT
// Floyd, Rivest, CACM, vol.18 (3), 1975
func select489(x []float64, L, R, K int) {
	for L < R {
		t := x[K]
		i, j := L, R
		x[L], x[K] = x[K], x[L]
		if f64LT(t, x[R]) {
			x[R], x[L] = x[L], x[R]
		}
		for i < j {
			x[i], x[j] = x[j], x[i]
			i++
			j--
			for f64LT(x[i], t) {
				i++
			}
			for f64LT(t, x[j]) {
				j--
			}

		}
		if f64EQ(x[L], t) {
			x[L], x[j] = x[j], x[L]
		} else {
			j++
			x[j], x[R] = x[R], x[j]
		}
		if j <= K {
			L = j + 1
		}
		if K <= j {
			R = j - 1
		}
	}
}

// low and high medians
func medians(x []float64) (lo, hi float64) {
	n := len(x)
	var m int
	if oddis(n) {
		m = (n + 1) / 2
		select489(x, 0, n-1, m-1)
		lo = x[m-1]
		hi = lo
	} else {
		m = n / 2
		select489(x, 0, n-1, m-1)
		lo = x[m-1]
		select489(x, 0, n-1, m)
		hi = x[m]
	}
	return
}
