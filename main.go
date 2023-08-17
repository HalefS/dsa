package main

import (
	"fmt"

	"github.com/HalefS/dsa/stack"
)

func main() {

	stack := stack.New()
	stack.Push(20)
	stack.Push(200)
	stack.Push(500)

	fmt.Println(*stack.Pop())
}
