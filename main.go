package main

import (
	array "dsa/arrays_and_hashes"
	"fmt"
)

func main() {

	str := []string{}

	s := array.Solution{}
	encoded := s.Encode(str)
	fmt.Println("Out of encoded : ", encoded)
	fmt.Println("Out of decoded : ", len(s.Decode(encoded)))
}
