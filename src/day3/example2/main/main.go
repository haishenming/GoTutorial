package main

import (
	"fmt"
	"day3/example2/str"
)

func main() {
	s := " 1 2 3 1 jjjdh jksk  1."

	s1, s2 ,s3, s4 := str.Str(s)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

