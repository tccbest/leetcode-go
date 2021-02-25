package _150_evaluate_reverse_polish_notation

import "strconv"

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	stack := []int{}
	str := map[string]struct{}{
		"+": struct{}{},
		"-": struct{}{},
		"*": struct{}{},
		"/": struct{}{},
	}

	for _, v := range tokens {
		if _, ok := str[v]; ok {
			if len(stack) < 2 {
				return -1
			}

			stackLen := len(stack)
			a := stack[stackLen-1]
			b := stack[stackLen-2]
			stack = stack[:stackLen-2]
			var result int
			switch v {
			case "+":
				result = b + a
			case "-":
				result = b - a
			case "*":
				result = b * a
			case "/":
				result = b / a
			}

			stack = append(stack, result)
		} else {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
