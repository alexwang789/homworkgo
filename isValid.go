func isValid(s string) bool {
	stack := []rune{} // 使用slice作为栈
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else { 
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != pairs[char] {
				return false
			}
		}
	}
	return len(stack) == 0
}
