package main

import (
	student ".."
	"fmt"
)

func main() {

	arg1 := -7
	arg2 := -2
	fmt.Println(student.RecursivePower(arg1, arg2))
	arg1 = -8
	arg2 = -7
	fmt.Println(student.RecursivePower(arg1, arg2))
	arg1 = 4
	arg2 = 8
	fmt.Println(student.RecursivePower(arg1, arg2))
	arg1 = 1
	arg2 = 3
	fmt.Println(student.RecursivePower(arg1, arg2))
	arg1 = -1
	arg2 = 1
	fmt.Println(student.RecursivePower(arg1, arg2))
	arg1 = -6
	arg2 = 5
	fmt.Println(student.RecursivePower(arg1, arg2))

}

// 0
// 0
// 65536
// 1
// -1
// -7776
