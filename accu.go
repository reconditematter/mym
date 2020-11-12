// Copyright (c) 2019-2020 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

func knuadd(u, v float64) (x, y float64) {
	// Knuth, Seminumerical Algorithms, 3rd ed (1998)
	x = u + v
	up := x - v
	vpp := x - up
	y = (u - up) + (v - vpp)
	return
}

func knumul(u, v float64) (x, y float64) {
	// Knuth, Seminumerical Algorithms, 3rd ed (1998)
	const C = 134217729 // 2^27+1
	up, vp := u*C, v*C
	u1, v1 := (u-up)+up, (v-vp)+vp
	u2, v2 := u-u1, v-v1
	x = u * v
	y = ((((u1 * v1) - x) + (u1 * v2)) + (u2 * v1)) + (u2 * v2)
	return
}

// AccuSum -- computes the sum f(0)+f(1)+...+f(n-1) using a compensated summation algorithm.
func AccuSum(n int, f func(int) float64) float64 {
	// Ogita, Rump, Oishi. Accurate sum and dot product. SIAM Journal on Scientific Computing, 26(6):1955–1988, 2005.
	var p, q, s float64
	for i := 0; i < n; i++ {
		p, q = knuadd(p, f(i))
		s += q
	}
	return p + s
}

// AccuSum2 -- computes the sum f(0,0)+f(0,1)+...+f(m-1,n-1) using a compensated summation algorithm.
func AccuSum2(m, n int, f func(int, int) float64) float64 {
	// Ogita, Rump, Oishi. Accurate sum and dot product. SIAM Journal on Scientific Computing, 26(6):1955–1988, 2005.
	var p, q, s float64
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p, q = knuadd(p, f(i, j))
			s += q
		}
	}
	return p + s
}

// AccuDot -- computes the inner (dot) product f(0)g(0)+f(1)g(1)+...+f(n-1)g(n-1) using a compensated summation algorithm.
func AccuDot(n int, f, g func(int) float64) float64 {
	// Ogita, Rump, Oishi. Accurate sum and dot product. SIAM Journal on Scientific Computing, 26(6):1955–1988, 2005.
	var h, p, q, r, s float64
	for i := 0; i < n; i++ {
		h, r = knumul(f(i), g(i))
		p, q = knuadd(p, h)
		s += q + r
	}
	return p + s
}

// AccuDot2 -- computes the inner (dot) product f(0,0)g(0,0)+f(0,1)g(0,1)+...+f(m-1,n-1)g(m-1,n-1) using a compensated summation algorithm.
func AccuDot2(m, n int, f, g func(int, int) float64) float64 {
	// Ogita, Rump, Oishi. Accurate sum and dot product. SIAM Journal on Scientific Computing, 26(6):1955–1988, 2005.
	var h, p, q, r, s float64
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			h, r = knumul(f(i, j), g(i, j))
			p, q = knuadd(p, h)
			s += q + r
		}
	}
	return p + s
}
