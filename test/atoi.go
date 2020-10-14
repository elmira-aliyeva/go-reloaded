package main

import (
	"fmt"

	student ".."
)

func main() {

	s9 := ""
	s10 := "-"
	s11 := "--123"
	s12 := "1"
	s13 := "-3"
	s14 := "8292"
	s15 := "9223372036854775807"
	s16 := "-9223372036854775808"

	n9 := student.Atoi(s9)
	n10 := student.Atoi(s10)
	n11 := student.Atoi(s11)
	n12 := student.Atoi(s12)
	n13 := student.Atoi(s13)
	n14 := student.Atoi(s14)
	n15 := student.Atoi(s15)
	n16 := student.Atoi(s16)

	fmt.Println(n9)
	fmt.Println(n10)
	fmt.Println(n11)
	fmt.Println(n12)
	fmt.Println(n13)
	fmt.Println(n14)
	fmt.Println(n15)
	fmt.Println(n16)
}

// 0
// 0
// 0
// 1
// -3
// 8292
// 9223372036854775807
// -9223372036854775808
