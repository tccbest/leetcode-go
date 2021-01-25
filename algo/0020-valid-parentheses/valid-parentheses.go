package _020_valid_parentheses

func IsValid(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	}

	m := map[string]string{"}": "{", "]": "[", ")": "("}

	stack := make([]string, 0)
	for i := 0; i < length; i++ {
		if len(stack) == 0 {
			stack = append(stack, string(s[i]))
		} else {
			slen := len(stack)
			istr := string(s[i])
			if m[istr] == stack[slen-1] {
				stack = stack[:slen-1]
			} else {
				stack = append(stack, istr)
			}
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
