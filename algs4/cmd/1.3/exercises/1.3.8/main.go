package main

import (
	"fmt"
	"strings"

	"github.com/summerKK/leetcode-Go/algs4/cmd/1.3"
)

/**
1.3.8
给定以下输入，给出 DoublingStackOfStrings 的数组的内容和大小。
   it was - the best - of times - - - it was - the - -
*/
func main() {
	input := "it was - the best - of times - - - it was - the - -"
	inputs := strings.Split(input, " ")
	MS := &__3.MyStack{}
	MS.Init(0)
	for _, input := range inputs {
		if input == "-" && !MS.IsEmpty() {
			fmt.Println(MS.Pop())
		} else {
			MS.Push(input)
		}
	}

	for s := range MS.Loop() {
		fmt.Printf("s: %v,len:%d\n", s, MS.Size())
	}
}
