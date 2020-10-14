package main

import (
	"fmt"

	student ".."
)

//List lol
type List = student.List

//Node lol
type Node = student.NodeL

//PrintList lol
func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " , ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {

	var link *student.NodeI

	// 1

	// link = listPushBack(link, 0)

	// PrintList(link)

	// link = student.SortListInsert(link, 39)
	// PrintList(link)

	// ry with the arguments: l = 0, <nil> and data_ref = 39
	// (0, 39, <nil>)

	// 2

	// link = listPushBack(link, 0)
	// link = listPushBack(link, 1)
	// link = listPushBack(link, 2)
	// link = listPushBack(link, 3)
	// link = listPushBack(link, 4)
	// link = listPushBack(link, 5)
	// link = listPushBack(link, 24)
	// link = listPushBack(link, 25)
	// link = listPushBack(link, 54)

	// PrintList(link)

	// link = student.SortListInsert(link, 33)
	// PrintList(link)

	// Does the function returns the value above?
	// Try with the arguments: l = 0, 1, 2, 3, 4, 5, 24, 25, 54, <nil> and data_ref = 33
	// (0, 1, 2, 3, 4, 5, 24, 25, 33, 54, <nil>)

	// 3

	// link = listPushBack(link, 0)
	// link = listPushBack(link, 2)
	// link = listPushBack(link, 18)
	// link = listPushBack(link, 33)
	// link = listPushBack(link, 37)
	// link = listPushBack(link, 37)
	// link = listPushBack(link, 39)
	// link = listPushBack(link, 52)
	// link = listPushBack(link, 53)
	// link = listPushBack(link, 57)

	// PrintList(link)

	// link = student.SortListInsert(link, 53)
	// PrintList(link)

	// Does the function returns the value above?
	// Try with the arguments: l = 0, 2, 18, 33, 37, 37, 39, 52, 53, 57, <nil> and data_ref = 53
	// (0, 2, 18, 33, 37, 37, 39, 52, 53, 53, 57, <nil>)

	// 4

	// link = listPushBack(link, 0)
	// link = listPushBack(link, 5)
	// link = listPushBack(link, 18)
	// link = listPushBack(link, 24)
	// link = listPushBack(link, 28)
	// link = listPushBack(link, 35)
	// link = listPushBack(link, 42)
	// link = listPushBack(link, 45)
	// link = listPushBack(link, 52)

	// PrintList(link)

	// link = student.SortListInsert(link, 52)
	// PrintList(link)

	// Does the function returns the value above?
	// Try with the arguments: l = 0, 5, 18, 24, 28, 35, 42, 45, 52, <nil> and data_ref = 52
	// (0, 5, 18, 24, 28, 35, 42, 45, 52, 52, <nil>)

	// 5

	// link = listPushBack(link, 0)
	// link = listPushBack(link, 12)
	// link = listPushBack(link, 20)
	// link = listPushBack(link, 23)
	// link = listPushBack(link, 23)
	// link = listPushBack(link, 24)
	// link = listPushBack(link, 30)
	// link = listPushBack(link, 41)
	// link = listPushBack(link, 53)
	// link = listPushBack(link, 57)
	// link = listPushBack(link, 59)

	// PrintList(link)

	// link = student.SortListInsert(link, 38)
	// PrintList(link)

	// Does the function returns the value above?
	// Try with the arguments: l = 0, 12, 20, 23, 23, 24, 30, 41, 53, 57, 59, <nil> and data_ref = 38
	// (0, 12, 20, 23, 23, 24, 30, 38, 41, 53, 57, 59, <nil>)
}
