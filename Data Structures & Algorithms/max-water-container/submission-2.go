func maxArea(heights []int) int {
	var (
		area int
		l, r = 0, len(heights) - 1
	)
	for l < r {
		if diff := (r - l) * min(heights[l], heights[r]); diff > area {
			area = diff
		}
		if heights[r] <= heights[l] {
			r--
			continue
		}
		l++
	}
	return area
}
