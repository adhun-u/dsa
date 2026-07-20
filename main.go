package main

import (
	"dsa/stack"
	"fmt"
)

func main() {
	out := stack.CarFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3})
	fmt.Println("Out : ", out)
}
