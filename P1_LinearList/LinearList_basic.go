package P1_LinearList

import "fmt"

// the fundamental operations for LinearList

type LinearList struct {
	list []int
}

// InitList init a LinearList
func (l *LinearList) InitList(size int) []int {
	l.list = make([]int, size)
	return l.list
}

// DestroyList destroy a LinearList
func (l *LinearList) DestroyList() {
	l.list = nil
}

// ListInsert insert one value into the LinearList's index
func (l *LinearList) ListInsert(index, value int) {
	// check the index whether out of the LinearList's boundary
	if index >= len(l.list) {
		return
	}
	for i := index + 1; i < len(l.list); i++ {
		l.list[i] = l.list[i-1]
	}
	l.list[index] = value
}

// ListDelete delete one value by the LinearList's index
func (l *LinearList) ListDelete(index int) {
	//  check the index whether out of the LinearList's boundary
	if index >= len(l.list) {
		return
	}
	if index == len(l.list)-1 {
		l.list = l.list[:index]
	}
	behind := l.list[index+1:]
	l.list = l.list[:index]
	l.list = append(l.list, behind...)
}

// LocateElem return the first element's index in the LinearList
func (l LinearList) LocateElem(e int) int {
	for i := 0; i < len(l.list); i++ {
		if l.list[i] == e {
			return i
		}
	}
	// return -1 means no such element
	return -1
}

// GetElem return the element by the index of the LinearList
func (l LinearList) GetElem(index int) (int, error) {
	if index >= len(l.list) {
		return 0, fmt.Errorf("index out of bonudary")
	}
	return l.list[index], nil
}

// Length get the length of the LinearList
func (l LinearList) Length() int {
	return len(l.list)
}

// PrintList print the LinearList
func (l LinearList) PrintList() {
	fmt.Println(l.list)
}

// Empty whether the LinearList is empty or not
func (l LinearList) Empty() bool {
	if len(l.list) == 0 {
		// false means the LinearList is empty
		return false
	}
	// true means the LinearList is not empty
	return true
}
