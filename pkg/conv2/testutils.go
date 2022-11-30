package conv2

import (
	"testing"

	"github.com/jo-m/goconv2/internal/pkg/testutil"
	"github.com/stretchr/testify/require"
	"gonum.org/v1/gonum/mat"
)

func loadImgAndPatch(t *testing.T) (*mat.Dense, *mat.Dense) {
	img := testutil.LoadImgToMat(t, "testdata/img.png")
	x, y, w, h := 3, 6, 8, 9
	patch := img.Slice(y, y+h, x, x+w).(*mat.Dense)

	a, b := patch.Dims()
	require.Equal(t, a, h)
	require.Equal(t, b, w)

	patchTruth := testutil.LoadMat64Txt(t, "testdata/gen/patch.txt")
	testutil.AssertMatEqual(t, patch, patchTruth)

	return img, patch
}
