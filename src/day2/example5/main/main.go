package main

import "fmt"

func main() {
	var a int
	var b int32
	a = 15
	b = int32(a) + int32(a) // 即使都是int，不同int的变量也是不能运算的，必须先转换
	b = b + 5

	fmt.Println(a)
	fmt.Println(b)
}
