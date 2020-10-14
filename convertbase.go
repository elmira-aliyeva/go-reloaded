package student

//Len lol
func Len(str string) int {
	var len int
	for range str {
		len++
	}
	return len
}

//Power lol
func Power(nb int, power int) int {
	if power > 0 {
		return nb * Power(nb, power-1)
	} else if power == 0 {
		return 1
	} else {
		return 0
	}
}

//Itoa lol
func Itoa(n int) string {
	ans, res := "", ""
	if n < 0 {
		n *= -1
		res += "-"
	} else if n == 0 {
		res = "0"
		return res
	}
	for i := n; i > 0; i /= 10 {
		r := rune(i%10 + 48)
		ans += string(r)
	}
	len := 0
	for range ans {
		len++
	}
	for i := len - 1; i >= 0; i-- {
		res += string(ans[i])
	}
	return res
}

//ConvertBase lol
func ConvertBase(nbr, baseFrom, baseTo string) string {
	lf, lt, ln := Len(baseFrom), Len(baseTo), Len(nbr)
	var dec = 0
	var nbr1, res, res1 string
	for i := ln - 1; i >= 0; i-- {
		nbr1 += string(nbr[i])
	}
	for i, j := range nbr1 {
		for a, b := range baseFrom {
			if j == b {
				dec += a * Power(lf, i)
			}
		}
	}
	// if Itoa(dec)[0] != nbr[0] {
	// 	return "0"
	// }
	for i := dec; i > 0; i = i / lt {
		res += string(baseTo[i%lt])
	}
	for i := Len(res) - 1; i >= 0; i-- {
		res1 += string(res[i])
	}
	return res1
}

// func ConvertBase(nbr, baseFrom, baseTo string) string {
// 	if nbr == student.NbrBase(AtoiBase(nbr, baseFrom), baseFrom) {
// 		return NbrBase(AtoiBase(nbr, baseFrom), baseTo)
// 	}
// 	return "0"
// }
