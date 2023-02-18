# Goconv2 - Two-dimensional convolutions in Go

This implements various 2d convolutions in Go, useful for signal processing and computer vision.

[Gonum's](https://www.gonum.org/) matrices are used as data types.

All implementations are tested against the output of the equivalent implementations in [SciPy](https://scipy.org/), assuming that those are probably correct.
The documentation is also conveniently outsourced to SciPy, links are included in godoc.

There are different implementations of the same functions:

- Naive Go - more or less copied 1:1 from the textbook. Those are very slow, because the Go compiler does not optimize a lot.
- Slightly optimized Go - a little faster. Only two most useful variants (to the author) are implemented.
- Cgo - much faster, even though they look like the naive implementations. The reason is that C compilers are much better at optimizing. Only two most useful variants (to the author) are implemented.
- An implementation using FFTs. This is only worth it for large images/matrices. Only the most useful variant (to the author) is implemented.

To choose the right implementation for you use case, you can tweak the image sizes used in the benchmark in `pkg/conv2/testutils.go`, and then run `make bench`.
In most of the case, the Cgo implementations seem to be the most useful ones.

For usage examples, see the tests in `pkg/conv2`.

## TODOs
- [ ] Tweak CPU optimizations (currently they are targeted to Raspberry Pis and seem to not always work).
- [ ] Implement more optimized versions.
- [ ] Optimize: do 2 FFTs in one.
- [ ] Test panics.
- [ ] Make it possible to configure what FFT is used.
