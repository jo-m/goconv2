package conv2

import (
	"testing"

	"github.com/jo-m/goconv2/internal/pkg/testutil"
)

func Test_FullFillOpt(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := FullFillOpt(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFillOpt(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFillOpt(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFillOpt(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
}
