package imutil

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"math/rand"
	"os"

	"gonum.org/v1/gonum/mat"
)

func Load(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func Dump(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}

func Sub(img image.Image, r image.Rectangle) (image.Image, error) {
	iface, ok := img.(interface {
		SubImage(r image.Rectangle) image.Image
	})

	if !ok {
		return nil, errors.New("img does not implement SubImage()")
	}

	return iface.SubImage(r), nil
}

func ToGray(img image.Image) *image.Gray {
	ret := image.NewGray(img.Bounds())
	draw.Draw(ret, ret.Bounds(), img, img.Bounds().Min, draw.Src)

	return ret
}

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

func FromMatNorm(m *mat.Dense) *image.Gray {
	normalized := matNormalize(m)
	return FromMat(normalized)
}

func Rand(seed int64, w, h int) *image.Gray {
	src := rand.NewSource(seed)
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
