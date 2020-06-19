package addbinary

// TODO REVIEW
func addBinary(a string, b string) string {
	var sum string
	helper(a, b, "0", len(a)-1, len(b)-1, &sum)
	return sum
}

func helper(s1, s2, carry string, i, j int, res *string) {
	if i < 0 && j < 0 {
		if carry == "1" {
			*res = "1" + *res
		}
		return
	}
	n1, n2, nextCarry := "0", "0", "0"
	if i >= 0 {
		n1 = string(s1[i])
	}
	if j >= 0 {
		n2 = string(s2[j])
	}
	switch {
	case n1 == "0" && n2 == "0":
		*res = carry + *res
	case n1 == "1" && n2 == "1":
		if carry == "0" {
			*res = "0" + *res
		} else {
			*res = "1" + *res
		}
		nextCarry = "1"
	default:
		if carry == "0" {
			*res = "1" + *res
		} else {
			*res = "0" + *res
			nextCarry = "1"
		}
	}
	helper(s1, s2, nextCarry, i-1, j-1, res)
}
