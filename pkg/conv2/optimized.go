package conv2

import (
	"gonum.org/v1/gonum/mat"
)

// FullFillOpt returns the 2-dimensional convolution of f and g.
//
// Slightly optimized Go implementation.
// It is equivalent to scipy.signal.convolve2d(f, g, mode="full", boundary="fill", fillvalue=0).
// See https://docs.scipy.org/doc/scipy/reference/generated/scipy.signal.convolve2d.html.
func FullFillOpt(f, g *mat.Dense) *mat.Dense {
	dy, dx := f.Dims()
	dv, du := g.Dims()

	n, m := dx+du-1, dy+dv-1
	out := mat.NewDense(m, n, nil)

	fd, fs := f.RawMatrix().Data, f.RawMatrix().Stride
	gd, gs := g.RawMatrix().Data, g.RawMatrix().Stride

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			var sum float64

			for v := 0; v < dv; v++ {
				yv := y - v
				if yv < 0 || yv >= dy {
					continue // zero
				}

				vgs := v * gs
				yvfs := yv * fs

				for u := 0; u < du; u++ {
					xu := x - u
					if xu < 0 || xu >= dx {
						continue // zero
					}

					sum += gd[vgs+u] * fd[yvfs+xu]
				}
			}

			out.Set(y, x, sum)
		}
	}

	return out
}

// ValidFillOpt returns the 2-dimensional convolution of f and g.
//
// Slightly optimized Go implementation.
// It is equivalent to scipy.signal.convolve2d(f, g, mode="valid", boundary="fill", fillvalue=0).
// See https://docs.scipy.org/doc/scipy/reference/generated/scipy.signal.convolve2d.html.
func ValidFillOpt(f, g *mat.Dense) *mat.Dense {
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

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			var sum float64

			ydv1 := y + dv - 1
			xduu := x + du - 1

			for v := 0; v < dv; v++ {
				vgs := v * gs
				ydvvfs := (ydv1 - v) * fs

				for u := 0; u < du; u++ {
					sum += gd[vgs+u] * fd[ydvvfs+(xduu-u)]
				}
			}

			out.Set(y, x, sum)
		}
	}

	return out
}
