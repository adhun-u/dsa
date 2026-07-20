package stack

import (
	"slices"
)

func CarFleet(target int, position []int, speed []int) int {
	posMap := map[int]int{}
	stack := []int{}

	for i := range position {
		pos := position[i]
		sp := speed[i]
		posMap[pos] = sp
	}

	slices.Sort(position)

	for i := len(position) - 1; i >= 0; i-- {
		pos := position[i]
		stack = append(stack, pos)

		if len(stack) >= 2 {
			last := stack[len(stack)-2]
			secLast := stack[len(stack)-1]
			lastT := float32((target - last) / posMap[last])
			secLastT := float32((target - secLast) / posMap[secLast])
			if secLastT <= lastT {
				stack = stack[:len(stack)-2]
				stack = append(stack, last)
			}
		}
	}

	return len(stack)
}
