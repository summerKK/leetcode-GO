package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPlusOne(t *testing.T) {
	data := []map[string]interface{}{
		{
			"arg0":   []int{1, 2, 3},
			"except": []int{1, 2, 4},
		},
		{
			"arg0":   []int{9, 9, 9},
			"except": []int{1, 0, 0, 0},
		},
	}

	for _, v := range data {
		nums := v["arg0"].([]int)
		newDigits := plusOne(nums)
		assert.Equal(t, newDigits, v["except"])
	}
}
