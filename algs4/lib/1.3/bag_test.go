package lib

import (
	"fmt"
	"strings"
	"testing"
)

func TestMyBag_Add(t *testing.T) {
	B := &MyBag{}
	B.Init()

	input := "to be or not to be that is"
	inputs := strings.Split(input, " ")
	for _, s := range inputs {
		B.Add(s)
	}

	fmt.Print(B)
}
