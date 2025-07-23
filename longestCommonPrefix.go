func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; ; i++ {
		for _, s := range strs {
			if i >= len(s) || s[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
}
