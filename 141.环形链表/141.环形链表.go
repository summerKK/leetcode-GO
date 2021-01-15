package leetcode

import . "github.com/summerKK/leetcode-Go/utils"

func hasCycle(head *ListNode) bool {
	// 慢指针
	slow := head
	// 快指针
	fast := head
	for fast != nil {
		// 快指针移动两次
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		// 当快指针(第二圈)追赶上慢指针.代表链表内有环(如果没有环不会再相遇)
		if fast == slow {
			return true
		}
		// 慢指针移动一次
		slow = slow.Next
	}

	return false
}
