package leetcode

import (
	"fmt"
	"sort"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestTwoSum(t *testing.T) {
	data := []map[string]interface{}{
		{
			"arg0":   []int{3, 3, 3, 3, 3, 3},
			"arg1":   6,
			"except": []int{0, 1},
		},
		{
			"arg0":   []int{2, 7, 11, 15},
			"arg1":   9,
			"except": []int{0, 1},
		},
		{
			"arg0":   []int{3, 2, 4},
			"arg1":   6,
			"except": []int{1, 2},
		},
	}

	for _, v := range data {
		arg0 := v["arg0"].([]int)
		arg1 := v["arg1"].(int)
		except := v["except"].([]int)
		result := twoSum(arg0, arg1)
		fmt.Println(arg0)
		sort.Ints(result)
		sort.Ints(except)
		assert.Equal(t, result, except)
	}
}
