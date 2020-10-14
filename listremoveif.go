package student

//NodeL lol
type NodeL struct {
	Data interface{}
	Next *NodeL
}

//List lol
type List struct {
	Head *NodeL
	Tail *NodeL
}

//ListPushBack lol
func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = n

}

//ListRemoveIf lol
func ListRemoveIf(l *List, data_ref interface{}) {
	newlist := &List{}
	it := l.Head
	for it != nil {
		if it.Data != data_ref {
			ListPushBack(newlist, it.Data)
		}
		it = it.Next
	}
	*l = *newlist
}
