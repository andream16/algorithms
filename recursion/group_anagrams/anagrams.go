package anagrams

// TODO WORK ON THIS
func groupAnagrams(strs []string) [][]string {
	return nil
}

func perm(s, r string, n int, res []string) []string {
	if n == len(s) {
		return append(res, s)
	}
	for i := 0; i < len(r); i++ {
		res = perm(s+string(r[i]), r[:i]+r[i+1:], n, res)
	}
	return res
}
