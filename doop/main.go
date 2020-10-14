package main

import (
	student ".."
	"os"

	"github.com/01-edu/z01"
)

const max = 9223372036854775807
const min = -9223372036854775808

// CheckOp returns true if operator is valid
func CheckOp(str string) bool {
	len := 0
	for range str {
		len++
	}
	for _, j := range str {
		if len == 1 && (j == '+' || j == '-' || j == '/' || j == '*' || j == '%') {
			return true
		}
	}
	return false
}

// PrintRes prints int
func PrintRes(num int) {
	if num == 0 {
		return
	}
	PrintRes(num / 10)
	z01.PrintRune(rune(num%10 + 48))
}

// PrintStr prints strings followed by newline
func PrintStr(msg string) {
	for _, j := range msg {
		z01.PrintRune(j)
	}
	z01.PrintRune(10)
}

func myAbs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func mySign(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	}
	return 1
}

//returns true if allowed
func checkResult(n1, n2, sign int) bool {
	if sign == 0 { // multiply
		if (n1 == -9223372036854775808 && n2 == 1) || (n2 == -9223372036854775808 && n1 == 1) {
			return true
		}
		if n2 != 0 && (myAbs(n1) > 9223372036854775807/myAbs(n2)) {
			return false
		} else if n1 != 0 && (myAbs(n2) > 9223372036854775807/myAbs(n1)) {
			return false
		} else {
			return true
		}
	} else if sign == 2 && ((n1 == -9223372036854775808 && n2 == -1) || (n2 == -9223372036854775808 && n1 == -1)) {
		return false
	} else if sign == 3 && (n1 == -9223372036854775808 && n2 == 1) {
		return false
	} else if sign == -1 || sign == 1 {
		if n1 == 0 && n2 == -9223372036854775808 && sign == -1 { // minus, exception
			return false
		}

		if n1 != 0 && n2 != 0 && mySign(n1)*mySign(n2)*sign > 0 { // res goes to overflow
			if myAbs(n1)+myAbs(n2) < 0 {
				// fmt.Println(myAbs(n1), myAbs(n2), myAbs(n1)+myAbs(n2))
				if sign == 1 && mySign(n1) == -1 && mySign(n2) == -1 && myAbs(n1)+myAbs(n2) != -9223372036854775808 {
					return false
				}
				return false
			}
		}
	}
	return true
}

func main() {
	var args = os.Args
	len, res := 0, 0
	msg1 := "No division by 0"
	msg2 := "No modulo by 0"
	for range args {
		len++
	}
	if len == 4 {
		if student.CheckNum(args[1]) && student.CheckNum(args[3]) && CheckOp(args[2]) {
			num1 := student.Atoi(args[1])
			num2 := student.Atoi(args[3])
			op := args[2]
			if op == "+" && checkResult(num1, num2, 1) {
				res = num1 + num2
			} else if op == "-" && checkResult(num1, num2, -1) {
				res = num1 - num2
			} else if op == "*" && checkResult(num1, num2, 0) {
				res = num1 * num2
			} else if op == "/" && checkResult(num1, num2, 2) {
				if num2 == 0 {
					PrintStr(msg1)
					return
				}
				res = num1 / num2
			} else if op == "%" && checkResult(num1, num2, 3) {
				if num2 == 0 {
					PrintStr(msg2)
					return
				}
				if mySign(num1)*mySign(num2) == -1 {
					res = num1%num2 + num2
				} else {
					res = num1 % num2
				}
			}
			if res == min {
				PrintStr("-9223372036854775808")
				return
			} else if res < 0 {
				res *= -1
				z01.PrintRune('-')
				PrintRes(res)
			} else if res == 0 {
				z01.PrintRune('0')
			} else {
				PrintRes(res)
			}
		} else {
			z01.PrintRune('0')
		}
		z01.PrintRune('\n')
	}
}


//  ./doop 861 + 870
// 1731
//  ./doop 861 - 870
// -9
// ./doop 861 "*" 870
// 749070
// ./doop 861 % 870
// 861
// ./doop  hello + 1
// 0
// ./doop 1 p 1
// 0
//  ./doop 1 / 0
// No division by 0
//  ./doop 1 % 0
// No modulo by 0
//  ./doop  1 "*" 1
// 1
//  ./doop 9223372036854775807 + 1
// 0
// ./doop 9223372036854775809 - 3
// 0
//  ./doop 9223372036854775807 "*" 3
// 0



