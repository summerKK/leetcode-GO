package lib

import (
	"fmt"
	"strings"
	"testing"
)

var S Stack

func TestMain(m *testing.M) {
	S = &FixedCapStack{}
	S.Init(1)
	m.Run()
}

func TestPush(t *testing.T) {
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
