package lib

import (
	"fmt"
	"strings"
	"testing"
)

var S *FixedCapStack

func TestMain(m *testing.M) {
	S = &FixedCapStack{}
	S.Init(1)
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
