package conv2

import (
	"testing"

	"github.com/jo-m/goconv2/internal/pkg/imutil"
	"github.com/jo-m/goconv2/internal/pkg/testutil"
)

func Test_FullFill(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := FullFill(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFill(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFill(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullFill(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-full-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_FullFill_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 20, 20))
	in1 := imutil.ToMat(imutil.Rand(129038, 20, 20))

	for i := 0; i < b.N; i++ {
		FFull(in0, in1)
	}
}

func Benchmark_FullFill_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(129038, 160, 120))

	for i := 0; i < b.N; i++ {
		FullFill(in0, in1)
	}
}

func Benchmark_FullFill_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		FullFill(in0, in1)
	}
}

func Benchmark_FullFill_PI(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		FullFill(in1, in0)
	}
}

func Test_FullWrap(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := FullWrap(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-full-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullWrap(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-full-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullWrap(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-full-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = FullWrap(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-full-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_FullWrap_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 20, 20))
	in1 := imutil.ToMat(imutil.Rand(129038, 20, 20))

	for i := 0; i < b.N; i++ {
		FFull(in0, in1)
	}
}

func Benchmark_FullWrap_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(129038, 160, 120))

	for i := 0; i < b.N; i++ {
		FullWrap(in0, in1)
	}
}

func Benchmark_FullWrap_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		FullWrap(in0, in1)
	}
}

func Benchmark_FullWrap_PI(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		FullWrap(in1, in0)
	}
}

func Test_ValidFill(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := ValidFill(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidFill(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidFill(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidFill(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-valid-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_ValidFill_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 20, 20))
	in1 := imutil.ToMat(imutil.Rand(129038, 20, 20))

	for i := 0; i < b.N; i++ {
		FFull(in0, in1)
	}
}

func Benchmark_ValidFill_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(129038, 160, 120))

	for i := 0; i < b.N; i++ {
		ValidFill(in0, in1)
	}
}

func Benchmark_ValidFill_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		ValidFill(in0, in1)
	}
}

func Benchmark_ValidFill_PI(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		ValidFill(in1, in0)
	}
}

func Test_ValidWrap(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := ValidWrap(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-valid-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidWrap(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-valid-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidWrap(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-valid-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = ValidWrap(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-valid-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_ValidWrap_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 20, 20))
	in1 := imutil.ToMat(imutil.Rand(129038, 20, 20))

	for i := 0; i < b.N; i++ {
		FFull(in0, in1)
	}
}

func Benchmark_ValidWrap_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(129038, 160, 120))

	for i := 0; i < b.N; i++ {
		ValidWrap(in0, in1)
	}
}

func Benchmark_ValidWrap_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		ValidWrap(in0, in1)
	}
}

func Benchmark_ValidWrap_PI(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		ValidWrap(in1, in0)
	}
}

func Test_SameFill(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := SameFill(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-same-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = SameFill(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-same-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = SameFill(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-same-fill.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = SameFill(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-same-fill.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_SameFill_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 20, 20))
	in1 := imutil.ToMat(imutil.Rand(129038, 20, 20))

	for i := 0; i < b.N; i++ {
		FFull(in0, in1)
	}
}

func Benchmark_SameFill_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(129038, 160, 120))

	for i := 0; i < b.N; i++ {
		SameFill(in0, in1)
	}
}

func Benchmark_SameFill_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		SameFill(in0, in1)
	}
}

func Benchmark_SameFill_PI(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		SameFill(in1, in0)
	}
}

func Test_SameWrap(t *testing.T) {
	img, patch := loadImgAndPatch(t)

	out := SameWrap(patch, patch)
	truth := testutil.LoadMat64Txt(t, "testdata/gen/conv-pp-same-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = SameWrap(img, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ii-same-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = SameWrap(img, patch)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-ip-same-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)

	out = SameWrap(patch, img)
	truth = testutil.LoadMat64Txt(t, "testdata/gen/conv-pi-same-wrap.txt")
	testutil.AssertMatEqual(t, truth, out)
}

func Benchmark_SameWrap_PP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 20, 20))
	in1 := imutil.ToMat(imutil.Rand(129038, 20, 20))

	for i := 0; i < b.N; i++ {
		FFull(in0, in1)
	}
}

func Benchmark_SameWrap_II(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(531535, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(129038, 160, 120))

	for i := 0; i < b.N; i++ {
		SameWrap(in0, in1)
	}
}

func Benchmark_SameWrap_IP(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		SameWrap(in0, in1)
	}
}

func Benchmark_SameWrap_PI(b *testing.B) {
	in0 := imutil.ToMat(imutil.Rand(340592732523, 160, 120))
	in1 := imutil.ToMat(imutil.Rand(359287343422, 20, 20))

	for i := 0; i < b.N; i++ {
		SameWrap(in1, in0)
	}
}
