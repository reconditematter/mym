// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

import (
	"math/rand"
	"sync"
)

// MT19937 -- returns a 64-bit source of uniform pseudo-random numbers.
// The source implements a Mersenne twister algorithm based on
// a large Mersenne prime number 2^19937-1. The default seed is selected
// to produce the same sequence as a default-constructed object
// of type mt19937_64 in `Standard for Programming Language C++`.
// The returned source is safe for concurrent use by multiple goroutines.
func MT19937() rand.Source64 {
	r := &mt19937{}
	r.Seed(5489)
	return r
}

// nmt19937 -- the size of the Mersenne twister's internal state.
const nmt19937 = 312

// mt19937 -- the Mersenne twister's internal state.
type mt19937 struct {
	state [nmt19937]uint64
	index int
	mutex sync.Mutex
}

// Seed -- seeds `r` with `seed`.
func (r *mt19937) Seed(seed int64) {
	const n = nmt19937
	r.mutex.Lock()
	r.state[0] = uint64(seed)
	for i := 1; i < n; i++ {
		r.state[i] = 6364136223846793005*(r.state[i-1]^(r.state[i-1]>>62)) + uint64(i)
	}
	r.index = n
	r.mutex.Unlock()
}

// Int63 -- returns a pseudo-random number in [0,2⁶³-1].
func (r *mt19937) Int63() int64 {
	return int64(r.Uint64() & 0x7FFFFFFFFFFFFFFF)
}

// Uint64 -- returns a pseudo-random number in [0,2⁶⁴-1].
func (r *mt19937) Uint64() uint64 {
	const n = nmt19937
	const m = n / 2
	const hi uint64 = 0xFFFFFFFF80000000
	const lo uint64 = 0x000000007FFFFFFF
	r.mutex.Lock()
	if r.index >= n {
		for i := 0; i < n-m; i++ {
			y := (r.state[i] & hi) | (r.state[i+1] & lo)
			r.state[i] = r.state[i+m] ^ (y >> 1) ^ ((y & 1) * 0xB5026F5AA96619E9)
		}
		for i := n - m; i < n-1; i++ {
			y := (r.state[i] & hi) | (r.state[i+1] & lo)
			r.state[i] = r.state[i+(m-n)] ^ (y >> 1) ^ ((y & 1) * 0xB5026F5AA96619E9)
		}
		y := (r.state[n-1] & hi) | (r.state[0] & lo)
		r.state[n-1] = r.state[m-1] ^ (y >> 1) ^ ((y & 1) * 0xB5026F5AA96619E9)
		r.index = 0
	}
	y := r.state[r.index]
	y ^= (y >> 29) & 0x5555555555555555
	y ^= (y << 17) & 0x71D67FFFEDA60000
	y ^= (y << 37) & 0xFFF7EEE000000000
	y ^= (y >> 43)
	r.index++
	r.mutex.Unlock()
	return y
}
