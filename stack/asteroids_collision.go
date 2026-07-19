package stack

func isSameOp(first int, second int) bool {
	return (first > 0 && second > 0) || (first < 0 && second < 0)
}
func AsteroidCollision(asteroids []int) []int {

	stack := []int{}

	for i := range asteroids {
		as := asteroids[i]
		stack = append(stack, as)

		for len(stack) >= 2 && !isSameOp(stack[len(stack)-2], as) {

			last := stack[len(stack)-1]
			secLast := stack[len(stack)-2]
			posLast := last

			if last > 0 {
				break
			} else {
				posLast = -(posLast)
			}

			if posLast == secLast {
				stack = stack[:len(stack)-2]
			} else if posLast > secLast {
				stack = stack[:len(stack)-2]
				stack = append(stack, last)
			} else {
				stack = stack[:len(stack)-2]
				stack = append(stack, secLast)
			}
		}

	}

	return stack
}
