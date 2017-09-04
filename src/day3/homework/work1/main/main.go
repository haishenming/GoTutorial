/*
输出九九乘法表
 */

package main

import (
	"fmt"
)

var rows, column int = 9, 9

func main() {
	for r:=1;r<=rows;r++{
		for c:=1;c<=column;c++{
			if r >= c {
				fmt.Printf("%d*%d=%d \t", r, c, r*c)
			}
		}
		fmt.Printf("\n")
	}
}
