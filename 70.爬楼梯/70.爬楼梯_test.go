package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestClimbStairs(t *testing.T) {
	testcases := []struct {
		arg0   int
		except int
	}{
		{
			1,
			1,
		},
		{
			2,
			2,
		},
		{
			3,
			3,
		},
		{
			4,
			5,
		},
	}
	for _, testcase := range testcases {
		result := climbStairs(testcase.arg0)
		assert.Equal(t, result, testcase.except)
	}
}
