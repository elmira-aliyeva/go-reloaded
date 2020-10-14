package main

import (
	"fmt"

	student ".."
)

func main() {
	// root := &student.TreeNode{Data: "01"}
	// student.BTreeInsertData(root, "07")
	// student.BTreeInsertData(root, "05")
	// student.BTreeInsertData(root, "12")
	// student.BTreeInsertData(root, "02")
	// student.BTreeInsertData(root, "10")
	// student.BTreeInsertData(root, "03")
	// node := student.BTreeSearchItem(root, "02")
	// fmt.Println("Before delete:")
	// student.BTreeApplyInorder(root, fmt.Println)
	// root = student.BTreeDeleteNode(root, node)
	// fmt.Println("After delete:")
	// student.BTreeApplyInorder(root, fmt.Println)

	// 01
	// 03
	// 05
	// 07
	// 10
	// 12

	// root := &student.TreeNode{Data: "33"}
	// student.BTreeInsertData(root, "5")
	// student.BTreeInsertData(root, "20")
	// student.BTreeInsertData(root, "31")
	// student.BTreeInsertData(root, "13")
	// student.BTreeInsertData(root, "11")
	// student.BTreeInsertData(root, "52")
	// node := student.BTreeSearchItem(root, "20")
	// fmt.Println("Before delete:")
	// student.BTreeApplyInorder(root, fmt.Println)
	// root = student.BTreeDeleteNode(root, node)
	// fmt.Println("After delete:")
	// student.BTreeApplyInorder(root, fmt.Println)

	// 11
	// 13
	// 31
	// 33
	// 5
	// 52

	// root := &student.TreeNode{Data: "03"}
	// student.BTreeInsertData(root, "39")
	// student.BTreeInsertData(root, "99")
	// student.BTreeInsertData(root, "44")
	// student.BTreeInsertData(root, "11")
	// student.BTreeInsertData(root, "14")
	// student.BTreeInsertData(root, "11")
	// node := student.BTreeSearchItem(root, "03")
	// fmt.Println("Before delete:")
	// student.BTreeApplyInorder(root, fmt.Println)
	// root = student.BTreeDeleteNode(root, node)
	// fmt.Println("After delete:")
	// student.BTreeApplyInorder(root, fmt.Println)

	// 11
	// 11
	// 14
	// 39
	// 44
	// 99

	root := &student.TreeNode{Data: "03"}
	student.BTreeInsertData(root, "03")
	student.BTreeInsertData(root, "94")
	student.BTreeInsertData(root, "19")
	student.BTreeInsertData(root, "24")
	student.BTreeInsertData(root, "111")
	student.BTreeInsertData(root, "01")
	node := student.BTreeSearchItem(root, "03")
	fmt.Println("Before delete:")
	student.BTreeApplyInorder(root, fmt.Println)
	root = student.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	student.BTreeApplyInorder(root, fmt.Println)

	// 01
	// 03
	// 111
	// 19
	// 24
	// 94
}

// root =
// 01
// └── 07
//     ├── 12
//     │   └── 10
//     └── 05
//         └── 02
//             └── 03
// and node = 02
// 01
// └── 07
//     ├── 12
//     │   └── 10
//     └── 05
//         └── 03
// Does the function returns the value above?
// Try with the arguments:
// root =
// 33
// ├── 5
// │   └── 52
// └── 20
//     ├── 31
//     └── 13
//         └── 11
// and node = 20
// 33
// ├── 5
// │   └── 52
// └── 31
//     └── 13
//         └── 11
// Does the function returns the value above?
// Try with the arguments:
// root =
// 03
// └── 39
//     ├── 99
//     │   └── 44
//     └── 11
//         └── 14
//             └── 11
//  and node = 03
// 39
// ├── 99
// │   └── 44
// └── 11
//     └── 14
//         └── 11
// Does the function returns the value above?
// Try with the arguments:
// root =
// 03
// ├── 03
// │   └── 94
// │       └── 19
// │           ├── 24
// │           └── 111
// └── 01
// and node = 03
// 03
// ├── 94
// │   └── 19
// │       ├── 24
// │       └── 111
// └── 01
