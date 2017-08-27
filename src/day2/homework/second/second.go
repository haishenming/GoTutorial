package second

/* 打印出100-999中所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字
立方和等于该数本身。例如：153 是一个“水仙花数”，因为 153=1 的三次
方＋5 的三次方＋3 的三次方。*/

import (
	"fmt"
)

func JudgeDaffodil() {
	var uni, decade, hundred int = 0, 0, 0
	var num int = 0
	for i := 101; i < 1000; i++ {
		uni = i % 10             // 个位
		decade = (i / 10) % 10   // 十位
		hundred = (i / 100) % 10 // 百位
		//iLi := []int(i)
		//fmt.Println(iLi)
		if uni*uni*uni+decade*decade*decade+hundred*hundred*hundred == i {
			fmt.Println(i)
			num ++
		}
	}
	fmt.Println("水仙花数的个数是：", num)

}
