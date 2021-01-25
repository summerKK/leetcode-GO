package lib

import "github.com/summerKK/leetcode-Go/algs4/utils"

// Bag implements a bag data type by utils
type Bag struct {
	*utils.LinkList
}

// NewBag ...
func NewBag() *Bag {
	return &Bag{
		utils.NewLinkList(),
	}
}
