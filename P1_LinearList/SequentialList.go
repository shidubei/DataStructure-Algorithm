package P1_LinearList

import "fmt"

const MaxSize = 100
const InitSize = 10

/*SequentialList1 static*/
type SequentialList1 struct {
	SqList [MaxSize]int
	length int
}

// InitList init
func (s *SequentialList1) InitList() {
	for i := 0; i < len(s.SqList); i++ {
		s.SqList[i] = 0
	}
	s.length = 0
}

// ListInsert insert
func (s *SequentialList1) ListInsert(index, value int) (bool, error) {
	if index > len(s.SqList) || index < 1 {
		return false, fmt.Errorf("index out of range")
	}
	for i := index; i < len(s.SqList); i++ {
		s.SqList[i] = s.SqList[i-1]
	}
	s.SqList[index-1] = value
	s.length++
	return true, nil
}

// ListDelete delete
func (s *SequentialList1) ListDelete(index int) (bool, error) {
	if index > len(s.SqList) || index < 1 {
		return false, fmt.Errorf("index out of range")
	}
	for i := index - 1; i < len(s.SqList)-1; i++ {
		s.SqList[i] = s.SqList[i+1]
	}
	s.length--
	return true, nil
}

// GetElem return the element at index of the List
func (s SequentialList1) GetElem(index int) (int, error) {
	if index > len(s.SqList) || index < 1 {
		return 0, fmt.Errorf("index out of range")
	}
	return s.SqList[index-1], nil
}

// LocateElem return the index
func (s SequentialList1) LocateElem(value int) (int, error) {
	for i := 0; i < len(s.SqList); i++ {
		if s.SqList[i] == value {
			return i, nil
		}
	}
	return -1, nil
}

/*SequentialList2 dynamic*/
type SequentialList2 struct {
	// In go, use slice as dynamic list
	SqList  []int
	length  int
	maxSize int
}

func (s *SequentialList2) InitList() {
	s.SqList = make([]int, InitSize)
	s.length = 0
	s.maxSize = InitSize
}

func (s *SequentialList2) IncreaseSize(addLen int) {
	// use temp to copy s.SqList temporarily
	temp := s.SqList
	s.SqList = make([]int, addLen+s.maxSize)
	copy(s.SqList, temp)
	s.maxSize = addLen + s.maxSize
	// release the temp
	temp = nil
}

func (s *SequentialList2) ListInsert(index, value int) (bool, error) {
	if index > s.maxSize {
		s.IncreaseSize(index - s.maxSize + 1)
	}
	if index < 1 {
		return false, fmt.Errorf("index should be at least 1")
	}
	for i := index; i < len(s.SqList); i++ {
		s.SqList[i] = s.SqList[i-1]
	}
	s.SqList[index-1] = value
	s.length++
	return true, nil
}

func (s *SequentialList2) ListDelete(index int) (bool, error) {
	if index < 1 || index > s.maxSize {
		return false, fmt.Errorf("index out of range")
	}
	for i := index - 1; i < len(s.SqList); i++ {
		s.SqList[i] = s.SqList[i+1]
	}
	s.length--
	return true, nil
}

func (s SequentialList2) GetElem(index int) (int, error) {
	if index > len(s.SqList) || index < 1 {
		return 0, fmt.Errorf("index out of range")
	}
	return s.SqList[index-1], nil
}

func (s SequentialList2) LocateElem(value int) (int, error) {
	for i := 0; i < len(s.SqList); i++ {
		if s.SqList[i] == value {
			return i, nil
		}
	}
	return -1, nil
}
