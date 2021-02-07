package lib

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUF_Union(t *testing.T) {
	uf := NewUF(10)

	uf.Union(4, 3)
	uf.Union(3, 8)
	uf.Union(6, 5)
	uf.Union(9, 4)
	uf.Union(2, 1)
	uf.Union(8, 9)
	uf.Union(5, 0)
	uf.Union(7, 2)
	uf.Union(6, 1)
	uf.Union(1, 0)
	uf.Union(6, 7)

	assert.Equal(t, 2, uf.Count())
}
