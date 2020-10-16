package main

import (
    "errors"
    "math"
)

func Add(a, b int) (int, error) {
    if a == math.MaxInt64 {
        return -1, errors.New("Overflow")
    }

    return a + b, nil
}
