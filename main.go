package main

import (
	"fmt"

	"github.com/HalefS/dsa/stack"
)

func main() {

	stack := stack.New()
	stack.Put(20)
	stack.Put(200)
	stack.Put(500)

	fmt.Println(*stack.Pop())
}
