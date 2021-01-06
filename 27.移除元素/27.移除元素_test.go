package leetcode

import (
	"sort"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRemoveElement(t *testing.T) {
	data := []map[string]interface{}{
		{
			"arg0":    []int{3, 2, 2, 3},
			"arg1":    3,
			"except0": 2,
			"except1": []int{2, 2},
		},
		{
			"arg0":    []int{0, 1, 2, 2, 3, 0, 4, 2},
			"arg1":    2,
			"except0": 5,
			"except1": []int{0, 1, 3, 0, 4},
		},
	}

	for _, v := range data {
		arg0 := v["arg0"].([]int)
		arg1 := v["arg1"].(int)
		except0 := v["except0"].(int)
		except1 := v["except1"].([]int)
		n := removeElement(arg0, arg1)
		assert.Equal(t, n, except0)
		sort.Ints(arg0[:except0])
		sort.Ints(except1)
		assert.Equal(t, arg0[:except0], except1)
	}
}
