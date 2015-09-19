dwt
===

*dwt* is simply a GO port of Gregoire Pau's fast discrete bi-orthogonal CDF wavelet transformation.

See [dwt97.c](https://code.google.com/p/axonlib/source/browse/trunk/extern/dwt97.c?r=19)
or [waveletcdf97](http://www.getreuer.info/home/waveletcdf97) for details.

# Installation
```
go get github.com/2bad4u/dwt

```

# Usage
```go

  import "github.com/2bad4u/dwt"

  func foo() {
    xn := make([]float64, 32)
    for i := 0; i < 32; i++ {
    	x[i] = 5.0 + float64(i) + 0.4*float64(i*i) - 0.02*float64(i*i*i)
    }

    // transform xn
    Fwt97(xn)

    // xn contains the transformation result
    ...

    // inverse transformation for signal restoration
    Iwt97(xn)

    ...
  }

```
