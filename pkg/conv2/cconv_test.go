package conv2

import (
	"testing"

	"github.com/jo-m/goconv2/internal/pkg/imutil"
	"github.com/jo-m/goconv2/internal/pkg/testutil"
)

func Test_FullFillC(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := FullFillC(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFillC(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFillC(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFillC(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_FullFillC_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		FullFillC(in0, in1)
	}
}

func Test_ValidFillC(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := ValidFillC(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidFillC(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidFillC(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidFillC(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_ValidFillC_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		ValidFillC(in0, in1)
	}
}
