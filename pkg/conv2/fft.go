package conv2

import (
	"github.com/davidkleiven/gosfft/sfft"
	"gonum.org/v1/gonum/cmplxs"
	"gonum.org/v1/gonum/mat"
)

func realToCmplxPad(f *mat.Dense) *mat.CDense {
	dy, dx := f.Dims()
	dy2, dx2 := dy*2-1, dx*2-1

	out := mat.NewCDense(dy2, dx2, nil)
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			out.Set(i, j, complex(f.At(i, j), 0.0))
		}
	}

	return out
}

func cmplxToRealScale(f *mat.CDense) *mat.Dense {
	dy, dx := f.Dims()

	out := mat.NewDense(dy, dx, nil)
	scale := 1 / (float64(dy) * float64(dx))
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			out.Set(i, j, real(f.At(i, j))*scale)
		}
	}

	return out
}

func FFull(f, g *mat.Dense) *mat.Dense {
	// make sure f is always larger than g
	if g.RawMatrix().Cols > f.RawMatrix().Cols {
		f, g = g, f
	}
	if g.RawMatrix().Rows > f.RawMatrix().Rows {
		panic("none of the inputs is at least as large as the other in every dimension")
	}

	dy, dx := f.Dims()
	dv, du := g.Dims()

	// pad g to size of f
	gpad := mat.NewDense(dy, dx, nil)
	gpad.Slice(
		center(dy-dv),
		center(dy-dv)+dv,
		center(dx-du),
		center(dx-du)+du,
	).(*mat.Dense).
		Copy(g)

	// convert to complex and pad
	cf := realToCmplxPad(f)
	cg := realToCmplxPad(gpad)

	// FFT in place
	fft := sfft.NewFFT2(cf.Dims())
	fft.FFT(cf.RawCMatrix().Data)
	fft.FFT(cg.RawCMatrix().Data)

	// element wise mul, into cf
	cmplxs.Mul(
		cf.RawCMatrix().Data,
		cg.RawCMatrix().Data,
	)

	// IFFT in place
	fft.IFFT(cf.RawCMatrix().Data)

	out := cmplxToRealScale(cf)
	return out.Slice(
		center(dy-dv),
		center(dy-dv)+dy+dv-1,
		center(dx-du),
		center(dx-du)+dx+du-1,
	).(*mat.Dense)
}
