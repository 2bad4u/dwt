[![GoDoc](https://godoc.org/github.com/2bad4u/dwt?status.svg)](https://godoc.org/github.com/2bad4u/dwt)
[![Build Status](https://travis-ci.org/2bad4u/dwt.svg)](https://travis-ci.org/2bad4u/dwt)
[![Coverage Status](https://coveralls.io/repos/2bad4u/dwt/badge.svg?branch=master&service=github)](https://coveralls.io/github/2bad4u/dwt?branch=master)

# dwt

*dwt* is simply a GO port of Gregoire Pau's 'Fast Discrete Bi-orthogonal CDF Wavelet Transform' examples.

See [dwt97.c](http://web.archive.org/web/20120305164605/http://www.embl.de/~gpau/misc/dwt97.c),
[CDF 5/3 Discrete Wavelet Transform: dwt53.c](https://github.com/VadimKirilchuk/jawelet/wiki/CDF-5-3-Discrete-Wavelet-Transform),
[CDF 9/7 Discrete Wavelet Transform: dwt97.c](https://github.com/VadimKirilchuk/jawelet/wiki/CDF-9-7-Discrete-Wavelet-Transform)
or [waveletcdf97](http://www.getreuer.info/home/waveletcdf97) for details.

#### Getting
```
go get github.com/2bad4u/dwt
```

#### Usage
```go

  import "github.com/2bad4u/dwt"

  func foo() {
    xn := make([]float64, 32)
    for i := 0; i < 32; i++ {
    	x[i] = 5.0 + float64(i) + 0.4*float64(i*i) - 0.02*float64(i*i*i)
    }

    // transform xn with CDF 9/7, or CDF 5/3: Fwt53(xn)
    Fwt97(xn)

    // xn contains the transformation result
    ...

    // restore signal with inverse transform, use Iwt53(xn) for CDF 5/3
    Iwt97(xn)

    ...
  }

```
