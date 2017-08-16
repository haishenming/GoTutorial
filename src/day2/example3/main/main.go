package main

import (
	"fmt"
	"time"
)

const (
	male    = 1
	female = 2
)

func main() {
	now := time.Now().Unix()
	if (now % female == 0) {
		fmt.Println("female")
	} else {
		fmt.Println("man")
	}
}
