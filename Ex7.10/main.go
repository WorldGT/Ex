// 编写一个IsPalindrome(s sort.Interface)bool函数表明序列s是否是回文序列
package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	if s.Len() == 0 {
		return false
	}
	i, j := 0, s.Len()-1
	for i < j {
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func main() {
	a := []int{1, 2, 3, 4, 3, 2, 1}
	fmt.Println(IsPalindrome(sort.IntSlice(a)))
	a = []int{2, 2, 3, 4, 3, 2, 1}
	fmt.Println(IsPalindrome(sort.IntSlice(a)))

}
