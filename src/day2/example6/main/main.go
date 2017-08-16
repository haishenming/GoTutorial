package main

import (
	"fmt"
	"time"
	"math/rand"
)

func init() {
	rand.Seed(time.Now().UnixNano())     // 初始化随机种子
}

func main() {

	fmt.Println("生成10个随机整数")
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}

	fmt.Println("随机生成10个小于100的整数")
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}

	fmt.Println("生成十个随机浮点数")
	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}

}
