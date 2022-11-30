#!/usr/bin/env python3

"""
Generates ground truth images for the Golang tests,
assuming that numpy/scipy is trustworthy.

    pip install opencv-python-headless numpy
"""

import cv2
import numpy as np
from scipy import signal


def write_float_mat_to_png(path, img):
    "Dump float image"
    img = (img * 255).astype(np.uint8)
    cv2.imwrite(path, img)


def write_float_mat_to_png_normalized(path, img):
    "Dump float image, normalize first"
    max = np.max(img)
    min = np.min(img)

    assert max != min, "cannot normalize"

    img = ((img - min) / (max - min) * 255).astype(np.uint8)
    cv2.imwrite(path, img)


def main():
    # a bit verbose, but be explicit about the operations
    img = cv2.imread("img.png", cv2.IMREAD_COLOR)
    img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
    img = img.astype(np.float64) / 255.0

    x, y, w, h = 60, 70, 100, 110
    patch = img[y : y + h, x : x + w]
    write_float_mat_to_png("gen/patch.png", patch)

    write_float_mat_to_png_normalized(
        "gen/conv-full-fill.png", signal.convolve2d(img, patch, mode="full", boundary="fill", fillvalue=0)
    )
    write_float_mat_to_png_normalized(
        "gen/conv-full-wrap.png", signal.convolve2d(img, patch, mode="full", boundary="wrap", fillvalue=0)
    )

    write_float_mat_to_png_normalized(
        "gen/conv-valid-fill.png", signal.convolve2d(img, patch, mode="valid", boundary="fill", fillvalue=0)
    )
    write_float_mat_to_png_normalized(
        "gen/conv-valid-wrap.png", signal.convolve2d(img, patch, mode="valid", boundary="wrap", fillvalue=0)
    )

    signal.correlate2d

    write_float_mat_to_png_normalized(
        "gen/conv-same-fill.png", signal.convolve2d(img, patch, mode="same", boundary="fill", fillvalue=0)
    )
    write_float_mat_to_png_normalized(
        "gen/conv-same-wrap.png", signal.convolve2d(img, patch, mode="same", boundary="wrap", fillvalue=0)
    )


if __name__ == "__main__":
    main()
