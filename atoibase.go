package student


// Slength lol
func Slength(str string) int {
	var res int
	for i := range str {
		res = i + 1
	}
	return res
}

//CheckBase lol
func CheckBase(base string) bool {
	var com string
	if Slength(base) < 2 {
		return false
	}
	for _, j := range base {
		if j == '-' || j == '+' {
			return false
		}
		for _, k := range com {
			if k == j {
				return false
			}
		}
		com += string(j)
	}
	return true
}

func CheckStrNum(num, base string) bool {
	for _, j := range num {
		flag := false
		for _, l := range base {
			if j == l  {
				flag = true
			}
		}
			if !flag {
				return false
			}
	}
	return true
}

//AtoiBase lol
func AtoiBase(s string, base string) int {
	var res = 0
	if CheckBase(base) && CheckStrNum(s, base) {
		for _, j := range s {
			for k, l := range base {
				if j == l {
					res = res*Slength(base) + k
				}
			}
		}
		return res
	}
	return 0
}
