package linkedlist

import (
	"errors"
	"fmt"
	"sync"
)

type Node struct {
	Data interface{}
	Next *Node
}

type SingleList struct {
	head 	*Node
	size	int
	mu 		sync.RWMutex
}

func GenSingleList(len int) (s *SingleList) {
	s = &SingleList{}
	for i:=0; i < len; i++ {
		s.Append(Node{ Data: i})
	}
	return
}

func (s *SingleList) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.head == nil {
		return true
	}
	return false
}

func (s *SingleList) Get(index int) *Node {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if index < 0 || index > s.size {
		return nil
	}
	n := s.head
	for j := 0; j < index-1;j++ {
		n = n.Next
	}
	return n
}

func (s *SingleList) Append(n Node) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.head != nil {
		last := s.head
		for last.Next != nil {
			last = last.Next
		}
		last.Next = &n
	} else {
		s.head = &n
	}
	s.size++
}

func (s *SingleList) String()  {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.head == nil {
		return
	}
	n := s.head
	for {
		if n == nil {
			break
		}
		fmt.Printf("%+v\n",n.Data)
		n = n.Next
	}
}

func (s *SingleList) Remove(i int) (*Node,error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if i < 0 || i > s.size {
		return nil, errors.New("index out of bounds")
	}
	n := s.head
	for j:= 0; j < i-1;j++{
		n = n.Next
	}
	remove := n.Next
	n.Next = remove.Next
	s.size--
	return remove,nil
}

func (s *SingleList) Reverse() {
	if s.head == nil || s.head.Next == nil {
		return
	}
	head := s.head
	var prev *Node
	for head != nil {
		head,prev,head.Next = head.Next, head,prev
	}
	s.head = prev
}
