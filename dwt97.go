package dwt

/**
 *  dwt97.c - Fast discrete biorthogonal CDF 9/7 wavelet forward and inverse transform (lifting implementation)
 *
 *  This code is provided "as is" and is given for educational purposes.
 *  2006 - Gregoire Pau - gregoire.pau@ebi.ac.uk
 */

const (
	p1_97  = -1.586134342
	ip1_97 = -p1_97

	u1_97  = -0.05298011854
	iu1_97 = -u1_97

	p2_97  = 0.8829110762
	ip2_97 = -p2_97

	u2_97  = 0.4435068522
	iu2_97 = -u2_97

	scale97  = 1.149604398
	iscale97 = 1.0 / scale97
)

// Fwt97 performs a bi-orthogonal 9/7 wavelet transformation (lifting implementation)
// of the signal in slice xn. The length of the signal n = len(xn) must be a power of 2.
//
// The input slice xn will be replaced by the transformation:
//
// The first half part of the output signal contains the approximation coefficients.
// The second half part contains the detail coefficients (aka. the wavelets coefficients).
func Fwt97(xn []float64) {
	var a float64
	var i int
	var n = validateLen(xn)

	// predict 1
	a = p1_97
	for i = 1; i < n-2; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[n-1] += 2 * a * xn[n-2]

	// update 1
	a = u1_97
	for i = 2; i < n; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[0] += 2 * a * xn[1]

	// predict 2
	a = p2_97
	for i = 1; i < n-2; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[n-1] += 2 * a * xn[n-2]

	// update 2
	a = u2_97
	for i = 2; i < n; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[0] += 2 * a * xn[1]

	// scale
	a = iscale97
	for i = 0; i < n; i++ {
		if i%2 != 0 {
			xn[i] *= a
		} else {
			xn[i] /= a
		}
	}

	// pack
	tb := make([]float64, n)
	for i = 0; i < n; i++ {
		if i%2 == 0 {
			tb[i/2] = xn[i]
		} else {
			tb[n/2+i/2] = xn[i]
		}
	}
	for i = 0; i < n; i++ {
		xn[i] = tb[i]
	}
}

// Iwt97 performs an inverse bi-orthogonal 9/7 wavelet transformation of xn.
// This is the inverse function of Fwt97 so that Iwt97(Fwt97(xn))=xn for every signal xn of length n.
//
// The length of slice xn must be a power of 2.
//
// The coefficients provided in slice xn are replaced by the original signal.
func Iwt97(xn []float64) {
	var a float64
	var i int
	var n = validateLen(xn)

	// unpack
	tb := make([]float64, n)
	for i = 0; i < n/2; i++ {
		tb[i*2] = xn[i]
		tb[i*2+1] = xn[i+n/2]
	}
	for i = 0; i < n; i++ {
		xn[i] = tb[i]
	}

	// undo scale
	a = scale97
	for i = 0; i < n; i++ {
		if i%2 != 0 {
			xn[i] *= a
		} else {
			xn[i] /= a
		}
	}

	// undo update 2
	a = iu2_97
	for i = 2; i < n; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[0] += 2 * a * xn[1]

	// undo predict 2
	a = ip2_97
	for i = 1; i < n-2; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[n-1] += 2 * a * xn[n-2]

	// undo update 1
	a = iu1_97
	for i = 2; i < n; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[0] += 2 * a * xn[1]

	// undo predict 1
	a = ip1_97
	for i = 1; i < n-2; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[n-1] += 2 * a * xn[n-2]
}
