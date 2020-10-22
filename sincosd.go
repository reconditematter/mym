// Copyright (c) 2019-2020 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.
// Original C code has the following:
// Cephes Math Library Release 2.1
// Copyright 1987 by Stephen L. Moshier

package mym

import (
	"math"
)

var sintbl = [92]float64{
	0.00000000000000000000e+0,
	1.74524064372835128194e-2,
	3.48994967025009716460e-2,
	5.23359562429438327221e-2,
	6.97564737441253007760e-2,
	8.71557427476581735581e-2,
	1.04528463267653471400e-1,
	1.21869343405147481113e-1,
	1.39173100960065444112e-1,
	1.56434465040230869010e-1,
	1.73648177666930348852e-1,
	1.90808995376544812405e-1,
	2.07911690817759337102e-1,
	2.24951054343864998051e-1,
	2.41921895599667722560e-1,
	2.58819045102520762349e-1,
	2.75637355816999185650e-1,
	2.92371704722736728097e-1,
	3.09016994374947424102e-1,
	3.25568154457156668714e-1,
	3.42020143325668733044e-1,
	3.58367949545300273484e-1,
	3.74606593415912035415e-1,
	3.90731128489273755062e-1,
	4.06736643075800207754e-1,
	4.22618261740699436187e-1,
	4.38371146789077417453e-1,
	4.53990499739546791560e-1,
	4.69471562785890775959e-1,
	4.84809620246337029075e-1,
	5.00000000000000000000e-1,
	5.15038074910054210082e-1,
	5.29919264233204954047e-1,
	5.44639035015027082224e-1,
	5.59192903470746830160e-1,
	5.73576436351046096108e-1,
	5.87785252292473129169e-1,
	6.01815023152048279918e-1,
	6.15661475325658279669e-1,
	6.29320391049837452706e-1,
	6.42787609686539326323e-1,
	6.56059028990507284782e-1,
	6.69130606358858213826e-1,
	6.81998360062498500442e-1,
	6.94658370458997286656e-1,
	7.07106781186547524401e-1,
	7.19339800338651139356e-1,
	7.31353701619170483288e-1,
	7.43144825477394235015e-1,
	7.54709580222771997943e-1,
	7.66044443118978035202e-1,
	7.77145961456970879980e-1,
	7.88010753606721956694e-1,
	7.98635510047292846284e-1,
	8.09016994374947424102e-1,
	8.19152044288991789684e-1,
	8.29037572555041692006e-1,
	8.38670567945424029638e-1,
	8.48048096156425970386e-1,
	8.57167300702112287465e-1,
	8.66025403784438646764e-1,
	8.74619707139395800285e-1,
	8.82947592858926942032e-1,
	8.91006524188367862360e-1,
	8.98794046299166992782e-1,
	9.06307787036649963243e-1,
	9.13545457642600895502e-1,
	9.20504853452440327397e-1,
	9.27183854566787400806e-1,
	9.33580426497201748990e-1,
	9.39692620785908384054e-1,
	9.45518575599316810348e-1,
	9.51056516295153572116e-1,
	9.56304755963035481339e-1,
	9.61261695938318861916e-1,
	9.65925826289068286750e-1,
	9.70295726275996472306e-1,
	9.74370064785235228540e-1,
	9.78147600733805637929e-1,
	9.81627183447663953497e-1,
	9.84807753012208059367e-1,
	9.87688340595137726190e-1,
	9.90268068741570315084e-1,
	9.92546151641322034980e-1,
	9.94521895368273336923e-1,
	9.96194698091745532295e-1,
	9.97564050259824247613e-1,
	9.98629534754573873784e-1,
	9.99390827019095730006e-1,
	9.99847695156391239157e-1,
	1.00000000000000000000e+0,
	9.99847695156391239157e-1,
}

// SinCosD -- returns the sine and cosine of the degree argument `x`.
// This function returns (NaN,NaN) when x∉[-360,360].
func SinCosD(x float64) (sin, cos float64) {
	if !(-360 <= x && x <= 360) {
		sin, cos = math.NaN(), math.NaN()
		return
	}
	//
	var xsign, ssign, csign int
	if x < 0 {
		xsign = -1
		x = -x
	}
	ix := int(x + 0.5)
	z := x - float64(ix)
	y := z * z
	//
	if ix <= 180 {
		ssign = 1
		csign = 1
	} else {
		ssign = -1
		csign = -1
		ix -= 180
	}
	if ix > 90 {
		csign = -csign
		ix = 180 - ix
	}
	//
	sx := sintbl[ix]
	if ssign < 0 {
		sx = -sx
	}
	//
	cx := sintbl[90-ix]
	if csign < 0 {
		cx = -cx
	}
	//
	sz := ((1.34959795251974073996e-11*y-8.86096155697856783296e-7)*y + 1.74532925199432957214e-2) * z
	cz := 1.0 - ((3.92582397764340914444e-14*y-3.86632385155548605680e-9)*y+1.52308709893354299569e-4)*y
	y = sx*cz + cx*sz
	if xsign < 0 {
		y = -y
	}
	//
	sin = y
	cos = cx*cz - sx*sz
	return
}
