#!/usr/bin/env python3

"""
Generates ground truth images for the Golang tests,
so that we can test our implementation against what numpy/scipy produces.

    pip install --upgrade opencv-python-headless scipy numpy
"""

import cv2
import numpy as np
from scipy import signal


def dump_mat(path, mat):
    "Dump a 2-dimensional float64 array, as txt file and as PNG for visualization."
    assert mat.dtype == np.float64
    assert mat.ndim == 2
    np.savetxt(path + ".txt", mat)

    max = np.max(mat)
    min = np.min(mat)

    if max == min:
        mat = np.zeros_like(mat)
    else:
        mat = ((mat - min) / (max - min) * 255).astype(np.uint8)

    cv2.imwrite(path + ".png", mat)


def main():
    # a bit verbose, but be explicit about the operations
    img = cv2.imread("img.png", cv2.IMREAD_COLOR)
    img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
    img = img.astype(np.float64) / 255.0
    dump_mat("gen/img", img)

    x, y, w, h = 3, 6, 8, 9
    patch = img[y : y + h, x : x + w]
    dump_mat("gen/patch", patch)

    dump_mat("gen/conv-ip-full-fill", signal.convolve2d(img, patch, mode="full", boundary="fill", fillvalue=0))
    dump_mat("gen/conv-ip-full-wrap", signal.convolve2d(img, patch, mode="full", boundary="wrap", fillvalue=0))
    dump_mat("gen/conv-pi-full-fill", signal.convolve2d(patch, img, mode="full", boundary="fill", fillvalue=0))
    dump_mat("gen/conv-pi-full-wrap", signal.convolve2d(patch, img, mode="full", boundary="wrap", fillvalue=0))
    dump_mat("gen/conv-ii-full-fill", signal.convolve2d(img, img, mode="full", boundary="fill", fillvalue=0))
    dump_mat("gen/conv-ii-full-wrap", signal.convolve2d(img, img, mode="full", boundary="wrap", fillvalue=0))
    dump_mat("gen/conv-pp-full-fill", signal.convolve2d(patch, patch, mode="full", boundary="fill", fillvalue=0))
    dump_mat("gen/conv-pp-full-wrap", signal.convolve2d(patch, patch, mode="full", boundary="wrap", fillvalue=0))

    dump_mat("gen/conv-ip-valid-fill", signal.convolve2d(img, patch, mode="valid", boundary="fill", fillvalue=0))
    dump_mat("gen/conv-ip-valid-wrap", signal.convolve2d(img, patch, mode="valid", boundary="wrap", fillvalue=0))
    dump_mat("gen/conv-pi-valid-fill", signal.convolve2d(patch, img, mode="valid", boundary="fill", fillvalue=0))
    dump_mat("gen/conv-pi-valid-wrap", signal.convolve2d(patch, img, mode="valid", boundary="wrap", fillvalue=0))
    dump_mat("gen/conv-ii-valid-fill", signal.convolve2d(img, img, mode="valid", boundary="fill", fillvalue=0))
    dump_mat("gen/conv-ii-valid-wrap", signal.convolve2d(img, img, mode="valid", boundary="wrap", fillvalue=0))
    dump_mat("gen/conv-pp-valid-fill", signal.convolve2d(patch, patch, mode="valid", boundary="fill", fillvalue=0))
    dump_mat("gen/conv-pp-valid-wrap", signal.convolve2d(patch, patch, mode="valid", boundary="wrap", fillvalue=0))

    dump_mat("gen/conv-ip-same-fill", signal.convolve2d(img, patch, mode="same", boundary="fill", fillvalue=0))
    dump_mat("gen/conv-ip-same-wrap", signal.convolve2d(img, patch, mode="same", boundary="wrap", fillvalue=0))
    dump_mat("gen/conv-pi-same-fill", signal.convolve2d(patch, img, mode="same", boundary="fill", fillvalue=0))
    dump_mat("gen/conv-pi-same-wrap", signal.convolve2d(patch, img, mode="same", boundary="wrap", fillvalue=0))
    dump_mat("gen/conv-ii-same-fill", signal.convolve2d(img, img, mode="same", boundary="fill", fillvalue=0))
    dump_mat("gen/conv-ii-same-wrap", signal.convolve2d(img, img, mode="same", boundary="wrap", fillvalue=0))
    dump_mat("gen/conv-pp-same-fill", signal.convolve2d(patch, patch, mode="same", boundary="fill", fillvalue=0))
    dump_mat("gen/conv-pp-same-wrap", signal.convolve2d(patch, patch, mode="same", boundary="wrap", fillvalue=0))


if __name__ == "__main__":
    main()
