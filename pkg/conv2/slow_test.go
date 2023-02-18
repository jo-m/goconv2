package conv2

import (
	"testing"

	"github.com/jo-m/goconv2/internal/pkg/imutil"
	"github.com/jo-m/goconv2/internal/pkg/testutil"
)

func Test_FullFillSlow(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := FullFillSlow(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFillSlow(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFillSlow(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFillSlow(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_FullFillSlow_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, patchW, patchH))
	in1 := imutil.ToMat(imutil.Rand(129038, patchW, patchH))

	for i := 0; i < b.N; i++ {
		FullFillSlow(in0, in1)
	}
}

func Benchmark_FullFillSlow_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(129038, imgW, imgH))

	for i := 0; i < b.N; i++ {
		FullFillSlow(in0, in1)
	}
}

func Benchmark_FullFillSlow_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		FullFillSlow(in0, in1)
	}
}

func Benchmark_FullFillSlow_PI(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		FullFillSlow(in1, in0)
	}
}

func Test_FullWrapSlow(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := FullWrapSlow(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-full-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullWrapSlow(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-full-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullWrapSlow(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-full-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullWrapSlow(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-full-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_FullWrapSlow_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, patchW, patchH))
	in1 := imutil.ToMat(imutil.Rand(129038, patchW, patchH))

	for i := 0; i < b.N; i++ {
		FullWrapSlow(in0, in1)
	}
}

func Benchmark_FullWrapSlow_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(129038, imgW, imgH))

	for i := 0; i < b.N; i++ {
		FullWrapSlow(in0, in1)
	}
}

func Benchmark_FullWrapSlow_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		FullWrapSlow(in0, in1)
	}
}

func Benchmark_FullWrapSlow_PI(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		FullWrapSlow(in1, in0)
	}
}

func Test_ValidFillSlow(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := ValidFillSlow(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidFillSlow(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidFillSlow(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidFillSlow(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_ValidFillSlow_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, patchW, patchH))
	in1 := imutil.ToMat(imutil.Rand(129038, patchW, patchH))

	for i := 0; i < b.N; i++ {
		ValidFillSlow(in0, in1)
	}
}

func Benchmark_ValidFillSlow_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(129038, imgW, imgH))

	for i := 0; i < b.N; i++ {
		ValidFillSlow(in0, in1)
	}
}

func Benchmark_ValidFillSlow_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		ValidFillSlow(in0, in1)
	}
}

func Benchmark_ValidFillSlow_PI(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		ValidFillSlow(in1, in0)
	}
}

func Test_ValidWrapSlow(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := ValidWrapSlow(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-valid-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidWrapSlow(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-valid-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidWrapSlow(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-valid-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidWrapSlow(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-valid-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_ValidWrapSlow_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, patchW, patchH))
	in1 := imutil.ToMat(imutil.Rand(129038, patchW, patchH))

	for i := 0; i < b.N; i++ {
		ValidWrapSlow(in0, in1)
	}
}

func Benchmark_ValidWrapSlow_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(129038, imgW, imgH))

	for i := 0; i < b.N; i++ {
		ValidWrapSlow(in0, in1)
	}
}

func Benchmark_ValidWrapSlow_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		ValidWrapSlow(in0, in1)
	}
}

func Benchmark_ValidWrapSlow_PI(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		ValidWrapSlow(in1, in0)
	}
}

func Test_SameFillSlow(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := SameFillSlow(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-same-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = SameFillSlow(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-same-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = SameFillSlow(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-same-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = SameFillSlow(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-same-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_SameFillSlow_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, patchW, patchH))
	in1 := imutil.ToMat(imutil.Rand(129038, patchW, patchH))

	for i := 0; i < b.N; i++ {
		SameFillSlow(in0, in1)
	}
}

func Benchmark_SameFillSlow_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(129038, imgW, imgH))

	for i := 0; i < b.N; i++ {
		SameFillSlow(in0, in1)
	}
}

func Benchmark_SameFillSlow_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		SameFillSlow(in0, in1)
	}
}

func Benchmark_SameFillSlow_PI(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		SameFillSlow(in1, in0)
	}
}

func Test_SameWrapSlow(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := SameWrapSlow(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-same-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = SameWrapSlow(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-same-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = SameWrapSlow(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-same-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = SameWrapSlow(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-same-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_SameWrapSlow_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, patchW, patchH))
	in1 := imutil.ToMat(imutil.Rand(129038, patchW, patchH))

	for i := 0; i < b.N; i++ {
		SameWrapSlow(in0, in1)
	}
}

func Benchmark_SameWrapSlow_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(129038, imgW, imgH))

	for i := 0; i < b.N; i++ {
		SameWrapSlow(in0, in1)
	}
}

func Benchmark_SameWrapSlow_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		SameWrapSlow(in0, in1)
	}
}

func Benchmark_SameWrapSlow_PI(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, imgW, imgH))
	in1 := imutil.ToMat(imutil.Rand(359287343422, patchW, patchH))

	for i := 0; i < b.N; i++ {
		SameWrapSlow(in1, in0)
	}
}
