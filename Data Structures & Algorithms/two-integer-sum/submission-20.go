func twoSum(nums []int, target int) (out []int) {
    hash := make(map[int]int)
	for i := range nums {
		seen, ok := hash[target-nums[i]]
		if !ok {
			hash[nums[i]] = i
			continue
		}
		if seen == i {
			continue
		}
		out = append(out, seen, i)
		return
	}
	return
}
