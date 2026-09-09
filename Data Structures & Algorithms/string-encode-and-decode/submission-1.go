type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var out strings.Builder
	for _, str := range strs {
		fmt.Fprintf(&out, "%d#%s", len(str), str)
	}
	return out.String()
}

func (s *Solution) Decode(encoded string) []string {
	fmt.Println(encoded)
	var (
		out []string
		i int
	)
	for i < len(encoded) {
		j := i + 1
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		out = append(out, encoded[j + 1: j + 1 + length])
		i = j + 1 + length
	}
	return out
}
