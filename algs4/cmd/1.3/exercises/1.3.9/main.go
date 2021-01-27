package main

import (
	"fmt"
	"strconv"
	"strings"

	lib "github.com/summerKK/leetcode-Go/algs4/lib/1.3"
)

/**
1.3.9
编写一段程序，从标准输入得到一个缺少左括号的表达式并打印出补全括号之后的中序表达式。 例如，给定输入:
	1+2)* 3 - 4 ) * 5 - 6 ) ) )
	你的程序应该输出:
	( ( 1 + 2 ) * ( ( 3 - 4 ) * ( 5 - 6 )))
*/
func main() {
	supplement("1 + 2 ) * 3 - 4 ) * 5 - 6 ) ) )")
}

func supplement(s string) {
	collection := strings.Split(s, " ")
	numbers := &lib.MyStack{}
	numbers.Init(0)
	opts := &lib.MyStack{}
	opts.Init(0)
	dataStack := lib.MyStack{}
	dataStack.Init(0)

	for _, s := range collection {
		if isDigital(s) {
			numbers.Push(s)
		} else if isOperator(s) {
			opts.Push(s)
		} else {
			// 括号
			if numbers.Size() >= 2 {
				n0 := numbers.Pop()
				n1 := numbers.Pop()
				opt := opts.Pop()
				dataStack.Push(fmt.Sprintf("( %s %s %s )", n1, opt, n0))
			}
		}
	}

	for opt := opts.Pop(); opt != nil; opt = opts.Pop() {
		data0 := dataStack.Pop()
		data1 := dataStack.Pop()
		dataStack.Push(fmt.Sprintf("( %s %s %s )", data1, opt, data0))
	}

	fmt.Println(dataStack.Pop())
}

func isDigital(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}

	return true
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}
