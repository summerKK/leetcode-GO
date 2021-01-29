package lib

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

var S *FixedCapStack
var MS *MyStack

func TestMain(m *testing.M) {
	S = &FixedCapStack{}
	S.Init(1)
	MS = &MyStack{}
	MS.Init(0)
	m.Run()
}

func TestFixedCapStack_Push(t *testing.T) {
	input := "to be or not to - be - - that - - - is"
	inputs := strings.Split(input, " ")
	for _, input := range inputs {
		if input == "-" && !S.IsEmpty() {
			fmt.Println(S.Pop())
		} else {
			S.Push(input)
		}
	}
}

func TestFixedCapStack_Loop(t *testing.T) {
	input := "to be or not to be that is"
	inputs := strings.Split(input, " ")
	for _, s := range inputs {
		S.Push(s)
	}

	for v := range S.Loop() {
		fmt.Printf("%v ", v)
	}
}

func TestMyStack_Push(t *testing.T) {
	input := "to be or not to - be - - that - - - is"
	inputs := strings.Split(input, " ")
	for _, input := range inputs {
		if input == "-" && !MS.IsEmpty() {
			fmt.Println(MS.Pop())
		} else {
			MS.Push(input)
		}
	}
}

func TestMyStack_Loop(t *testing.T) {
	input := "to be or not to be that is"
	inputs := strings.Split(input, " ")
	for _, s := range inputs {
		MS.Push(s)
	}

	for v := range MS.Loop() {
		fmt.Printf("%v ", v)
	}
}

func TestMyStack_Peek(t *testing.T) {
	input := "to be or not to be that is"
	inputs := strings.Split(input, " ")
	for _, s := range inputs {
		MS.Push(s)
	}

	assert.Equal(t, MS.Peek(), "is")
}

func TestMyStack_Copy(t *testing.T) {
	input := "to be or not to be that is"
	inputs := strings.Split(input, " ")
	for _, s := range inputs {
		MS.Push(s)
	}

	stack0 := &MyStack{}
	stack0.Init(0)
	MS.Copy(stack0)

	for s := range stack0.Loop() {
		fmt.Printf("%s ", s)
	}
	fmt.Println()
}
