// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package mym

import (
	cr "crypto/rand"
	"encoding/binary"
	"math/rand"
	"sync"
)

// prisrc -- a private source of pseudo-random numbers `MT19937`.
var prisrc rand.Source64
var onesrc sync.Once

// inisrc -- initializes `prisrc` with a 64-bit seed obtained from `crypto/rand`.
func inisrc() {
	onesrc.Do(func() {
		var buf [8]byte
		var b = buf[:]
		_, err := cr.Read(b)
		if err != nil {
			panic(err)
		}
		seed := binary.LittleEndian.Uint64(b)
		prisrc = MT19937()
		prisrc.Seed(int64(seed))
	})
}
