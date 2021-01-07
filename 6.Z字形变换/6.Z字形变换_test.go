package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestConvert(t *testing.T) {
	data := []map[string]interface{}{
		{
			"arg0":   "PAYPALISHIRING",
			"arg1":   4,
			"except": "PINALSIGYAHRPI",
		},
		{
			"arg0":   "LEETCODEISHIRING",
			"arg1":   3,
			"except": "LCIRETOESIIGEDHN",
		},
	}

	for _, v := range data {
		arg0 := v["arg0"].(string)
		arg1 := v["arg1"].(int)
		result := convert(arg0, arg1)
		assert.Equal(t, result, v["except"])
	}
}
