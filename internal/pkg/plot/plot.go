// Package plot in an internal package that plots. Used for development and debugging.
package plot

import (
	"math/cmplx"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// GridMatrix is a type used to be able to plot a 2D matrix with Gonum's heatmap
type GridMatrix struct {
	mat mat.Matrix
}

// Dims return the size of the underlying matrix
func (g GridMatrix) Dims() (int, int) { r, c := g.mat.Dims(); return c, r }

// X returns the coordinate value corresponding to column c
func (g GridMatrix) X(c int) float64 { return float64(c) }

// Y returns the coordinate value corresponiding to the row r
func (g GridMatrix) Y(r int) float64 { return float64(r) }

// Z returns the function value corresponding to the (r, c) element in the underlying matrix
func (g GridMatrix) Z(c, r int) float64 { return g.mat.At(r, c) }

// Mat plots a matrix to an image file.
func Mat(path string, mat mat.Matrix) {
	plt := plot.New()
	grid := GridMatrix{mat}
	plt.Add(plotter.NewHeatMap(grid, palette.Heat(10, 1)))

	if err := plt.Save(10*vg.Centimeter, 10*vg.Centimeter, path); err != nil {
		panic(err)
	}
}

// CMat plots a complex matrix to an image file.
func CMat(path string, cmat mat.CMatrix) {
	di, dj := cmat.Dims()

	amp := mat.NewDense(di, dj, nil)
	for i := 0; i < di; i++ {
		for j := 0; j < dj; j++ {
			amp.Set(i, j, cmplx.Abs(cmat.At(i, j)))
		}
	}

	plt := plot.New()
	grid := GridMatrix{mat: amp}
	plt.Add(plotter.NewHeatMap(grid, palette.Heat(10, 1)))

	if err := plt.Save(10*vg.Centimeter, 10*vg.Centimeter, path); err != nil {
		panic(err)
	}
}
