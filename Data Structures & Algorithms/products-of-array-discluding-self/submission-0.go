func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	prefix := make([]int, len(nums))
	prefix[0] = 1
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i - 1] * nums[i - 1]
	}
	posfix := make([]int, len(nums))
	posfix[len(nums) - 1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		posfix[i] = posfix[i + 1] * nums[i + 1]
	}
	out := make([]int, len(nums))
	for i := range nums {
		out[i] = prefix[i] * posfix[i]
	}
	return out
}
