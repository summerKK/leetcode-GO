package main

import (
	"fmt"
	"strconv"
	"strings"

	lib "github.com/summerKK/leetcode-Go/algs4/lib/1.3"
)

/**
1.3.10
编写一个过滤器 InfixToPostfix，将算术表达式由中序表达式转为后序表达式。
*/
func main() {
	//s1 := "( ( 1 + 2 ) * ( ( 3 - 4 ) * ( 5 - 6 ) ) )"
	//s1 := "( 3 + 2 ) * ( 2 + 3 )"
	s1 := "( ( 2 * 3 ) / ( 2 - 1 ) ) + ( 3 * ( 4 - 1 ) )"
	// s2 := "1 4 + 3 *  10 5 / +"
	collection := strings.Split(s1, " ")
	stack0 := &lib.MyStack{}
	stack0.Init(0)
	for _, s := range collection {
		if isDigital(s) {
			fmt.Printf("%s ", s)
		} else {
			if s == "(" {
				stack0.Push(s)
				continue
			}
			if s == ")" {
				for v := stack0.Pop(); v != "("; v = stack0.Pop() {
					fmt.Printf("%s ", v)
				}
				continue
			}
			v, ok := stack0.Peek().(string)
			if !ok {
				v = ""
			}
			if priority(s) >= priority(v) {
				stack0.Push(s)
			} else {
				fmt.Printf("%s ", s)
			}
		}
	}

	for s := range stack0.Loop() {
		fmt.Printf("%s ", s)
	}
}

func isDigital(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}

	return true
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/" || s == "(" || s == ")"
}

func priority(s string) int {
	if s == "+" || s == "-" {
		return 1
	}
	if s == "*" || s == "/" {
		return 2
	}

	return 0
}
