package decodeways

// To solve this problem we can use Dynamic Programming.
// We analyse digits starting from left and going to the right.
// Any digit by itself is valid. Two digits together are valid if summed up they are in the range 10 <= D <= 26.
// Every time we find a digit, it means that we have a potential new way of decoding. Same thing for when we have
// two consecutive digits that satisfy what mentioned above. Last thing is, we need to have a criteria to stop and
// to discard invalid decode ways. We stop when the remaining string is empty - it means that we've processed
// all the decode ways for that branch and we can return 1, as it's a new decode way. A decode way is invalid if it
// starts with 0 as there's no mapping for it. Now, for the DP bit - we use a map to  contain the analysed strings and
// their respective results. We do this because in big inputs, there's a lot of repetition, so, having these resulta
// in memory is very useful.
//
// E.G.: 226
//
// (2 + decodeWays(26)), (22 + decodeWays(6))
// (22 + decodeWays(6)), (2,2 + decodeWays(6)) BZ, VF
// VF, BBF, BZ, VF -> BBF, BZ, VF
//
// T: O(2^n) with no memoization, O(n) with memoization
// S: O(n)
func numDecodings(s string) int {
	return numDecodingsRec(s, map[string]int{})
}

func numDecodingsRec(s string, m map[string]int) int {
	if s == "" {
		return 1
	} else if '0' == s[0] {
		return 0
	}

	if r, ok := m[s]; ok {
		return r
	}

	res := numDecodingsRec(s[1:], m)
	if len(s) > 1 {
		if s[0:2] >= "10" && s[0:2] <= "26" {
			res += numDecodingsRec(s[2:], m)
		}
	}

	m[s] = res
	return res
}
