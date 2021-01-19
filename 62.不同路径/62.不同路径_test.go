package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUniquePaths(t *testing.T) {
	testcases := []struct {
		arg0   int
		arg1   int
		except int
	}{
		{3, 7, 28},
		{3, 2, 3},
		{3, 3, 6},
	}
	for _, testcase := range testcases {
		result := uniquePaths(testcase.arg0, testcase.arg1)
		assert.Equal(t, result, testcase.except)
	}
}
