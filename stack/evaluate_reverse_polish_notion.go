package stack

import "strconv"

func calc(stack []int, operator string) (newStack []int) {

	first := stack[len(stack)-2]
	second := stack[len(stack)-1]
	ans := 0

	switch operator {
	case "+":
		ans = first + second
	case "-":
		ans = first - second
	case "*":
		ans = first * second
	default:
		ans = first / second
	}

	stack = stack[:len(stack)-2]
	stack = append(stack, ans)

	return stack
}
func EvalRPN(tokens []string) int {

	if len(tokens) == 0 {
		return 0
	}

	stack := []int{}

	for i := range tokens {

		char := tokens[i]

		switch char {
		case "-":
			stack = calc(stack, char)

		case "+":
			stack = calc(stack, char)

		case "*":
			stack = calc(stack, char)

		case "/":
			stack = calc(stack, char)

		default:

			intValue, _ := strconv.Atoi(char)
			stack = append(stack, intValue)

		}
	}

	return stack[0]
}
