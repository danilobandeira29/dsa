func isPalindrome(s string) bool {
	var (
		right = len(s) - 1
		left int
	)
	for left < right {
		for left < right && isNotAlphaNumeric(s[left]) {
			left++
		}
		for left < right && isNotAlphaNumeric(s[right]) {
			right--
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isNotAlphaNumeric(b byte) bool {
	return !((b >= 'a' && b <= 'z') || 
		   (b >= 'A' && b <= 'Z') || 
		   (b >= '0' && b <= '9'))
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
