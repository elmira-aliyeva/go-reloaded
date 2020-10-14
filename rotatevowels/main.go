package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	len := 0
	for range args {
		len++
	}
	if len > 0 {
		vowels := "AEIOUaeiou"
		str, vow, vow1 := "", "", ""
		for i, j := range args {
			str += j
			if i != len-1 {
				str += " "
			}
			for _, b := range j {
				for _, c := range vowels {
					if b == c {
						vow += string(b)
					}
				}
			}
		}
		var lenv int
		for range vow {
			lenv++
		}
		for i := lenv - 1; i >= 0; i-- {
			vow1 += string(vow[i])
		}
		a, res := 0, ""
		for _, j := range str {
			vowel := false
			for _, b := range vowels {
				if j == b {
					res += string(vow1[a])
					a++
					vowel = true
					break
				}
			}
			if !vowel {
				res += string(j)
			}
		}
		for _, j := range res {
			z01.PrintRune(j)
		}
	}
	z01.PrintRune('\n')
}


// ./rotatevowels "Hello World"
// Hollo Werld
// ./rotatevowels "HEllO World" "problem solved"
// Hello Werld problom sOlvEd
//  ./rotatevowels "str" "shh" "psst"
// str shh psst
// ./rotatevowels "happy thoughts" "good luck"
// huppy thooghts guod lack
// ./rotatevowels "al's elEphAnt is overly underweight!"
// il's elephunt es ovirly AndErweaght!