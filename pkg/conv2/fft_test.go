package conv2

import (
	"testing"

	"github.com/jo-m/goconv2/internal/pkg/imutil"
	"github.com/jo-m/goconv2/internal/pkg/testutil"
)

func Test_FFull(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := FFull(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/fftconv-pp-full.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FFull(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/fftconv-ii-full.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FFull(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/fftconv-ip-full.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FFull(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/fftconv-pi-full.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_FFull_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 20, 20))
	in1 := imutil.ToMat(imutil.Rand(129038, 20, 20))

	for i := 0; i < b.N; i++ {
		FFull(in0, in1)
	}
}

func Benchmark_FFull_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(129038, 160, 120))

	for i := 0; i < b.N; i++ {
		FFull(in0, in1)
	}
}

func Benchmark_FFull_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		FFull(in0, in1)
	}
}
