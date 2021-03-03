// Copyright (c) 2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

import (
	"math"
)

// Medians -- returns the low and high medians of `x`.
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

// MedianMAD -- computes the median (`med`) and the median absolute deviation (`mad`) of `x`.
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
