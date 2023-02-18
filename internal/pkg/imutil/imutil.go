// Package imutil is an internal package which contains various utilities for loading, storing, and generating image files.
package imutil

import (
	"errors"
	"image"
	"image/color"
	"image/draw"

	// import JPEG decoder
	_ "image/jpeg"
	"image/png"
	"math/rand"
	"os"

	"gonum.org/v1/gonum/mat"
)

// Load loads an image from a file path.
func Load(path string) (image.Image, error) {
	// #nosec G304
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// #nosec G307
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// Dump dumps an image to a file.
func Dump(path string, img image.Image) error {
	// #nosec G304
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	// #nosec G307
	defer f.Close()

	return png.Encode(f, img)
}

// Sub selects a subregion from an image and returns a new image. Shares memory.
func Sub(img image.Image, r image.Rectangle) (image.Image, error) {
	iface, ok := img.(interface {
		SubImage(r image.Rectangle) image.Image
	})

	if !ok {
		return nil, errors.New("img does not implement SubImage()")
	}

	return iface.SubImage(r), nil
}

// ToGray converts an image to a grayscale image.
func ToGray(img image.Image) *image.Gray {
	ret := image.NewGray(img.Bounds())
	draw.Draw(ret, ret.Bounds(), img, img.Bounds().Min, draw.Src)

	return ret
}

// ToMat converts a grayscale image to a float matrix with values in [0, 1].
func ToMat(img *image.Gray) *mat.Dense {
	rect := img.Bounds()
	ret := mat.NewDense(rect.Dy(), rect.Dx(), nil)

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			px := img.GrayAt(x, y)
			ret.Set(y-rect.Min.Y, x-rect.Min.X, float64(px.Y)/255)
		}
	}

	return ret
}

// FromMat converts a float [0, 1] matrix to a grayscale image.
func FromMat(m *mat.Dense) *image.Gray {
	r, c := m.Dims()

	img := image.NewGray(image.Rect(0, 0, c, r))
	rect := img.Bounds()
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			val := m.At(y-rect.Min.Y, x-rect.Min.X)
			img.Set(x, y, color.Gray{Y: uint8(val * 255)})
		}
	}

	return img
}

func matNormalize(m *mat.Dense) *mat.Dense {
	r, c := m.Dims()
	ret := mat.NewDense(r, c, nil)

	min := mat.Min(m)
	max := mat.Max(m)

	if min == max {
		return ret // zero
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			ret.Set(i, j, (m.At(i, j)-min)/(max-min))
		}
	}

	return ret
}

// FromMatNorm normalizes a float matrix linearly to [0, 1] and then calls FromMat() on the result.
func FromMatNorm(m *mat.Dense) *image.Gray {
	normalized := matNormalize(m)
	return FromMat(normalized)
}

// Rand generates a random grayscale image.
func Rand(seed int64, w, h int) *image.Gray {
	src := rand.NewSource(seed)
	// #nosec G404
	rnd := rand.New(src)

	rect := image.Rect(0, 0, w, h)
	img := image.NewGray(rect)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.Gray{Y: uint8(rnd.Int())})
		}
	}

	return img
}
