package conv2

// The flags chosen below optimize for the following platforms:
// - amd64: compiler host, with OpenMP
// - arm: Raspberry Pi Zero (W)
// - arm64: Raspberry Pi 4
//
// To show the flags which -march=native would produce, run
//
// 	gcc -march=native -E -v - </dev/null 2>&1 | grep cc1
//
// For more details, see
// - https://gist.github.com/fm4dd/c663217935dc17f0fc73c9c81b0aa845
// - https://gcc.gnu.org/onlinedocs/gcc/x86-Options.html
// - https://gcc.gnu.org/onlinedocs/gcc/ARM-Options.html
// - https://gcc.gnu.org/onlinedocs/gcc/AArch64-Options.html

// #cgo CFLAGS: -Wall -Werror -Wextra -pedantic -std=c99
// #cgo CFLAGS: -O2
//
// #cgo amd64 CFLAGS: -march=native
// #cgo amd64 CFLAGS: -fopenmp
// #cgo amd64 LDFLAGS: -fopenmp
//
// #cgo arm CFLAGS: -mfloat-abi=hard -mfpu=vfp -marm -march=armv6kz+fp
//
// #cgo arm64 CFLAGS: -mcpu=cortex-a72 -mtune=cortex-a72
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
