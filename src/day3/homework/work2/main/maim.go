
/*
一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3.
编程找出1000以内的所有完数。

 */

package main

import (
	"fmt"
)

func isPerfectNumber(num int) bool {

	if num == 0 {
		//fmt.Printf("0没有因子")
	} else if num == 1 {
		//fmt.Printf("1的因子是1, 1是完数")
		return true
	} else {
		var sumNum int
		for i:=1;i<num;i++{
			if num%i==0 {
				sumNum = sumNum + i
			}
		}
		if sumNum == num {
			//fmt.Printf("%d 是完数", num)
			return true
		}
	}

	return false
}

func main() {
	var num int
	for num=0;num<=1000;num++{
		if isPerfectNumber(num) {
			fmt.Printf("%d 是完数\n", num)
		}
	}

}

