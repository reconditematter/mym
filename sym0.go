// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

// Sym0 -- represents a symmetric matrix with zero diagonal.
type Sym0 struct {
	n int
	a []float64
}

// NewSym0 -- returns an n-by-n symmetric matrix with all zero elements.
func NewSym0(n int) Sym0 {
	if n < 1 {
		panic("mym.NewSym0: n < 1")
	}
	return Sym0{n, make([]float64, n*(n+1)/2)}
}

// N -- returns the matrix size.
func (a Sym0) N() int {
	return a.n
}

// Get -- returns a[i,j].
func (a Sym0) Get(i, j int) float64 {
	if !(0 <= i && i < a.n && 0 <= j && j < a.n) {
		panic("mym.Sym0.Get: index")
	}
	if i < j {
		return a.a[i*a.n-(i-1)*i/2+j-i]
	}
	if i > j {
		return a.a[j*a.n-(j-1)*j/2+i-j]
	}
	return 0
}

// Set -- assigns a[i,j] = x.
func (a Sym0) Set(i, j int, x float64) {
	if !(0 <= i && i < a.n && 0 <= j && j < a.n) {
		panic("mym.Sym0.Set: index")
	}
	if i < j {
		a.a[i*a.n-(i-1)*i/2+j-i] = x
		return
	}
	if i > j {
		a.a[j*a.n-(j-1)*j/2+i-j] = x
		return
	}
}
