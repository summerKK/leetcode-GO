package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)
import . "github.com/summerKK/leetcode-Go/utils"

func TestHasCycle(t *testing.T) {
	testcases := []struct {
		arg0   *ListNode
		except bool
	}{
		{
			(func() *ListNode {
				linked0 := GenLinked([]int{3, 2, 0, 4})
				linked1 := linked0
				for linked0 != nil {
					if linked0.Next == nil {
						linked0.Next = linked1
						break
					}
					// 这里 linked0的指针改变了.所以对linked0做操作不会影响到linked1
					linked0 = linked0.Next
				}

				return linked1
			})(),
			true,
		},
	}

	for _, testcase := range testcases {
		result := hasCycle(testcase.arg0)
		assert.Equal(t, result, testcase.except)
	}
}
