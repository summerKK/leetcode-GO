package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMinPathSum(t *testing.T) {
	testcases := []struct {
		arg0   [][]int
		except int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			12,
		},
		{
			[][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			7,
		},
	}

	for _, testcase := range testcases {
		result := minPathSum(testcase.arg0)
		assert.Equal(t, result, testcase.except)
	}
}
