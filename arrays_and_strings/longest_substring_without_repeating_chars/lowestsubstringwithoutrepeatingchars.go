package main

import "strings"

func main() {}

func lengthOfLongestSubstring(s string) int {

	ls := len(s)

	switch {
	case s == "":
		return 0
	case ls == 1:
		return 1
	}

	var (
		w string
		l int
	)

	for i, c := range s {
		cs := string(c)
		if i == 0 {
			w += cs
			continue
		}
		idx := strings.Index(w, cs)
		if idx == -1 {
			w += cs
			if i != ls-1 {
				continue
			}
		}
		lw := len(w)
		if lw > l {
			l = lw
		}
		w = string(w[idx+1:] + cs)
	}
	return l
}
