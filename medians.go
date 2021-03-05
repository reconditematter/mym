// Copyright (c) 2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

import (
	"math"
)

// Medians -- returns the low and high medians of a sample `x`.
func Medians(x []float64) (lomed, himed float64) {
	if len(x) == 0 {
		lomed, himed = math.NaN(), math.NaN()
		return
	}
	//
	y := make([]float64, len(x))
	copy(y, x)
	lomed, himed = medians(y)
	return
}

// FactorMAD -- MAD scale factor for the N(0,1) Gaussian distribution.
const FactorMAD = 1.48260221850560186054707652936042343132670320259031

// MedianMAD -- computes the median (`med`) and the median absolute deviation (`mad`) of a sample `x`.
func MedianMAD(x []float64) (med, mad float64) {
	if len(x) == 0 {
		med, mad = math.NaN(), math.NaN()
		return
	}
	//
	y := make([]float64, len(x))
	copy(y, x)
	lo, hi := medians(y)
	med = lo + (hi-lo)/2
	//
	for k, yk := range y {
		y[k] = math.Abs(yk - med)
	}
	lo, hi = medians(y)
	mad = lo + (hi-lo)/2
	return
}

// Summary5 -- computes the five-point summary of a sample `x`:
//
//    s[0] - minimum
//    s[1] - lower hinge
//    s[2] - median
//    s[3] - upper hinge
//    s[4] - maximum
func Summary5(x []float64) (s [5]float64) {
	var (
		n, m, k int
	)
	n = len(x)
	if n < 5 {
		s[0], s[1], s[2], s[3], s[4] = math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN()
		return
	}
	//
	y := make([]float64, n)
	copy(y, x)
	//
	n1 := n - 1
	select489(y, 0, n1, 0)
	s[0] = y[0]
	select489(y, 0, n1, n1)
	s[4] = y[n1]
	//
	if oddis(n) {
		m = (n + 1) / 2
		select489(y, 0, n1, m-1)
		// median
		s[2] = y[m-1]
	} else {
		m = n / 2
		select489(y, 0, n1, m-1)
		// low median
		lo := y[m-1]
		select489(y, 0, n1, m)
		// high median
		hi := y[m]
		// median
		s[2] = lo + (hi-lo)/2
	}
	//
	if oddis(m) {
		k = (m + 1) / 2
		select489(y, 0, n1, k-1)
		// lower hinge
		s[1] = y[k-1]
		select489(y, 0, n1, n-k)
		// upper hinge
		s[3] = y[n-k]
	} else {
		k = m / 2
		select489(y, 0, n1, k-1)
		lo := y[k-1]
		select489(y, 0, n1, k)
		hi := y[k]
		// lower hinge
		s[1] = lo + (hi-lo)/2
		select489(y, 0, n1, n-k-1)
		lo = y[n-k-1]
		select489(y, 0, n1, n-k)
		hi = y[n-k]
		// upper hinge
		s[3] = lo + (hi-lo)/2
	}
	return
}
