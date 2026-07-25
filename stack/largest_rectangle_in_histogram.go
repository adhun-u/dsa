package stack

func LargestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0
	n := len(heights)

	for i := 0; i <= n; i++ {

		var currHeight int
		if i == n {
			currHeight = 0
		} else {
			currHeight = heights[i]
		}

		for len(stack) != 0 && heights[stack[len(stack)-1]] > currHeight {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}

			width := i - left - 1
			area := height * width

			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}
