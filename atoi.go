package student

//StrLen returns strings length
func StrLen(str string) int {
	n := 0
	for range str {
		n++
	}
	return n
}

// CheckNum returns true if number is correct
func CheckNum(str string) bool {
	negative := false
	if StrLen(str) >= 20 && str[0] == '-' {
		negative = true
		str = str[1:]
	}
	if StrLen(str) > 19 {
		return false
	} else if StrLen(str) == 19 {
		if negative && str > "9223372036854775808" || !negative && str > "9223372036854775807" {
			return false
		}
	}
	for i, j := range str {
		if j == '-' || j == '+' {
			if i != 0 || StrLen(str) == 1 {
				return false
			}
		} else if j < '0' || j > '9' {
			return false
		}
	}
	return true
}

// Atoi returns int from given string
func Atoi(str string) int {
	if CheckNum(str) {
		res, sign := 0, 1
		for _, j := range str {
			c := 0
			if j == '-' {
				sign = -1
			}
			for i := '0'; i < j; i++ {
				c++
			}
			res = res*10 + c
		}
		return res * sign
	}
	return 0
}
