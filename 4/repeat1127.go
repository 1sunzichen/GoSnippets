package main

import (
	"errors"
)

type ListNode1127 struct {
	Next *ListNode1127
	Val  int
}

func NewListNode1127(val int) *ListNode1127 {
	return &ListNode1127{
		Val:  val,
		Next: nil,
	}
}

func insertNode1127(n0 *ListNode1127, p *ListNode1127) error {
	if n0 == nil {
		return errors.New("没有信息")
	}
	n1 := n0.Next
	n0.Next = p
	p.Next = n1
	return nil
}

func findNode1127(head *ListNode1127, Val int) (index int) {

	if head != nil {
		if head.Val == Val {
			return index
		}
		head = head.Next
		index++
	}
	return -1
}

func accessNode1127(head *ListNode1127, index int) *ListNode1127 {
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}
