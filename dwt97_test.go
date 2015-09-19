package dwt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIwt97Fwt97(t *testing.T) {

	xn := make([]float64, 32)

	for i := 0; i < 32; i++ {
		xn[i] = 5.0 + float64(i) + 0.4*float64(i*i) - 0.02*float64(i*i*i)
	}
	yn := make([]float64, 32)
	copy(yn, xn)

	fmt.Printf("xn is %v.", xn)

	Fwt97(xn)
	fmt.Printf("Fwt97(xn) is %v.", xn)

	Iwt97(xn)
	fmt.Printf("Iwt97(Fwt97(xn)) is %v.", xn)

	assert.InDeltaSlice(t, yn, xn, 0.0000000001, "Iwt97(Fwt97(xn)) != xn")
}

func TestValidateLen(t *testing.T) {
	assert.NotPanics(t, func() { validateLen(make([]float64, 16)) }, "No panic expected for n=16")

	assert.Panics(t, func() { validateLen(make([]float64, 17)) }, "Should panic for n=17")
	assert.Panics(t, func() { validateLen(make([]float64, 1)) }, "Should panic for n=1!")
}

func TestIsPowerOfTwo(t *testing.T) {
	const bits = 32

	for l, b := 1, 0; b < bits; b++ {
		l = l << 1
		assert.True(t, isPowerOfTwo(l), fmt.Sprintf("True expected, %d is a power of 2!", l))
	}

	assert.False(t, isPowerOfTwo(0), "0 is not a power of 2!")
	assert.False(t, isPowerOfTwo(6), "6 is not a power of 2!")
}
