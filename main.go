package main

import (
	"dsa/stack"
	"fmt"
)

func main() {
	// out := stack.DailyTemperatures([]int{30, 30})
	// fmt.Println("Out : ", out)

	s := stack.StockConstructor()
	span := s.Next(100)
	fmt.Println("Span : ", span)
	span = s.Next(120)
	fmt.Println("Span : ", span)
	span = s.Next(100)
	fmt.Println("Span : ", span)
	span = s.Next(110)
	fmt.Println("Span : ", span)
	span = s.Next(120)
	fmt.Println("Span : ", span)
	span = s.Next(50)
	fmt.Println("Span : ", span)
}
