package main

import (
	"dsa/stack"
	"fmt"
)

func main() {
	out := stack.CarFleet(10, []int{6, 8}, []int{3, 2})
	fmt.Println("Out : ", out)
}
