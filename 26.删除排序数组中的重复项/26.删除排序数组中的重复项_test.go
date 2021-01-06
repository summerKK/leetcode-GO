package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func testRemoveDuplicates(t *testing.T, process func(nums []int) int) {
	data := []map[string]interface{}{
		{
			"arg0":    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			"except0": 5,
			"except1": []int{0, 1, 2, 3, 4},
		},
		{
			"arg0":    []int{1, 1, 2},
			"except0": 2,
			"except1": []int{1, 2},
		},
	}

	for _, v := range data {
		arg0 := v["arg0"].([]int)
		except0 := v["except0"].(int)
		duplicates := process(arg0)
		assert.Equal(t, duplicates, v["except0"])
		assert.Equal(t, arg0[:except0], v["except1"])
	}
}

func TestRemoveDuplicates0(t *testing.T) {
	testRemoveDuplicates(t, removeDuplicates0)
}

func TestRemoveDuplicates1(t *testing.T) {
	testRemoveDuplicates(t, removeDuplicates1)
}

func TestRemoveDuplicates2(t *testing.T) {
	testRemoveDuplicates(t, removeDuplicates2)
}
