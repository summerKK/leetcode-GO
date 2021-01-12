package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenLinked(l []int) *ListNode {
	head := &ListNode{}
	result := head
	for _, v := range l {
		node := &ListNode{}
		node.Val = v
		head.Next = node
		head = head.Next
	}

	return result.Next
}

func Traverse(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}
