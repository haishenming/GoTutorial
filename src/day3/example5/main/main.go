package main

import (
	"fmt"
)

func main() {
	var a int
	defer fmt.Println("defer",a)

	a++
	fmt.Println("正常：",a)
}
