package main

import "fmt"

// 1.1.11 编写一段代码，打印出一个二维布尔数组的内容。其中，使用 * 表示真，空格表示假。打印出行号和列号。
func main() {
	arrs := [][]bool{{true, false, true, false}, {false, false, true, true}, {true, true, true, false}}
	fmt.Printf("%2s", "")
	for i := 0; i < len(arrs[0]); i++ {
		fmt.Printf("%2d", i)
	}
	fmt.Println()
	for i := 0; i < len(arrs); i++ {
		fmt.Printf("%2d", i)
		for j := 0; j < len(arrs[0]); j++ {
			if arrs[i][j] {
				fmt.Printf("%2s", "*")
			} else {
				fmt.Printf("%2s", "^")
			}
		}
		fmt.Println()
	}
}
