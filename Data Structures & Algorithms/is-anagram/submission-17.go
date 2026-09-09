// func isAnagram(s string, t string) bool {
// 	l1, l2 := len(s), len(t)
// 	if l1 != l2 {
// 		return false
// 	}
// 	h1 := make(map[rune]int)
// 	h2 := make(map[rune]int)
// 	for _, e := range s {
// 		h1[e]++
// 	}
// 	for _, e := range t {
// 		h2[e]++
// 	}
// 	for k, v := range h1 {
// 		if _, ok := h2[k]; !ok {
// 			return false
// 		}
// 		if h2[k] != v {
// 			return false
// 		}
// 	}
// 	return true
// }

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var counter[26] int
	for i := range s {
		counter[s[i] - 'a']++
		counter[t[i] - 'a']--
	}
	for i := range counter {
		if counter[i] != 0 {
			return false
		}
	}
	return true
}
