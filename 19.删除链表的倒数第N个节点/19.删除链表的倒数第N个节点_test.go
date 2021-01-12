package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testcases := []struct {
		arg0   *ListNode
		arg1   int
		except []int
	}{
		{
			arg0: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val:  6,
									Next: nil,
								},
							},
						},
					},
				},
			},
			arg1:   2,
			except: []int{1, 2, 3, 4, 6},
		},
	}

	traverse := func(head *ListNode) []int {
		var result []int
		for head != nil {
			result = append(result, head.Val)
			head = head.Next
		}

		return result
	}

	for _, testcase := range testcases {
		head := removeNthFromEnd(testcase.arg0, testcase.arg1)
		assert.Equal(t, traverse(head), testcase.except)
	}
}
