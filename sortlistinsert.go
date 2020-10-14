package student

// NodeI a
type NodeI struct {
	Data int
	Next *NodeI
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

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

//SortListInsert l
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	it := l
	flag := false
	var link *NodeI
	for it != nil {
		if data_ref < it.Data && !flag {
			link = listPushBack(link, data_ref)
			link = listPushBack(link, it.Data)
			flag = true
		} else if it.Next == nil && !flag {
			link = listPushBack(link, it.Data)
			link = listPushBack(link, data_ref)
		} else {
			link = listPushBack(link, it.Data)
		}

		it = it.Next
	}
	return link
}
