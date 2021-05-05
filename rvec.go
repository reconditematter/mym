// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

// RVec -- represents a vector of variable size with real (float64) elements.
type RVec struct {
	vec []float64
	lwb int
}

const erriib = "invalid index bound"
const errioob = "index out of bounds"

// imin -- min{a,b}.
func imin(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// imax -- max{a,b}.
func imax(a, b int) int {
	if b < a {
		return a
	}
	return b
}

// Maxdex -- maximum valid index (281474976710656).
const Maxdex = 1 << 48

// Mindex -- minimum valid index (-281474976710656).
const Mindex = -Maxdex

// Genrvec -- generates a new vector with the given lower and upper index bounds.
func Genrvec(lwb, upb int) RVec {
	if lwb < Mindex || lwb > Maxdex || upb < Mindex || upb > Maxdex {
		panic(erriib)
	}
	if upb < lwb {
		return RVec{}
	}
	return RVec{make([]float64, upb-lwb+1), lwb}
}

// Ix -- returns the lower and upper index bounds of `x`.
func (x RVec) Ix() (lwb, upb int) {
	if len(x.vec) == 0 {
		return Maxdex, Mindex
	}
	return x.lwb, x.lwb + (len(x.vec) - 1)
}

// E -- returns the element x[i].
func (x RVec) E(i int) float64 {
	if i < Mindex || i > Maxdex {
		panic(errioob)
	}
	L, U := x.Ix()
	if i < L || i > U {
		return 0
	}
	return x.vec[i-L]
}

// U -- updates the element x[i]=e.
func (x *RVec) U(i int, e float64) {
	if i < Mindex || i > Maxdex {
		panic(errioob)
	}
	L, U := x.Ix()
	//
	if i < L || i > U {
		newL, newU := imin(L, i), imax(U, i)
		newvec := make([]float64, newU-newL+1)
		for k := L; k <= U; k++ {
			newvec[k-newL] = x.vec[k-L]
		}
		x.vec = newvec
		x.lwb = newL
		L = newL
	}
	//
	x.vec[i-L] = e
}

// meet -- returns the intersection of index bounds of two vectors.
func (x RVec) meet(y RVec) (int, int) {
	xL, xU := x.Ix()
	yL, yU := y.Ix()
	return imax(xL, yL), imin(xU, yU)
}

// span -- returns the union of index bounds of two vectors.
func (x RVec) span(y RVec) (int, int) {
	xL, xU := x.Ix()
	yL, yU := y.Ix()
	return imin(xL, yL), imax(xU, yU)
}

// Add -- returns the sum x+y.
func (x RVec) Add(y RVec) RVec {
	L, U := x.span(y)
	z := Genrvec(L, U)
	for i := L; i <= U; i++ {
		z.U(i, x.E(i)+y.E(i))
	}
	return z
}

// Sub -- returns the difference x-y.
func (x RVec) Sub(y RVec) RVec {
	L, U := x.span(y)
	z := Genrvec(L, U)
	for i := L; i <= U; i++ {
		z.U(i, x.E(i)-y.E(i))
	}
	return z
}

// Mul -- returns the product s*x.
func (x RVec) Mul(s float64) RVec {
	L, U := x.Ix()
	z := Genrvec(L, U)
	for i := L; i <= U; i++ {
		z.U(i, x.E(i)*s)
	}
	return z
}

// Div -- returns the product (1/s)*x.
func (x RVec) Div(s float64) RVec {
	L, U := x.Ix()
	z := Genrvec(L, U)
	for i := L; i <= U; i++ {
		z.U(i, x.E(i)/s)
	}
	return z
}

// Dot -- returns the scalar (dot) product of `x` and `y`.
func (x RVec) Dot(y RVec) float64 {
	L, U := x.meet(y)
	s := 0.0
	for i := L; i <= U; i++ {
		s += x.E(i) * y.E(i)
	}
	return s
}
