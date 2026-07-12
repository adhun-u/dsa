package stack

func IsValid(s string) bool {

	stack := []byte{}

	for i := range s {

		bracket := s[i]

		switch bracket {
		case '{', '[', '(':
			stack = append(stack, bracket)
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] == '[' || stack[len(stack)-1] == '(' {
				return false
			}

			stack = stack[:len(stack)-1]
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] == '{' || stack[len(stack)-1] == '[' {
				return false
			}

			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] == '{' || stack[len(stack)-1] == '(' {
				return false
			}

			stack = stack[:len(stack)-1]

		}

	}

	return len(stack) == 0
}
