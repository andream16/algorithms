package strobogrammaticnumber

// A number is a stroboscopic only if it's in the map below.
// Concatenate the numbers the other way around and compare with the input.
//
// T: O(n)
// S: O(1)
func isStrobogrammatic(num string) bool {

	m := map[rune]rune{
		'0' : '0',
		'1' : '1',
		'6' : '9',
		'8' : '8',
		'9' : '6',
	}

	var str string

	for _, c := range num {
		r, ok := m[c]
		if !ok {
			return false
		}
		str = string(r) + str
	}

	return num == str

}
