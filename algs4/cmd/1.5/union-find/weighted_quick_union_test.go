package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestWeightedQuickUnion_Union(t *testing.T) {
	qu := NewWeightedQuickUnion(10)

	qu.Union(4, 3)
	qu.Union(3, 8)
	qu.Union(6, 5)
	qu.Union(9, 4)
	qu.Union(2, 1)
	qu.Union(8, 9)
	qu.Union(5, 0)
	qu.Union(7, 2)
	qu.Union(6, 1)
	qu.Union(1, 0)
	qu.Union(6, 7)

	assert.Equal(t, 2, qu.Count())
}
