package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 哨兵节点,用于返回链表的头
	result := &ListNode{}
	// 哨兵节点的next节点指向链表的头
	result.Next = head
	var pre *ListNode
	cur := result
	i := 1
	// 当head == nil的时候代表链表已经遍历完.距离head n 的距离就是需要删除的元素
	for head != nil {
		// 当 head和cur的距离为n的时候.cur和pre开始移动
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	// 删除导数第n个元素
	if pre == nil {
		return nil
	}

	pre.Next = pre.Next.Next

	return result.Next
}
