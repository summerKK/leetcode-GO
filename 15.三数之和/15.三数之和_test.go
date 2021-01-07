package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestThreeSum(t *testing.T) {
	data := []map[string]interface{}{
		{
			"arg0": []int{-1, 0, 1, 2, -1, -4},
		},
	}

	for _, v := range data {
		arg0 := v["arg0"].([]int)
		result := threeSum(arg0)
		for _, item0 := range result {
			r := 0
			for _, item1 := range item0 {
				r += item1
			}
			assert.Equal(t, r, 0)
			r = 0
		}
	}
}
