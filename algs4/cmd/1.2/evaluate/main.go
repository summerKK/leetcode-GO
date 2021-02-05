package main

import (
	"fmt"
	"strings"
)

var e = NewEvaluate()

// Dijkstra 的双栈算术表达式求值算法
func main() {
	run("( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )")
	run("( ( 1 + sqrt ( 5.0 ) ) / 2.0 )")
}

func run(s string) {
	exprs := strings.Split(s, " ")
	result := e.Calc(exprs)
	fmt.Printf("express: %s, result: %f\n", s, result)
}
