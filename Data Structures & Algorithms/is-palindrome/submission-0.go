func isPalindrome(s string) bool {
	s = strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, ""))
	j := len(s) - 1
	i := 0
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
