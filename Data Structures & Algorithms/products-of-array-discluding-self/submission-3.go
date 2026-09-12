func productExceptSelf(nums []int) []int {
	suffix := 1
	out := make([]int, len(nums))
	out[0] = 1
	for i := 0; i < len(nums) - 1; i++ {
		suffix *= nums[i]
		out[i+1] = suffix
	}
	suffix = 1
	for i := len(nums) - 1; i >= 0; i-- {
		out[i] = suffix * out[i]
		suffix *= nums[i]
	}
	return out
}
