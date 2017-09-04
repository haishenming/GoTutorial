package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n, input int
	n = rand.Intn(100)

	for {
		var flag bool = false
		fmt.Println("I have a number, guess!")
		fmt.Scanf("%d\n", &input)

		switch {
		case input == n:
			fmt.Println("Right!")
			flag = true
		case input < n:
			fmt.Println("Low")
		case input > n:
			fmt.Println("Large")
		}

	if flag {
		break
	}

	}

}

