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

	fmt.Printf("xn is %v\n.", xn)

	Fwt97(xn)
	fmt.Printf("Fwt97(xn) is %v\n.", xn)

	Iwt97(xn)
	fmt.Printf("Iwt97(Fwt97(xn)) is %v\n.", xn)

	assert.InDeltaSlice(t, yn, xn, 0.0000000001, "Iwt97(Fwt97(xn)) != xn")
}
