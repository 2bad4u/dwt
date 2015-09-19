// Package dwt provides CDF wavelet transformation.
// This is as port of Gregoire Pau's C implementation:
// see http://web.archive.org/web/20120305164605/http://www.embl.de/~gpau/misc/dwt97.c for details.
package dwt

import "fmt"

const (
	p1  = -1.586134342
	ip1 = -p1

	u1  = -0.05298011854
	iu1 = -u1

	p2  = 0.8829110762
	ip2 = -p2

	u2  = 0.4435068522
	iu2 = -u2

	scale = 1.149604398
)

//  Fwt97 performs a bi-orthogonal 9/7 wavelet transform (lifting implementation)
//  of input signal xn. The length of the signal n = len(xn), and must be a power of 2.
//
//  The signal xn will be replaced by the transformations output:
//  The first half part of the output signal contains the approximation coefficients.
//  The second half part contains the detail coefficients (aka. the wavelets coefficients).
func Fwt97(xn []float64) {
	var a float64
	var i int
	var n = validateLen(xn)

	a = p1
	for i = 1; i < n-2; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[n-1] += 2 * a * xn[n-2]

	a = u1
	for i = 2; i < n; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[0] += 2 * a * xn[1]

	a = p2
	for i = 1; i < n-2; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[n-1] += 2 * a * xn[n-2]

	a = u2
	for i = 2; i < n; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[0] += 2 * a * xn[1]

	a = 1 / scale
	for i = 0; i < n; i++ {
		if i%2 != 0 {
			xn[i] *= a
		} else {
			xn[i] /= a
		}
	}

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
// The length of the signal n = len(xn), and must be a power of 2. The coefficients provided as xn are
// replaced by the original signal.
func Iwt97(xn []float64) {
	var a float64
	var i int
	var n = validateLen(xn)

	tb := make([]float64, n)

	for i = 0; i < n/2; i++ {
		tb[i*2] = xn[i]
		tb[i*2+1] = xn[i+n/2]
	}
	for i = 0; i < n; i++ {
		xn[i] = tb[i]
	}

	a = scale
	for i = 0; i < n; i++ {
		if i%2 != 0 {
			xn[i] *= a
		} else {
			xn[i] /= a
		}
	}

	a = iu2
	for i = 2; i < n; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[0] += 2 * a * xn[1]

	a = ip2
	for i = 1; i < n-2; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[n-1] += 2 * a * xn[n-2]

	a = iu1
	for i = 2; i < n; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[0] += 2 * a * xn[1]

	a = ip1
	for i = 1; i < n-2; i += 2 {
		xn[i] += a * (xn[i-1] + xn[i+1])
	}
	xn[n-1] += 2 * a * xn[n-2]
}

// validateLen checks that the input length is greater than 1 and a power of 2.
// The function returns the slices length if valid or panics otherwise.
func validateLen(xn []float64) (n int) {
	n = len(xn)
	if n < 2 || !isPowerOfTwo(n) {
		panic(fmt.Sprintf("Signal can't be transformed, the input length %d is not a power of 2!", n))
	}
	return n
}

// isPowerOfTwo returns true if the given number n is a positive power of 2.
func isPowerOfTwo(n int) bool {
	un := uint(n)
	if un == 0 || (un&(un-1)) != 0 {
		return false
	}
	return true
}
