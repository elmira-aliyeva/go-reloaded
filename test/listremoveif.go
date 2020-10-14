package main

import (
	"fmt"

	student ".."
)

//PrintList lo
func PrintList(l *student.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
	a := &student.List{}

	fmt.Println("----normal state----")
	PrintList(a)
	student.ListRemoveIf(a, 1)
	fmt.Println("------answer-----")
	PrintList(a)
	fmt.Println()

	b := &student.List{}

	fmt.Println("----normal state----")
	PrintList(b)
	student.ListRemoveIf(b, 96)
	fmt.Println("------answer-----")
	PrintList(b)
	fmt.Println()

	c := &student.List{}

	fmt.Println("----normal state----")

	student.ListPushBack(c, 98)
	student.ListPushBack(c, 98)
	student.ListPushBack(c, 33)
	student.ListPushBack(c, 34)
	student.ListPushBack(c, 33)
	student.ListPushBack(c, 34)
	student.ListPushBack(c, 33)
	student.ListPushBack(c, 89)
	student.ListPushBack(c, 33)

	PrintList(c)
	student.ListRemoveIf(c, 34)
	fmt.Println("------answer-----")
	PrintList(c)
	fmt.Println()

	d := &student.List{}

	fmt.Println("----normal state----")

	student.ListPushBack(d, 79)
	student.ListPushBack(d, 74)
	student.ListPushBack(d, 99)
	student.ListPushBack(d, 79)
	student.ListPushBack(d, 7)

	PrintList(d)
	student.ListRemoveIf(d, 99)
	fmt.Println("------answer-----")
	PrintList(d)
	fmt.Println()

	e := &student.List{}

	fmt.Println("----normal state----")

	student.ListPushBack(e, 56)
	student.ListPushBack(e, 93)
	student.ListPushBack(e, 68)
	student.ListPushBack(e, 56)
	student.ListPushBack(e, 87)
	student.ListPushBack(e, 68)
	student.ListPushBack(e, 56)
	student.ListPushBack(e, 68)

	PrintList(e)
	student.ListRemoveIf(e, 68)
	fmt.Println("------answer-----")
	PrintList(e)
	fmt.Println()

	f := &student.List{}

	fmt.Println("----normal state----")

	student.ListPushBack(f, "mvkUxbqhQve4l")
	student.ListPushBack(f, "4Zc4t hnf SQ")
	student.ListPushBack(f, "q2If E8BPuX")

	PrintList(f)
	student.ListRemoveIf(f, "4Zc4t hnf SQ")
	fmt.Println("------answer-----")
	PrintList(f)
	fmt.Println()
}

// (<nil>)
// (<nil>)
// (98-> 98-> 33-> 33-> 33-> 89-> 33-> <nil>)
// (79-> 74-> 79-> 7-> <nil>)
// (56-> 93-> 56-> 87-> 56-> <nil>)
// (mvkUxbqhQve4l-> q2If E8BPuX -> <nil>)
