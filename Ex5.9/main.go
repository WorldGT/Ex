/*编写函数expand，将s中的"foo"替换为f("foo")的返回值。
func expand(s string, f func(string) string) string*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("1234 foo 56 foo", toUpper))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "foo", f("foo1 "), 2)
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}
