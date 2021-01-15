package leetcode

import . "github.com/summerKK/leetcode-Go/utils"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{}
	head := list
	tmp := 0
	// tmp != 0的条件很重.就像测试用例中的.如果没有这个tmp != 0 的条件最后一次tmp的值是1(进位)没办法进入for中
	// 最后就会导致缺少一位
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}

		// 这里list的第一个元素的head.head.next才是链表的第一个元素
		// 每个位上面的数 <= 9 ,所以 mod 10.取个位.如果需要进位存在tmp变量里面.下次相加的时候会自动加上
		list.Next = &ListNode{Val: tmp % 10}
		list = list.Next
		tmp = tmp / 10
	}

	return head.Next
}
