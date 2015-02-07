#!/bin/env python2
# -*- coding:utf-8 -*-
import cv2

import sys


SKIP_FRAME = 10


if __name__ == "__main__":
    frame = None
    try:
        cam0 = cv2.VideoCapture(0)
        # cam1 = cv2.VideoCapture(1)

        # 何フレームか飛ばしてから本撮影する
        for i in range(SKIP_FRAME):
            ret, frame = cam0.read()

    except:
        print "Failed to capture image"

    finally:
        cam0.release()

    cv2.imwrite("tmp.jpg", frame) # save
