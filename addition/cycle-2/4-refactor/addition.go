package main

import (
    "errors"
    "math"
)

func Add(a, b int) (int, error) {
    if b > 0 {
        if a > math.MaxInt64 - b {
            return -1, errors.New("Overflow")
        }
    }

    return a + b, nil
}
