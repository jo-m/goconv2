package conv2

import (
	"gonum.org/v1/gonum/mat"
)

func FullFill(f, g *mat.Dense) *mat.Dense {
	dy, dx := f.Dims()
	dv, du := g.Dims()

	n, m := dx+du-1, dy+dv-1
	out := mat.NewDense(m, n, nil)

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			var sum float64

			for v := 0; v < dv; v++ {
				for u := 0; u < du; u++ {
					if y-v < 0 || y-v >= dy || x-u < 0 || x-u >= dx {
						continue // zero
					}
					sum += g.At(v, u) * f.At(y-v, x-u)
				}
			}

			out.Set(y, x, sum)
		}
	}

	return out
}

func FullWrap(f, g *mat.Dense) *mat.Dense {
	dy, dx := f.Dims()
	dv, du := g.Dims()

	n, m := dx+du-1, dy+dv-1
	out := mat.NewDense(m, n, nil)

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			var sum float64

			for v := 0; v < dv; v++ {
				for u := 0; u < du; u++ {
					sum += g.At(v%dv, u%du) * f.At(wrap(y-v, dy), wrap(x-u, dx))
				}
			}

			out.Set(y, x, sum)
		}
	}

	return out
}

func ValidFill(f, g *mat.Dense) *mat.Dense {
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

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			var sum float64

			for v := 0; v < dv; v++ {
				for u := 0; u < du; u++ {
					sum += g.At(v, u) * f.At(y+dv-1-v, x+du-1-u)
				}
			}

			out.Set(y, x, sum)
		}
	}

	return out
}

func ValidWrap(f, g *mat.Dense) *mat.Dense {
	return ValidFill(f, g)
}

func SameFill(f, g *mat.Dense) *mat.Dense {
	dy, dx := f.Dims()
	dv, du := g.Dims()

	// Would be equivalent here (but waste computations):
	//   out := FullFill(f, g)
	//   return out.Slice(center(dv), center(dv)+dy, center(du), center(du)+dx).(*mat.Dense)

	out := mat.NewDense(dy, dx, nil)

	for y := center(dv); y < center(dv)+dy; y++ {
		for x := center(du); x < center(du)+dx; x++ {
			var sum float64

			for v := 0; v < dv; v++ {
				for u := 0; u < du; u++ {
					if y-v < 0 || y-v >= dy || x-u < 0 || x-u >= dx {
						continue // zero
					}
					sum += g.At(v, u) * f.At(y-v, x-u)
				}
			}

			out.Set(y-center(dv), x-center(du), sum)
		}
	}

	return out
}

func SameWrap(f, g *mat.Dense) *mat.Dense {
	dy, dx := f.Dims()
	dv, du := g.Dims()

	// Would be equivalent here (but waste computations):
	//   out := FullWrap(f, g)
	//   return out.Slice(center(dv), center(dv)+dy, center(du), center(du)+dx).(*mat.Dense)

	out := mat.NewDense(dy, dx, nil)

	for y := center(dv); y < center(dv)+dy; y++ {
		for x := center(du); x < center(du)+dx; x++ {
			var sum float64

			for v := 0; v < dv; v++ {
				for u := 0; u < du; u++ {
					sum += g.At(v%dv, u%du) * f.At(wrap(y-v, dy), wrap(x-u, dx))
				}
			}

			out.Set(y-center(dv), x-center(du), sum)
		}
	}

	return out
}
