package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/summerKK/leetcode-Go/algs4/cmd/1.3"
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
	numbers := &__3.MyStack{}
	numbers.Init(0)
	opts := &__3.MyStack{}
	opts.Init(0)
	dataStack := __3.MyStack{}
	dataStack.Init(0)

	for _, s := range collection {
		if isDigital(s) {
			numbers.Push(s)
		} else if isOperator(s) {
			opts.Push(s)
		} else {
			// 括号
			n0 := numbers.Pop()
			n1 := numbers.Pop()
			opt := opts.Pop()
			numbers.Push(fmt.Sprintf("( %s %s %s )", n1, opt, n0))
		}
	}

	fmt.Println(numbers.Pop())
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
