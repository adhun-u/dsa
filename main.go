package main

import (
	"dsa/stack"
	"fmt"
)

func main() {

	/*

		["MinStack", "push", 10, "pop", "push", 20, "top", "push", -20, "getMin"]

	*/
	minS := stack.Constructor()
	minS.Push(10)
	minS.Pop()
	minS.Push(20)
	fmt.Println("Top : ", minS.Top())
	minS.Push(-20)
	fmt.Println("Min : ", minS.GetMin())
	// fmt.Println("Min : ", minS.GetMin())
	// minS.Pop()
	// fmt.Println("Min  : ", minS.GetMin())
}
