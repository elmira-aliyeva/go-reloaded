package main

import (
	"fmt"

	student ".."
)

func main() {

	root := &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root, "07")
	student.BTreeInsertData(root, "12")
	student.BTreeInsertData(root, "10")
	student.BTreeInsertData(root, "05")
	student.BTreeInsertData(root, "02")
	student.BTreeInsertData(root, "03")
	student.BTreeApplyByLevel(root, fmt.Println)

	// 01
	// 07
	// 05
	// 12
	// 02
	// 10
	// 03

	// root := &student.TreeNode{Data: "01"}
	// student.BTreeInsertData(root, "07")
	// student.BTreeInsertData(root, "12")
	// student.BTreeInsertData(root, "10")
	// student.BTreeInsertData(root, "05")
	// student.BTreeInsertData(root, "02")
	// student.BTreeInsertData(root, "03")
	// student.BTreeApplyByLevel(root, fmt.Print)

	// 01070512021003

}

// btreeapplybylevel
// Try with the arguments:
// root =
// 01
// └── 07
//     ├── 12
//     │   └── 10
//     └── 05
//         └── 02
//             └── 03
// and f = fmt.Println

// Does the function prints the value above?
// Try with the arguments:
// root =
// 01
// └── 07
//     ├── 12
//     │   └── 10
//     └── 05
//         └── 02
//             └── 03
// and f = fmt.Print
// 01070512021003
// Does the function prints the value above?
