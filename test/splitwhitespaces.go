package main

import (
	student ".."
	"fmt"
)

func main() {

	str := "The earliest foundations of what would become   computer science predate the invention of the modern digital computer"
	fmt.Println(student.SplitWhiteSpaces(str))
	str = "Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"
	fmt.Println(student.SplitWhiteSpaces(str))
	str = "aiding in  computations such as multiplication and division ."
	fmt.Println(student.SplitWhiteSpaces(str))
	str = "Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."
	fmt.Println(student.SplitWhiteSpaces(str))
	str = "Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"
	fmt.Println(student.SplitWhiteSpaces(str))
	str = "In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"
	fmt.Println(student.SplitWhiteSpaces(str))

}

// [The earliest foundations of what would become computer science predate the invention of the modern digital computer]
// [Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,]
// [aiding in computations such as multiplication and division .]
// [Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment.]
// [Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]]
// [In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,]
