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
	// student.BTreeInsertData(root, "10")
	// student.BTreeInsertData(root, "02")
	// student.BTreeInsertData(root, "03")
	// node := student.BTreeSearchItem(root, "12")
	// replacement := &student.TreeNode{Data: "55"}
	// student.BTreeInsertData(replacement, "60")
	// student.BTreeInsertData(replacement, "33")
	// student.BTreeInsertData(replacement, "12")
	// student.BTreeInsertData(replacement, "15")
	// root = student.BTreeTransplant(root, node, replacement)
	// student.BTreeApplyInorder(root, fmt.Println)

	// 01
	// 02
	// 03
	// 05
	// 07
	// 12
	// 15
	// 33
	// 55
	// 60

	// root := &student.TreeNode{Data: "33"}
	// student.BTreeInsertData(root, "5")
	// student.BTreeInsertData(root, "52")
	// student.BTreeInsertData(root, "20")
	// student.BTreeInsertData(root, "31")
	// student.BTreeInsertData(root, "13")
	// student.BTreeInsertData(root, "11")
	// node := student.BTreeSearchItem(root, "20")
	// replacement := &student.TreeNode{Data: "55"}
	// student.BTreeInsertData(replacement, "60")
	// student.BTreeInsertData(replacement, "33")
	// student.BTreeInsertData(replacement, "12")
	// student.BTreeInsertData(replacement, "15")
	// root = student.BTreeTransplant(root, node, replacement)
	// student.BTreeApplyInorder(root, fmt.Println)

	// 12
	// 15
	// 33
	// 55
	// 60
	// 33
	// 5
	// 52

	root := &student.TreeNode{Data: "03"}
	student.BTreeInsertData(root, "39")
	student.BTreeInsertData(root, "99")
	student.BTreeInsertData(root, "44")
	student.BTreeInsertData(root, "11")
	student.BTreeInsertData(root, "14")
	student.BTreeInsertData(root, "11")
	node := student.BTreeSearchItem(root, "11")
	replacement := &student.TreeNode{Data: "55"}
	student.BTreeInsertData(replacement, "60")
	student.BTreeInsertData(replacement, "33")
	student.BTreeInsertData(replacement, "12")
	student.BTreeInsertData(replacement, "15")
	root = student.BTreeTransplant(root, node, replacement)
	student.BTreeApplyInorder(root, fmt.Println)

	// 03
	// 12
	// 15
	// 33
	// 55
	// 60
	// 39
	// 44
	// 99
}

// btreetransplant
// Try with the arguments:
// root =
// 01
// └── 07
//     ├── 12
//     │   └── 10
//     └── 05
//         └── 02
//             └── 03
// , node = 12 and
// repla =
// 55
// ├── 60
// └── 33
//     └── 12
//         └── 15

// 01
// └── 07
//     ├── 55
//     │   ├── 60
//     │   └── 33
//     │       └── 12
//     │           └── 15
//     └── 05
//         └── 02
// 			└── 03

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
// , node = 20 and
// repla =
// 55
// ├── 60
// └── 33
//     └── 12
// 		└── 15

// 33
// ├── 5
// │   └── 52
// └── 55
//     ├── 60
//     └── 33
//         └── 12
//             └── 15
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
// , node = 11 and
// repla =
// 55
// ├── 60
// └── 33
//     └── 12
//         └── 15
// 03
// └── 39
//     ├── 99
//     │   └── 44
//     └── 55
//         ├── 60
//         └── 33
//             └── 12
//                 └── 15
