package main

import (
	"dsa/stack"
	"fmt"
)

func main() {
	// out := stack.DailyTemperatures([]int{30, 30})
	// fmt.Println("Out : ", out)

	queue := stack.QueueConstructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	fmt.Println("Peek : ", queue.Peek())
	fmt.Println("Pop : ", queue.Pop())
	fmt.Println("Peek : ", queue.Peek())
	fmt.Println("Empty : ", queue.Empty())
}
