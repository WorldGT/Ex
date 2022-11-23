//编写多参数版本的strings.Join。

package main

import (
	"bytes"
	"fmt"
)

func stringsJoin(sep string, elems ...string) string {
	var buf bytes.Buffer
	if len(elems) == 0 {
		return "elems为空"
	}
	buf.WriteString(elems[0])
	for _, elem := range elems[1:] {
		buf.WriteString(sep)
		buf.WriteString(elem)
	}
	return buf.String()
}

func main() {

	fmt.Println(stringsJoin(" 分隔 "))
	fmt.Println(stringsJoin(" 分隔 ", "ab"))
	fmt.Println(stringsJoin(" 分隔 ", "ab", "cd"))
	fmt.Println(stringsJoin(" 分隔 ", "ab", "cd", "dasd"))
}
