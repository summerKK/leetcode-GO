package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
	. "github.com/summerKK/leetcode-Go/utils"
)

func TestAddTwoNumbers(t *testing.T) {
	testcases := []struct {
		arg0   *ListNode
		arg1   *ListNode
		except *ListNode
	}{
		{
			arg0:   GenLinked([]int{9, 9, 9, 9, 9, 9, 9}),
			arg1:   GenLinked([]int{9, 9, 9, 9}),
			except: GenLinked([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			arg0:   GenLinked([]int{2, 4, 3}),
			arg1:   GenLinked([]int{5, 6, 4}),
			except: GenLinked([]int{7, 0, 8}),
		},
	}

	for _, testcase := range testcases {
		result := addTwoNumbers(testcase.arg0, testcase.arg1)
		assert.Equal(t, Traverse(result), Traverse(testcase.except))
	}
}
