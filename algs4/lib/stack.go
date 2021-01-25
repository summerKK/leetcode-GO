package lib

import "github.com/summerKK/leetcode-Go/algs4/utils"

// Stack implements a stack by utils
type Stack struct {
	*utils.LinkList
}

// NewStack ...
func NewStack() *Stack {
	return &Stack{utils.NewLinkList()}
}

// Push ...
func (s *Stack) Push(item interface{}) {
	s.Add(item)
}

// Pop ...
func (s *Stack) Pop() (item interface{}) {
	if s.IsEmpty() {
		panic("Stack Empty")
	}
	item = s.Del()
	return
}
