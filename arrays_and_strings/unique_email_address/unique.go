package main

import (
	"fmt"
	"strings"
)

// T: O(N * (log N)) + O(N)
// S: l = |e|; 3*l
func numUniqueEmails(emails []string) int {
	localDomains, domainNames := make([]string, 0, len(emails)), make([]string, 0, len(emails))
	for _, e := range emails {
		parts := strings.Split(e, "@")
		localDomain := strings.Split(strings.Replace(parts[0], ".", "", -1), "+")[0]
		localDomains = append(localDomains, localDomain)
		domainNames = append(domainNames, parts[1])
	}
	uniqueEmails := make(map[string]struct{}, len(emails))
	for i := 0; i < len(emails); i++ {
		email := localDomains[i] + "@" + domainNames[i]
		uniqueEmails[email] = struct{}{}
	}
	return len(uniqueEmails)
}

func main() {
	fmt.Println(numUniqueEmails([]string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}))
}
