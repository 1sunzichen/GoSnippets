package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(u int) *ListNode {
	return &ListNode{
		Val:  u,
		Next: nil,
	}
}

func insertNode(no *ListNode, p *ListNode) {
	n1 := no.Next
	p.Next = n1
	no.Next = p
}

func removeItem(no *ListNode) {
	if no.Next == nil {
		return
	}
	p := no.Next
	n1 := p.Next
	no.Next = n1
}

func access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}
func findNode(head *ListNode, target int) int {
	index := 0
	for head != nil {
		if head.Val == target {
			return index
		}
		head = head.Next
		index++
	}
	return -1
}
