package conv2

// Optimization flags: see
// - https://gcc.gnu.org/onlinedocs/gcc/x86-Options.html
// - https://gcc.gnu.org/onlinedocs/gcc/ARM-Options.html#ARM-Options
// - https://gcc.gnu.org/onlinedocs/gcc/AArch64-Options.html#AArch64-Options

// #cgo CFLAGS: -O3
// #cgo amd64 CFLAGS: -march=skylake -mtune=skylake
// #cgo arm CFLAGS: -mcpu=cortex-a53 -mfpu=neon-vfpv4 -mtune=cortex-a53
// #cgo arm64 CFLAGS: -march=armv8-a+crc -mcpu=cortex-a72 -mtune=cortex-a72
// #include "cconv.h"
import "C"

import (
	"gonum.org/v1/gonum/mat"
)

// FullFillC returns the 2-dimensional convolution of f and g.
//
// Implemented in C (cgo).
// It is equivalent to scipy.signal.convolve2d(f, g, mode="full", boundary="fill", fillvalue=0).
// See https://docs.scipy.org/doc/scipy/reference/generated/scipy.signal.convolve2d.html.
func FullFillC(f, g *mat.Dense) *mat.Dense {
	dy, dx := f.Dims()
	dv, du := g.Dims()

	n, m := dx+du-1, dy+dv-1
	out := mat.NewDense(m, n, nil)

	fd, fs := f.RawMatrix().Data, f.RawMatrix().Stride
	gd, gs := g.RawMatrix().Data, g.RawMatrix().Stride
	outd, outs := out.RawMatrix().Data, out.RawMatrix().Stride

	C.FullFillC(
		C.int(dy),
		C.int(dx),
		C.int(dv),
		C.int(du),
		C.int(n),
		C.int(m),
		C.int(fs),
		C.int(gs),
		C.int(outs),
		(*C.float64)(&fd[0]),
		(*C.float64)(&gd[0]),
		(*C.float64)(&outd[0]),
	)

	return out
}

// ValidFillC returns the 2-dimensional convolution of f and g.
//
// Implemented in C (cgo).
// It is equivalent to scipy.signal.convolve2d(f, g, mode="valid", boundary="fill", fillvalue=0).
// See https://docs.scipy.org/doc/scipy/reference/generated/scipy.signal.convolve2d.html.
func ValidFillC(f, g *mat.Dense) *mat.Dense {
	// make sure f is always larger than g
	if g.RawMatrix().Cols > f.RawMatrix().Cols {
		f, g = g, f
	}
	if g.RawMatrix().Rows > f.RawMatrix().Rows {
		panic("none of the inputs is at least as large as the other in every dimension")
	}

	dy, dx := f.Dims()
	dv, du := g.Dims()

	n, m := dx-du+1, dy-dv+1
	out := mat.NewDense(m, n, nil)

	fd, fs := f.RawMatrix().Data, f.RawMatrix().Stride
	gd, gs := g.RawMatrix().Data, g.RawMatrix().Stride
	outd, outs := out.RawMatrix().Data, out.RawMatrix().Stride

	C.ValidFillC(
		C.int(dy),
		C.int(dx),
		C.int(dv),
		C.int(du),
		C.int(n),
		C.int(m),
		C.int(fs),
		C.int(gs),
		C.int(outs),
		(*C.float64)(&fd[0]),
		(*C.float64)(&gd[0]),
		(*C.float64)(&outd[0]),
	)

	return out
}
