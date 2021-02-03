package main

import "fmt"

/**
1.3.37
Josephus 问题。在这个古老的问题中，N 个身陷绝境的人一致同意通过以下方式减少生存人 数。他们围坐成一圈(位置记为 0 到 N-1)并从第一个人开始报数，报到 M 的人会被杀死， 直到最后一个人留下来。传说中 Josephus 找到了不会被杀死的位置。编写一个 Queue 的用例
Josephus，从命令行接受 N 和 M 并打印出人们被杀死的顺序(这也将显示 Josephus 在圈中的位置)。
*/
func main() {
	i := josephus0(11, 3) + 1
	fmt.Println(i)

	i = josephus1(11, 3)
	fmt.Println(i)
}

func josephus0(nums int, flag int) int {
	if nums == 1 {
		return 0
	}

	return (josephus0(nums-1, flag) + flag) % nums
}

func josephus1(nums int, flag int) int {
	arr := make([]int, nums)
	for i := 0; i < nums; i++ {
		arr[i] = i + 1
	}
	endIndex := nums - 1
	currentFlag := 1
	currentIndex := 0
	for endIndex > 0 {
		// 数到第n个人
		if currentFlag == flag {
			fmt.Printf("kill %d\n", arr[currentIndex])
			// 删掉第 currentIndex 个人,后面元素往前移动
			for i := currentIndex; i < nums-1; i++ {
				arr[i] = arr[i+1]
			}
			arr[nums-1] = 0
			currentFlag = 1
			endIndex--
			// 数组元素往前移动需要减一
			currentIndex--
		} else {
			currentFlag++
		}
		// 一组循环结束.下次循环从尾到头(比如:10 11 1)
		if currentIndex == endIndex {
			currentIndex = 0
		} else {
			currentIndex++
		}
	}

	return arr[0]
}
