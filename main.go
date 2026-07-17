package main

import (
	"dsa/stack"
	"fmt"
)

func main() {
	// out := stack.DailyTemperatures([]int{30, 30})
	// fmt.Println("Out : ", out)

	queue := stack.StackConstructor()
	queue.Push(1)
	queue.Push(2)

	fmt.Println("Top : ", queue.Top())
	// fmt.Println("Empty : ", queue.Empty())
	fmt.Println("Pop : ", queue.Pop())
	// queue.Push(10)
	fmt.Println("Top : ", queue.Top())
	fmt.Println("Pop : ", queue.Pop())
}
