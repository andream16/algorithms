package romantointeger

func toInt(s byte) int {
	switch s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

// IV
// c "I", n "V" if c < n -> -c
// c "V", n ""  if n == "" || c >= n -> +c

// VIII
// c "V", n "I" if c >= n -> +c
// c "I", n "I" if c >= n -> +c
// c "I", n "I" if c >= n -> +c
// c "I", n "" if n == "" || c >= n -> +c

func getToAdd(curr, next byte) int {
	c, n := toInt(curr), toInt(next)

	switch {
	case c >= n:
		return c
	case c < n:
		return c * -1
	default:
		return 0
	}

}

// TODO REVIEW
func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	var (
		n = 0
		l = len(s)
	)

	for i := range s {

		curr := s[i]

		if i == l-1 {
			n += toInt(s[i])
			break
		}

		n += getToAdd(curr, s[i+1])

	}

	return n
}
