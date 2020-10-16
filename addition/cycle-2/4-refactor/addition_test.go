package main

import (
    "math"
    "errors"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    result, _ := Add(1, 2)
    assert.Equal(t, result, 3)
}

func TestAdd_UpperBound(t *testing.T) {
    _, err := Add(math.MaxInt64, 1)
    assert.Equal(t, errors.New("Overflow"), err)
}
