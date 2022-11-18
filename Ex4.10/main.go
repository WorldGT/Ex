// 4.10
// 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。

// go run main.go repo:golang/go is:open json decoder
package main

import (
	"GoEx/Ex4.10/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	mbefore := time.Now().AddDate(0, -1, 0) // 一个月前
	ybefore := time.Now().AddDate(-1, 0, 0) // 一年前的时间

	monbefore := []*github.Issue{}
	yeabefore := []*github.Issue{}
	yeaafter := []*github.Issue{}
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatalf("解析错误 : %s", err)
	}

	// 遍历所有items
	fmt.Println("创建时间不到一个月:")
	for _, item := range result.Items {
		//时间
		// fmt.Printf("item.CreatedAt: %v\n", item.CreatedAt)
		if mbefore.Before(item.CreatedAt) { // 创建时间不到一个月
			// fmt.Println("一月:\t ", item.CreatedAt)
			monbefore = append(monbefore, item)

		} else if ybefore.Before(item.CreatedAt) { //创建时间不到一年
			// fmt.Println("不到一年:\t", item.CreatedAt)
			yeabefore = append(yeabefore, item)

		} else if ybefore.After(item.CreatedAt) { // 创建时间超过一年
			// fmt.Println("超过:\t", item.CreatedAt)
			yeaafter = append(yeaafter, item)
		}
	}

	fmt.Println("创建时间不到一个月:")
	for _, mbefores := range monbefore {
		fmt.Println(mbefores.CreatedAt)
		fmt.Printf("#%-5d %9.9s %.55s\n", mbefores.Number, mbefores.User.Login, mbefores.Title)
	}

	fmt.Println("创建时间不到一年:")
	for _, yeabefores := range monbefore {
		fmt.Println(yeabefores.CreatedAt)
		fmt.Printf("#%-5d %9.9s %.55s\n", yeabefores.Number, yeabefores.User.Login, yeabefores.Title)
	}
	fmt.Println("创建时间超过一年:")
	for _, yeaafters := range monbefore {
		fmt.Println(yeaafters.CreatedAt)
		fmt.Printf("#%-5d %9.9s %.55s\n", yeaafters.Number, yeaafters.User.Login, yeaafters.Title)
	}
}
