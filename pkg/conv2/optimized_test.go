package conv2

import (
	"testing"

	"github.com/jo-m/goconv2/internal/pkg/imutil"
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

func Benchmark_FullFillOpt_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		FullFillOpt(in0, in1)
	}
}

func Test_ValidFillOpt(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := ValidFillOpt(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidFillOpt(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidFillOpt(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidFillOpt(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_ValidFillOpt_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		ValidFillOpt(in0, in1)
	}
}
