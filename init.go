package mym

import (
	cr "crypto/rand"
	"encoding/binary"
	"math/rand"
)

// prisrc -- a private source of pseudo-random numbers `MT19937`.
var prisrc rand.Source64

// init -- initializes `prisrc` with a 64-bit seed obtained from `crypto/rand`.
func init() {
	var buf [8]byte
	var b = buf[:]
	_, err := cr.Read(b)
	if err != nil {
		panic(err)
	}
	seed := binary.LittleEndian.Uint64(b)
	prisrc = MT19937()
	prisrc.Seed(int64(seed))
}
