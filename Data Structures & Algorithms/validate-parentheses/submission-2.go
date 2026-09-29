func isValid(s string) bool {
	closeToOpen := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0, len(s))
	for i := range s {
		if _, ok := closeToOpen[s[i]]; !ok {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		lastElement := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if lastElement != closeToOpen[s[i]] {
			return false
		}
	}
	return len(stack) == 0
}
