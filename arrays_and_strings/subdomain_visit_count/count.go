package count

import (
	"fmt"
	"strconv"
	"strings"
)

// We use a map where the key is a subdomain and the value is the count.
// For each entry, the first substring before a space is the count, the rest is the domain.
// The first valid domain is the last of the array obtained splitting the second substring by '.'.
// For the next valid domains, we concatenate the current entry to the previous one and increment it in the map.
// After we filled our map, we can simply loop through it to get the correct string format. This is linear.
//
// T: O(n)
// S: O(n)
func subdomainVisits(cpdomains []string) []string {
	m := map[string]int{}

	for _, s := range cpdomains {
		getCountAndDomains(s, m)
	}

	var res []string

	for k, v := range m {
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}

	return res
}

func getCountAndDomains(s string, m map[string]int) {

	p := strings.Split(s, " ")

	var (
		visits, _ = strconv.Atoi(p[0])
		rest      = strings.Split(p[1], ".")
	)

	curr := rest[len(rest)-1]
	m[curr] += visits

	for i := len(rest) - 2; i >= 0; i-- {
		curr = rest[i] + "." + curr
		m[curr] += visits
	}

}
