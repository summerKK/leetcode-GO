package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIntersect0(t *testing.T) {
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2, 2, 3}
	except := []int{2, 2}

	result := intersect0(num1, num2)

	assert.Equal(t, result, except)
}
