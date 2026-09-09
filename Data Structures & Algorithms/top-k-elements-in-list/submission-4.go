func topKFrequent(nums []int, k int) []int {
    hash := make(map[int]int)
    for _, n := range nums {
        hash[n]++
    }
    freq := make([][]int, len(nums) + 1)
    for num, count := range hash {
        freq[count] = append(freq[count], num)
    }
    var out []int
    for i := len(freq) - 1; i > 0; i-- {
        for _, num := range freq[i] {
            out = append(out, num)
            if len(out) == k {
                return out
            }
        }
    }
    return out
}
