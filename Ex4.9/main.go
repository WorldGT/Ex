//编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用
//Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	words := make(map[string]int)

	// 终端使用ctrl d 停止
	for input.Scan() {
		words[input.Text()]++
	}

	for w, j := range words {
		fmt.Printf("%s\t%d\n", w, j)
	}

}
