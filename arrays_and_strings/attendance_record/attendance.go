package main

import "fmt"

func checkRecord(s string) bool {
	if len(s) == 2 {
		return !(s[0] == 'A' && s[1] == 'A')
	}

	var (
		numAbsences = 0
		numLates    = 0
	)

	for _, c := range s {
		if numAbsences > 1 {
			return false
		}
		if c == 'A' {
			numAbsences++
		}
		if c != 'L' {
			numLates = 0
		} else {
			numLates++
		}
		if numLates >= 3 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(checkRecord("PPALLP"))  // true
	fmt.Println(checkRecord("PPALLL"))  // false
	fmt.Println(checkRecord(""))        // true
	fmt.Println(checkRecord("A"))       // true
	fmt.Println(checkRecord("AA"))      // false
	fmt.Println(checkRecord("L"))       // true
	fmt.Println(checkRecord("LLL"))     // false
	fmt.Println(checkRecord("LALL"))    // true
	fmt.Println(checkRecord("AL"))      // true
	fmt.Println(checkRecord("ALLAPPL")) // false
}
