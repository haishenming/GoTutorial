package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)

	fmt.Println(now.Format("2006/1/2 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("15:04:05 2006/1/2"))
	fmt.Println(now.Format("2006/1/2"))
}

