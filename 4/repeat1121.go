package main

type ListNode1122 struct {
	Val  int
	Next *ListNode1122
}

func NewListNode1122(i int) *ListNode1122 {
	return &ListNode1122{
		Val:  i,
		Next: nil,
	}
}

func insertNode1122(no *ListNode1122, p0 *ListNode1122) {
	n1 := no.Next
	p0.Next = n1
	no.Next = p0
}

func findNode1122(head *ListNode1122, Val int) (index int) {

	for head != nil {
		if head.Val == Val {
			return index
		}
		head = head.Next
		index++
	}
	return -1
}

func accessNode1122(head *ListNode1122, index int) *ListNode1122 {
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}
