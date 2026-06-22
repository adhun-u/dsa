package main

import (
	arraysandhashes "dsa/arrays_and_hashes"
	"fmt"
)

func main() {

	output := arraysandhashes.GroupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"})

	fmt.Println("Value : ", output)

}
