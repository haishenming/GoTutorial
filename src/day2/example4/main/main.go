package main

import "fmt"


func swap(a int, b int) (int ,int ) {

	return b, a

}


func main() {
	a := 1
	b := 2
	fmt.Println(a, b)
	a, b = swap(a, b)
	fmt.Println(a, b)

}