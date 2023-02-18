package conv2

import (
	"testing"

	"github.com/jo-m/goconv2/internal/pkg/imutil"
	"github.com/jo-m/goconv2/internal/pkg/testutil"
)

func Test_FullFFT(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := FullFFT(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/fftconv-pp-full.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFFT(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/fftconv-ii-full.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFFT(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/fftconv-ip-full.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFFT(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/fftconv-pi-full.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_FullFFT_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, patchW, patchH))
	in1 := imutil.ToMat(imutil.Rand(129038, patchW, patchH))

	for i := 0; i < b.N; i++ {
		FullFFT(in0, in1)
	}
}

func Benchmark_FullFFT_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(129038, imgW, imgH))

	for i := 0; i < b.N; i++ {
		FullFFT(in0, in1)
	}
}

func Benchmark_FullFFT_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		FullFFT(in0, in1)
	}
}
