func hasDuplicate(nums []int) bool {
    hash := make(map[int]struct{}, len(nums))
    for _, e := range nums {
        if _, ok := hash[e]; ok {
            return true
        }
        hash[e] = struct{}{}
    }
    return false
}
