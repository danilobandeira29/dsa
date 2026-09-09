import (
	"slices"
	"maps"
)

func topKFrequent(nums []int, k int) []int {
	slices.Sort(nums)
	j := len(nums) - 1
	i := j - 1
	hash := make(map[int][]int)
	for i >= 0 {
		if nums[j] == nums[i] {
			i--
			continue
		}
		freq := j - i
		hash[freq] = append(hash[freq], nums[j])
		j = i
		i--
	}
	freq := j - i
	hash[freq] = append(hash[freq], nums[j])
	sortedKeys := slices.Sorted(maps.Keys(hash))
	slices.Reverse(sortedKeys)
	var out []int
	for _, key := range sortedKeys {
		out = append(out, hash[key]...)
	}
	return out[:k]
}
