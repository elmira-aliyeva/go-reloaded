package student

import (
	"github.com/01-edu/z01"
)

// Length - strlen
func Length(str string) int {
	var res int
	for i := range str {
		res = i + 1
	}
	return res
}

// Check returns true if base is valid
func Check(base string) bool {
	var com string
	if Length(base) < 2 {
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

// PrintNbrBase you know what i do
func PrintNbrBase(nbr int, base string) {
	if Check(base) {
		isneg := false
		var ovf uint64 = 0
		max := false
		if nbr < 0 {
			isneg = true
			if nbr == -9223372036854775808 {
				ovf = 9223372036854775808
				max = true
			} else {
				nbr = nbr * -1
			}
		}
		var res string
		l := Length(base)
		for nbr != 0 {
			if !max {
				res += string(base[nbr%l])
			} else if max {
				res += string(base[ovf%uint64(l)])
			}
			nbr = nbr / l
			ovf = ovf / uint64(l)
		}
		if isneg {
			res += "-"
		}
		for i := Length(res) - 1; i >= 0; i-- {
			z01.PrintRune(rune(res[i]))
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}
