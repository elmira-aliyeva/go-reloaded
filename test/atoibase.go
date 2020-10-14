package main

import (
	student ".."
	"fmt"
)

func main() {

	fmt.Println(student.AtoiBase("bcbbbbaab", "abc"))
	fmt.Println(student.AtoiBase("0001", "01"))
	fmt.Println(student.AtoiBase("00", "01"))
	fmt.Println(student.AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	fmt.Println(student.AtoiBase("AAho?Ao", "WhoAmI?"))
	fmt.Println(student.AtoiBase("thisinputshouldnotmatter", "abca"))
	fmt.Println(student.AtoiBase("125", "0123456789"))
	fmt.Println(student.AtoiBase("uoi", "choumi"))
	fmt.Println(student.AtoiBase("bbbbbab", "-ab"))

}

// 12016
// 1
// 0
// 11557277891
// 406772
// 0
// 125
// 125
// 0
