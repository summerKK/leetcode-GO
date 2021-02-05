package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/summerKK/leetcode-Go/algs4/cmd/1.3"
)

var colo *color.Color

/**
1.3.43
文件列表。文件夹就是一列文件和文件夹的列表。编写一个程序，从命令行接受一个文件夹名
作为参数，打印出该文件夹下的所有文件并用递归的方式在所有子文件夹的名下(缩进)列出
其下的所有文件。提示:使用队列，并参考 java.io.File。
*/
func main() {
	colo = color.New(color.FgYellow)
	queue := &__3.MyQueue{}
	_ = dir("/Users/summer/Docker/go-summer/leetcode-Go", queue, -1)
}

func dir(path string, queue *__3.MyQueue, depth int) (err error) {
	depth++
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return
	}

	level := ""
	for i := 0; i < depth; i++ {
		level += fmt.Sprint(" ")
	}
	fileName := fileInfo.Name()
	if fileInfo.IsDir() {
		fileName += "~~"
	}
	queue.Enqueue(level + fileName)

	if fileInfo.IsDir() {
		fileInfos, err := file.Readdir(0)
		if err != nil {
			return err
		}
		for _, fileInfo := range fileInfos {
			err = dir(path+"/"+fileInfo.Name(), queue, depth)
			if err != nil {
				return err
			}
		}
	}

	for fileName := queue.Dequeue(); fileName != nil; fileName = queue.Dequeue() {
		if strings.Contains(fileName.(string), "~~") {
			_, _ = colo.Println(strings.TrimRight(fileName.(string), "~~"))
		} else {
			fmt.Println(fileName)
		}
	}

	return nil
}
