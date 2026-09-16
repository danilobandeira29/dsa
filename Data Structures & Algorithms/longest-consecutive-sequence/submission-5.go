func longestConsecutive(nums []int) int {
	hash := make(map[int]struct{})
	for _, n := range nums {
		hash[n] = struct{}{}
	}
	out := make(map[int]int)
	var maxCount int
	for value := range hash {
		if _, ok := hash[value-1]; ok {
			continue
		}
		for current := value + 1; ; current++ {
			out[value]++
			if _, ok := hash[current]; !ok {
				break
			}
		}
		if out[value] > maxCount {
			maxCount = out[value]
		}
	}
	return maxCount
}
