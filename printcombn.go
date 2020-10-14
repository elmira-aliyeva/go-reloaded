package student

import (
	"github.com/01-edu/z01"
)

//MyItoa lol
func MyItoa(n int) string {
	ans, res := "", ""
	if n == 0 {
		return "0"
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

//PrintCombN lol
func PrintCombN(n int) {
	var null int = 0
	var stringTemp string = ""
	RecursionFunc(null, n, stringTemp)
	z01.PrintRune('\n')
}

//RecursionFunc lol
func RecursionFunc(null, n int, stringTemp string) {
	var stringOut string
	if n == 0 {
		if stringTemp == "9" || stringTemp == "89" || stringTemp == "789" ||
			stringTemp == "6789" || stringTemp == "56789" || stringTemp == "456789" ||
			stringTemp == "3456789" || stringTemp == "23456789" ||
			stringTemp == "123456789" {
			printStringInRune(stringTemp)
		} else {
			// fmt.Println("printing stringTemp:", stringTemp)
			printStringInRune(stringTemp)
			z01.PrintRune(',')
			z01.PrintRune(' ')
			return
		}
	}
	for a := null; a <= 9; a++ {
		// fmt.Println("a:", a)
		digit := MyItoa(a)
		// fmt.Println("digit:", digit)
		stringOut = stringTemp + digit
		// fmt.Println("stringTemp:", stringTemp)
		// fmt.Println("stringOut:", stringOut)
		// fmt.Println("sending this to recursion", "a+1:", a+1, "n-1", n-1, "stringOut:", stringOut)
		RecursionFunc(a+1, n-1, stringOut)
	}
}

// printStringInRune Function to print string in rune
func printStringInRune(stringIn string) {
	for _, r := range stringIn {
		z01.PrintRune(r)
	}
}
