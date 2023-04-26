package P1_LinearList

type LinkList struct {
	Val  int
	Next *LinkList
}

// InitList init a LinkList without headLink
func (l *LinkList) InitList() bool {
	l = nil
	return true
}

// InitListHead init a LinkList with headLink
func (l *LinkList) InitListHead() bool {
	if l == nil {
		return false
	}
	l.Next = nil
	return true
}

// LintInsert insert in the index of the LinkList without headLink
func (l *LinkList) LintInsert(index, value int) bool {
	// if l == nil,means there have no val in LinkList,so just set the head
	if l == nil {
		l.Val = value
		l.Next = nil
		return true
	}
	cur := l
	i := 1
	for cur != nil && i < index {
		cur = cur.Next
		i++
	}
	if cur == nil {
		return false
	}
	temp := &LinkList{Val: value}
	temp.Next = cur.Next
	cur.Next = temp
	return true
}
