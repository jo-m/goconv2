#include "cconv.h"

void FullFillC(int dy, int dx, int dv, int du, int n, int m, int fs, int gs, int outs, float64 *fd, float64 *gd,
               float64 *outd) {
  for (int y = 0; y < m; y++) {
    for (int x = 0; x < n; x++) {
      float64 sum = 0;

      for (int v = 0; v < dv; v++) {
        for (int u = 0; u < du; u++) {
          if (y - v < 0 || y - v >= dy || x - u < 0 || x - u >= dx) {
            continue;  // zero
          }

          sum += fd[(y - v) * fs + (x - u)] * gd[v * gs + u];
        }
      }

      outd[y * outs + x] = sum;
    }
  }
}

void ValidFillC(int dy, int dx, int dv, int du, int n, int m, int fs, int gs, int outs, float64 *fd, float64 *gd,
                float64 *outd) {
  for (int y = 0; y < m; y++) {
    for (int x = 0; x < n; x++) {
      float64 sum = 0;

      for (int v = 0; v < dv; v++) {
        for (int u = 0; u < du; u++) {
          sum += gd[v * gs + u] * fd[(y + dv - 1 - v) * fs + (x + du - 1 - u)];
        }
      }

      outd[y * outs + x] = sum;
    }
  }
}
