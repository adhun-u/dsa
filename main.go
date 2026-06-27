package main

import (
	array "dsa/arrays_and_hashes"
	"fmt"
)

func main() {

	output := array.TopKFrequent([]int{1, 2, 2, 3, 3, 3}, 2)

	fmt.Println("Value : ", output)

}
