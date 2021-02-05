package main

import (
	"math"

	"github.com/summerKK/leetcode-Go/algs4/lib"
	"github.com/summerKK/leetcode-Go/algs4/utils"
)

type evaluate struct {
	ops  *lib.Stack
	vals *lib.Stack
}

func NewEvaluate() *evaluate {
	return &evaluate{
		ops:  lib.NewStack(),
		vals: lib.NewStack(),
	}
}

func (e *evaluate) Calc(exprs []string) float64 {
	for _, expr := range exprs {
		switch expr {
		case "(":
		case "+":
			fallthrough
		case "-":
			fallthrough
		case "*":
			fallthrough
		case "/":
			fallthrough
		case "sqrt":
			e.ops.Push(expr)
		case ")":
			op := e.ops.Pop().(string)
			v := e.vals.Pop().(float64)
			switch op {
			case "+":
				v = e.vals.Pop().(float64) + v
			case "-":
				v = e.vals.Pop().(float64) - v
			case "*":
				v = e.vals.Pop().(float64) * v
			case "/":
				v = e.vals.Pop().(float64) / v
			case "sqrt":
				v = math.Sqrt(v)
			}
			e.vals.Push(v)
		default:
			e.vals.Push(utils.StrTo(expr).MustFloat64())
		}
	}

	return e.vals.Pop().(float64)
}
