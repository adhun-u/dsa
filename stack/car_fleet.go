package stack

import (
	"slices"
)

func CarFleet(target int, position []int, speed []int) int {

	speedMap := map[int]int{}
	stack := []int{}

	for i := range position {
		speedMap[position[i]] = speed[i]
	}

	slices.Sort(position)

	for i := len(position) - 1; i >= 0; i-- {
		stack = append(stack, position[i])

		if len(stack) >= 2 {
			last := stack[len(stack)-1]
			secLast := stack[len(stack)-2]
			lastTime := (float32(target) - float32(last)) / float32(speedMap[last])
			secLastTime := (float32(target) - float32(secLast)) / float32(speedMap[secLast])

			if lastTime <= secLastTime {
				stack = stack[:len(stack)-2]
				stack = append(stack, secLast)
			}
		}
	}
	return len(stack)
}
