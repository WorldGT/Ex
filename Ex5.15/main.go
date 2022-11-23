/*
编写类似sum的可变参数函数max和min。
考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。
*/
package main

import "fmt"

func main() {
	fmt.Println(max(1, 2, 3, 4, 5, 6))
	fmt.Println(mix(1, 2, 3, 4, 5, 0))
	//fmt.Println(max())
}

func max(vs ...int) int {
	if len(vs) == 0 {
		panic("vs is nil")
	}
	var m int = 1
	for _, v := range vs {
		if v > m {
			m = v
		}
	}
	return m
}

func mix(vs ...int) int {
	if len(vs) == 0 {
		panic("vs is nil")
	}
	var m int = 1
	for _, v := range vs {
		if v < m {
			m = v
		}
	}
	return m
}
