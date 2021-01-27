package lib

import (
	"fmt"
	"strings"
	"testing"
)

func TestMyQueue_Enqueue(t *testing.T) {
	Q := &MyQueue{}
	Q.Init()

	input := "to be or not to be that is"
	inputs := strings.Split(input, " ")
	for _, s := range inputs {
		Q.Enqueue(s)
	}

	fmt.Println(Q.Dequeue())
	fmt.Println(Q.Dequeue())
	fmt.Println(Q.Dequeue())
}
