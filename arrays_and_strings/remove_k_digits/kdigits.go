package kdigits

// TODO CHECK THIS
func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

	for ; k > 0; k-- {
		for i := 0; i < len(num)-1; i++ {
			if num[i] > num[i+1] {
				if i+1 != len(num)-1 && num[i+1] == '0' {
					num = num[:i] + num[i+2:len(num)]
				} else {
					num = num[:i] + num[i+1:len(num)]
				}
				break
			} else if num[i] <= num[i+1] {
				num = num[:i+1] + num[i+2:len(num)]
			}
		}
	}

	return num
}
