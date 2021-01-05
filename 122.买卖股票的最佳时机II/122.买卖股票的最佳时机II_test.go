package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMaxProfit(t *testing.T) {
	data := []map[string]interface{}{
		{
			"test":   []int{7, 6, 5, 4, 3, 2, 1},
			"except": 0,
		},
		{
			"test":   []int{7, 1, 5, 3, 6, 4},
			"except": 7,
		},
		{
			"test":   []int{1, 2, 3, 4, 5},
			"except": 4,
		},
	}

	for _, v := range data {
		d := v["test"].([]int)
		profit := maxProfit(d)
		assert.Equal(t, profit, v["except"])
	}
}
