package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRotate(t *testing.T) {
	data := []map[string]interface{}{
		{
			"arg0":   []int{1, 2},
			"arg1":   3,
			"except": []int{2, 1},
		},
		{
			"arg0":   []int{1, 2, 3, 4, 5, 6, 7},
			"arg1":   3,
			"except": []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"arg0":   []int{-1, -100, 3, 99},
			"arg1":   2,
			"except": []int{3, 99, -1, -100},
		},
		{
			"arg0":   []int{-1},
			"arg1":   2,
			"except": []int{-1},
		},
	}

	for _, v := range data {
		arg0 := v["arg0"].([]int)
		arg1 := v["arg1"].(int)
		rotate(arg0, arg1)
		assert.Equal(t, arg0, v["except"])
	}
}
