package main

import (
	student ".."
	"fmt"
)

func main() {

	result := student.ConvertBase("4506C", "0123456789ABCDEF", "choumi")
	fmt.Println(result)
	result = student.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF")
	fmt.Println(result)
	result = student.ConvertBase("256850", "0123456789", "01")
	fmt.Println(result)
	result = student.ConvertBase("uuhoumo", "choumi", "Zone01")
	fmt.Println(result)
	result = student.ConvertBase("683241", "0123456789", "0123456789")
	fmt.Println(result)

}

// "hccocimc"
// "9B611"
// "111110101101010010"
// "eeone0n"
// "683241"
