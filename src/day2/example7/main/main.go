package main

import (
	"fmt"
)

type Rect struct {
	a, b float64
	width, height float64
}

func (r *Rect) Area() float64  {
	return r.width * r.height
}

func main() {
	r := &Rect{0, 0, 100, 200}

	fmt.Println(r.Area())
}
