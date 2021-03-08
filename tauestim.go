// Copyright (c) 2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

import (
	"math"
)

// FactorTau -- τ-estimator dispersion scale factor for the N(0,1) Gaussian distribution.
const FactorTau = 1.0 / 0.962

// TauEstim -- computes robust τ-estimators of
// location (μ) and dispersion (σ) for a sample `x`.
//
// Reference: Maronna, Zamar,
// Robust Estimates of Location and Dispersion for High-Dimensional Datasets,
// Technometrics, vol 44 (4), pp 307-317 (2002).
//
// DOI: https://doi.org/10.1198/004017002188618509
func TauEstim(x []float64) (μ, σ float64) {
	μ0, σ0 := MedianMAD(x)
	if !(σ0 > Epsilon) {
		μ, σ = μ0, σ0
		return
	}
	//
	fnR := func(c, x float64) float64 { return math.Min(c*c, x*x) }
	fnW := func(c, x float64) float64 {
		if math.Abs(x) <= c {
			return Sq(1 - Sq(x/c))
		}
		return 0
	}
	//
	n := len(x)
	const c1 = 4.5
	w := make([]float64, n)
	for i := range w {
		w[i] = fnW(c1, (x[i]-μ0)/σ0)
	}
	μ = AccuDot(n, func(i int) float64 { return x[i] }, func(i int) float64 { return w[i] })
	μ /= AccuSum(n, func(i int) float64 { return w[i] })
	//
	const c2 = 3.0
	σ = AccuSum(n, func(i int) float64 { return fnR(c2, (x[i]-μ)/σ0) })
	σ = σ0 * math.Sqrt(σ/float64(n))
	//
	return
}
