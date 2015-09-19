package dwt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
