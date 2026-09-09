func groupAnagrams(strs []string) [][]string {
    hash := make(map[[26]int][]string)
    for _, s := range strs {
        var freq [26]int
        for _, r := range s {
            freq[rune(r) - 'a']++
        }
        hash[freq] = append(hash[freq], s)
    }
    output := make([][]string, 0, len(hash))
    for _, v := range hash {
        output = append(output, v)
    }
    return output
}
