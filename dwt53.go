package dwt

/**
 *  dwt53.c - Fast discrete biorthogonal CDF 5/3 wavelet forward and inverse transform (lifting implementation)
 *
 *  This code is provided "as is" and is given for educational purposes.
 *  2007 - Gregoire Pau - gregoire.pau@ebi.ac.uk
 */

const (
	p1_53  = -0.5
	ip1_53 = -p1_53

	u1_53  = 0.25
	iu1_53 = -u1_53

	scale53  = 1.4142135623730951 // math.Sqrt(2.0)
	iscale53 = 1 / scale53
)

//  Fwt53 performs a bi-orthogonal 5/3 wavelet transformation (lifting implementation)
//  of the signal in slice xn. The length of the signal n = len(xn) must be a power of 2.
//
//  The input in slice xn will be replaced by the transformation:
//
//  The first half part of the output signal contains the approximation coefficients.
//  The second half part contains the detail coefficients (aka. the wavelets coefficients).
func Fwt53(xn []float64) {
	var a float64
	var i int
	n := validateLen(xn)

	// predict 1
	a = p1_53
	for i = 1; i < n-2; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[n-1] += 2 * a * xn[n-2]

	// update 1
	a = u1_53
	for i = 2; i < n; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[0] += 2 * a * xn[1]

	// scale
	a = scale53
	for i = 0; i < n; i++ {
		if i%2 != 0 {
			xn[i] = xn[i] * a
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

// Iwt53 performs an inverse bi-orthogonal 5/3 wavelet transformation of xn.
// This is the inverse function of Fwt53 so that Iwt53(Fwt53(xn))=xn for every signal xn of length n.
//
// The length of slice xn must be a power of 2.
//
// The coefficients provided in slice xn are replaced by the original signal.
func Iwt53(xn []float64) {
	var a float64
	var i int
	n := validateLen(xn)

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
	a = iscale53
	for i = 0; i < n; i++ {
		if i%2 != 0 {
			xn[i] *= a
		} else {
			xn[i] /= a
		}
	}

	// undo update 1
	a = iu1_53
	for i = 2; i < n; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[0] += 2 * a * xn[1]

	// undo predict 1
	a = ip1_53
	for i = 1; i < n-2; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[n-1] += 2 * a * xn[n-2]
}
