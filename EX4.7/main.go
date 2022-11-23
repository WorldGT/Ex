/* 4.7修改函数 reverse，来翻转一个 UTF-8 编码的字符串中的字符元素，
传入参数是该字符串对应的字节 slice 类型（[]byte）。
你可以做到不需要重新分配内存就实现该功能吗？  */

package main

import "fmt"

func reverse(arr []byte) string {
	arrlen := len(arr)
	for i := 0; i < arrlen/2; i++ {
		arr[i], arr[arrlen-i-1] = arr[arrlen-i-1], arr[i]
	}
	return string(arr)
}

func main() {
	a := "123456"
	fmt.Println(reverse([]byte(a)))
}
