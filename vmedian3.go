package mym

import (
	"math"
)

// Vmedian3 -- computes the geometric median of u[0],u[1],...,u[len(u)-1].
// This function implements the Weiszfeld iterative algorithm as modified by
// Vardi and Zhang.
//
// Reference: Vardi and Zhang, The multivariate L1-median and associated data depth,
// Proceedings of the National Academy of Sciences Feb 2000, 97 (4) 1423-1426.
//
// DOI: https://doi.org/10.1073/pnas.97.4.1423
func Vmedian3(u [][3]float64) [3]float64 {
	n := len(u)
	if n == 0 {
		return [3]float64{0, 0, 0}
	}
	//
	// initial approximation
	mu := Vmean3(u)
	//
	// convergence test compares the sum of distances between two iterations
	convtest := func(mu1, mu2 [3]float64) bool {
		s1 := AccuSum(len(u), func(i int) float64 { return Vabs3(Vsub3(mu1, u[i])) })
		s2 := AccuSum(len(u), func(i int) float64 { return Vabs3(Vsub3(mu2, u[i])) })
		return math.Abs(s1-s2) < Epsilon*math.Max(s1, s2)
	}
	//
	for iter := 1; iter <= 5000; iter++ {
		eta := 0.0
		S1 := [3]float64{0, 0, 0}
		S2 := 0.0
		R := [3]float64{0, 0, 0}
		//
		for _, v := range u {
			w := Vsub3(v, mu)
			wabs := Vabs3(w)
			//
			if wabs < Epsilon {
				eta += 1
				continue
			}
			//
			S1 = Vadd3(S1, Vdiv3(v, wabs))
			S2 += 1 / wabs
			//
			R = Vadd3(R, Vdiv3(w, wabs))
		}
		//
		T := Vdiv3(S1, S2)
		gamma := math.Min(1, eta/Vabs3(R))
		munew := Vadd3(Vmul3(T, 1-gamma), Vmul3(mu, gamma))
		if convtest(mu, munew) {
			return munew
		}
		//
		mu = munew
	}
	//
	return mu
}
