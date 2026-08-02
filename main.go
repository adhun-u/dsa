package main

import (
	"dsa/sorting"
	"fmt"
)

func main() {
	out := sorting.InsertSort([]int{9, 4, 5, 9, 6, 0, 7, 0})
	fmt.Println("Out : ", out)
}
