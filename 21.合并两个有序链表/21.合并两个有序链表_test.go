package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
	. "github.com/summerKK/leetcode-Go/utils"
)

func TestMergeTwoLists(t *testing.T) {

	testcases := []struct {
		arg0   *ListNode
		arg1   *ListNode
		except []int
	}{
		{
			arg0:   GenLinked([]int{1, 2, 4}),
			arg1:   GenLinked([]int{1, 3, 4}),
			except: []int{1, 1, 2, 3, 4, 4},
		},
		{
			arg0:   GenLinked([]int{1, 2, 4, 5, 6, 7, 8}),
			arg1:   GenLinked([]int{4, 5, 6}),
			except: []int{1, 2, 4, 4, 5, 5, 6, 6, 7, 8},
		},
	}

	for _, testcase := range testcases {
		head := mergeTwoLists(testcase.arg1, testcase.arg0)
		assert.Equal(t, Traverse(head), testcase.except)
	}
}
