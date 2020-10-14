package student

// SortedListMerge lol
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	var n3 *NodeI
	iter1 := n1
	iter2 := n2
	for {
		if iter1 != nil {
			n3 = listPushBack(n3, iter1.Data)
			iter1 = iter1.Next
			continue
		} else if iter2 != nil {
			n3 = listPushBack(n3, iter2.Data)
			iter2 = iter2.Next
			continue
		} else if iter2 == nil && iter1 == nil {
			break
		}
	}
	iter1 = n3
	for iter1 != nil {
		iter2 = n3
		for iter2 != nil {
			if iter1.Data < iter2.Data {
				iter1.Data, iter2.Data = iter2.Data, iter1.Data
			}
			iter2 = iter2.Next
		}
		iter1 = iter1.Next
	}
	return n3
}
