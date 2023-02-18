#pragma once

#define float64 double

void FullFillC(
    int dy,
    int dx,
    int dv,
    int du,
    int n,
    int m,
    int fs,
    int gs,
    int outs,
    float64 *fd,
    float64 *gd,
    float64 *outd
);

void ValidFillC(
    int dy,
    int dx,
    int dv,
    int du,
    int n,
    int m,
    int fs,
    int gs,
    int outs,
    float64 *fd,
    float64 *gd,
    float64 *outd
);
