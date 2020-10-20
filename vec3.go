package mym

import (
	"math"
)

// Vneg3 -- returns -u.
func Vneg3(u [3]float64) [3]float64 {
	return [3]float64{-u[0], -u[1], -u[2]}
}

// Vadd3 -- returns u+v.
func Vadd3(u, v [3]float64) [3]float64 {
	return [3]float64{u[0] + v[0], u[1] + v[1], u[2] + v[2]}
}

// Vsub3 -- returns u-v.
func Vsub3(u, v [3]float64) [3]float64 {
	return [3]float64{u[0] - v[0], u[1] - v[1], u[2] - v[2]}
}

// Vmul3 -- returns s·u.
func Vmul3(u [3]float64, s float64) [3]float64 {
	return [3]float64{s * u[0], s * u[1], s * u[2]}
}

// Vdiv3 -- returns u/s.
func Vdiv3(u [3]float64, s float64) [3]float64 {
	return [3]float64{u[0] / s, u[1] / s, u[2] / s}
}

// Vdot3 -- returns u·v (scalar, or dot product).
func Vdot3(u, v [3]float64) float64 {
	return u[0]*v[0] + u[1]*v[1] + u[2]*v[2]
}

// Vcrs3 -- returns u⨯v (vector, or cross product).
func Vcrs3(u, v [3]float64) [3]float64 {
	x := u[1]*v[2] - u[2]*v[1]
	y := u[2]*v[0] - u[0]*v[2]
	z := u[0]*v[1] - u[1]*v[0]
	return [3]float64{x, y, z}
}

// Vabs3 -- returns |u| (L2 norm).
func Vabs3(u [3]float64) float64 {
	return math.Hypot(math.Hypot(u[0], u[1]), u[2])
}

// Vhat3 -- returns u/|u|. Returns a zero vector when |u|<ε.
func Vhat3(u [3]float64) [3]float64 {
	s := math.Hypot(math.Hypot(u[0], u[1]), u[2])
	if s >= Epsilon {
		return [3]float64{u[0] / s, u[1] / s, u[2] / s}
	}
	return [3]float64{0, 0, 0}
}

// Vmean3 -- returns the mean vector of u[0],u[1],...,u[len(u)-1].
func Vmean3(u [][3]float64) [3]float64 {
	n := len(u)
	if n == 0 {
		return [3]float64{0, 0, 0}
	}
	xs := AccuSum(n, func(i int) float64 { return u[i][0] })
	ys := AccuSum(n, func(i int) float64 { return u[i][1] })
	zs := AccuSum(n, func(i int) float64 { return u[i][2] })
	return [3]float64{xs / float64(n), ys / float64(n), zs / float64(n)}
}
