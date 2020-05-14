package replacestr

import (
	"sort"
	"strings"
)

type group struct {
	index  int
	source string
	target string
}

// Since indexes is not guaranteed to be sorted, we can sort it, together with its corresponding sources & targets
// to help us in this problem. Once the sorting is done (from highest index to lowest) we can apply the following steps:
//  - if S at index i doesn't begin with source src, we can simply skip it.
//  - if S at index i begins with source src, then we copy everything on the right of the current index, concatenate
//    the target and concatenate everything else on the left of the current index after the length of the target.
//
// T: O(n log n)
// S: O(n)
func findReplaceString(S string, indexes []int, sources []string, targets []string) string {

	groups := make([]group, len(indexes))
	for i := 0; i < len(indexes); i++ {
		groups[i] = group{
			index:  indexes[i],
			source: sources[i],
			target: targets[i],
		}
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].index > groups[j].index
	})

	for _, group := range groups {
		postfix := S[group.index:]
		if !strings.HasPrefix(postfix, group.source) {
			continue
		}
		S = S[:group.index] + group.target + S[group.index+len(group.source):]
	}

	return S
}
