package stack

func DailyTemperatures(temperatures []int) []int {

	stack := []int{}
	out := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {

		temp := temperatures[i]

		if len(stack) == 0 {
			stack = append(stack, i)
			out[i] = 0
			continue
		}

		prevTemp := temperatures[i+1]

		if temp < prevTemp {
			out[i] = 1
		} else {

			for len(stack) != 0 && temperatures[stack[len(stack)-1]] <= temp {
				stack = stack[:len(stack)-1]
			}

			if len(stack) != 0 {
				out[i] = stack[len(stack)-1] - i
			} else {
				out[i] = 0
			}

		}

		stack = append(stack, i)

	}

	return out

}
