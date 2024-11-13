package main

func IsValid(s string) bool {
	stack, stackPointer := make([]byte, len(s)), 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '{', '[':
			stack[stackPointer] = s[i]
			stackPointer++
		case ')':
			if stackPointer == 0 || stack[stackPointer-1] != '(' {
				return false
			}
			stackPointer--
		case ']':
			if stackPointer == 0 || stack[stackPointer-1] != '[' {
				return false
			}
			stackPointer--
		case '}':
			if stackPointer == 0 || stack[stackPointer-1] != '{' {
				return false
			}
			stackPointer--
		}
	}

	return stackPointer == 0
}
