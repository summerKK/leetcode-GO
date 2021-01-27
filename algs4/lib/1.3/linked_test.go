package lib

import (
	"fmt"
	"testing"
)

func TestLinked_Add(t *testing.T) {
	L := &Linked{}
	L.Init()
	collection := []string{"to", "be", "or", "not", "to", "be", "that", "is"}
	for _, s := range collection {
		L.Add(s)
	}
	L.Del()
	for node := L.current; node.next != nil; node = node.Next() {
		fmt.Printf("%v ", node.Val())
	}
}
