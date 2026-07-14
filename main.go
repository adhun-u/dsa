package main

import (
	"dsa/stack"
	"fmt"
)

func main() {
	out := stack.EvalRPN([]string{"100", "0", "*"})
	fmt.Println("Out : ", out)
}
