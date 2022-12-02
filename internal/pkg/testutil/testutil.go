package testutil

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/jo-m/goconv2/internal/pkg/imutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gonum.org/v1/gonum/mat"
)

func LoadImg(t *testing.T, path string) image.Image {
	img, err := imutil.Load(path)
	require.NoError(t, err)

	return img
}

func LoadImgToMat(t *testing.T, path string) *mat.Dense {
	img := LoadImg(t, path)
	gray := imutil.ToGray(img)
	return imutil.ToMat(gray)
}

func LoadMat64Txt(t *testing.T, path string) *mat.Dense {
	data, err := os.ReadFile(path)
	require.NoError(t, err)

	cols := strings.Split(strings.TrimSpace(string(data)), "\n")
	require.GreaterOrEqual(t, len(cols), 1)

	r := len(cols)
	c := len(strings.Split(cols[0], " "))

	mat := mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		cells := strings.Split(cols[i], " ")
		require.Equal(t, len(cells), c)

		for j := 0; j < c; j++ {
			val, err := strconv.ParseFloat(cells[j], 64)
			require.NoError(t, err)

			mat.Set(i, j, val)
		}
	}

	return mat
}

func AssertMatEqual(t *testing.T, expected, actual mat.Matrix) {
	m, n := expected.Dims()
	k, l := actual.Dims()

	assert.Equal(t, m, k, "rows")
	assert.Equal(t, n, l, "cols")

	if m != k || n != l {
		return
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			assert.InDelta(t, expected.At(i, j), actual.At(i, j), 0.0000000000001, fmt.Sprintf("cell at (%d,%d)", i, j))
		}
	}
}
