
/*
输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数。
 */

package main

import (
	"fmt"
	"unicode"
)

func main() {

	var st string
	var digit, letter, space, other int = 0,0,0,0

	fmt.Printf("输入一个字符串：")
	fmt.Scanf("%s", &st)
	for _,s := range st {
		if unicode.IsDigit(s) {
			fmt.Printf("%c 是数字\n", s)
			digit ++
		} else if unicode.IsLetter(s) {
			fmt.Printf("%c是字母\n", s)
			letter ++
		} else if unicode.IsSpace(s) {
			fmt.Printf("%c是空格\n", s)
			space ++
		} else {
			fmt.Printf("%c是其他字符\n", s)
			other ++
		}
	}

	fmt.Printf("数字：%d，字母：%d，空格：%d，其他字符：%d", digit, letter, space, other)

}

