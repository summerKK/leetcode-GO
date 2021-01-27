package main

import (
	"fmt"
	"strings"

	lib "github.com/summerKK/leetcode-Go/algs4/lib/1.3"
)

/**
1.3.4
编写一个 Stack 的用例 Parentheses，从标准输入中读取一个文本流并使用栈判定其中的括 号是否配对完整。
例如，对于 [()]{}{[()()]()} 程序应该打印 true，对于 [(]) 则打印 false。
*/
func main() {
	parentheses("[ ( ) ] { } { [ ( ) ( ) ] ( ) }")
	parentheses("[ ( ] )")
	parentheses("[ ( ) ]")
}

func parentheses(parentheses string) {
	collection := strings.Split(parentheses, " ")
	S := &lib.MyStack{}
	S.Init(0)
	ok := true
	for _, s := range collection {
		if s == "[" || s == "(" || s == "{" {
			S.Push(s)
		} else {
			switch s {
			case ")":
				ok = S.Pop() == "("
			case "]":
				ok = S.Pop() == "["
			case "}":
				ok = S.Pop() == "{"
			}
			if !ok {
				break
			}
		}
	}

	if ok {
		fmt.Printf("parentheses:%s %v\n", parentheses, true)
	} else {
		fmt.Printf("parentheses:%s %v\n", parentheses, false)
	}
}
