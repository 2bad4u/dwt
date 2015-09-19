// Copyright 2015 André Schubert. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package dwt provides CDF 5/3 & 9/7 Wavelet Transformations.
//	This is as port of Gregoire Pau's C implementation dwt97.c and dwt53.c.
package dwt

import "fmt"

// validateLen checks that the input length is greater than 1 and a power of 2.
// The function returns the slice's length if valid or panics otherwise.
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
