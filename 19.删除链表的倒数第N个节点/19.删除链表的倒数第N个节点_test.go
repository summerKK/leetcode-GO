package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
	. "github.com/summerKK/leetcode-Go/utils"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testcases := []struct {
		arg0   *ListNode
		arg1   int
		except []int
	}{
		{
			arg0:   GenLinked([]int{1, 2, 3, 4, 5, 6}),
			arg1:   2,
			except: []int{1, 2, 3, 4, 6},
		},
	}

	for _, testcase := range testcases {
		head := removeNthFromEnd(testcase.arg0, testcase.arg1)
		assert.Equal(t, Traverse(head), testcase.except)
	}
}
