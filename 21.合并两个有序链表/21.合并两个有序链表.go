package leetcode

import . "github.com/summerKK/leetcode-Go/utils"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	result := preHead

	// 至少保证一个链表遍历完
	for l1 != nil && l2 != nil {
		// 按从小到大的顺序排列
		if l1.Val > l2.Val {
			preHead.Next = l2
			// 指针下移
			l2 = l2.Next
		} else {
			preHead.Next = l1
			l1 = l1.Next
		}

		// 指针下移
		preHead = preHead.Next
	}

	// 两个链表长度不一致 如果l1还是l2中任意一方还有余下元素没有用到，那余下的这些元素一定大于prehead已经合并完的链表（因为是有序链表）。我们只需要将这些元素全部追加到prehead合并完的链表后
	if l1 != nil {
		preHead.Next = l1
	}

	if l2 != nil {
		preHead.Next = l2
	}

	return result.Next
}
