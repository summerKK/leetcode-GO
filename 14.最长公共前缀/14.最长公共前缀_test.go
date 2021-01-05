package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLongestCommonPrefix(t *testing.T) {
	data := []map[string]interface{}{
		{
			"test":   []string{"flower", "flow", "flight"},
			"except": "fl",
		},
		{
			"test":   []string{"summer", "summ", "summmmmmmmmmm"},
			"except": "summ",
		},
	}

	for _, v := range data {
		s, _ := v["test"].([]string)
		prefix := longestCommonPrefix(s)
		assert.Equal(t, prefix, v["except"])
	}
}
