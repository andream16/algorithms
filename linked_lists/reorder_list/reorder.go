package main

import "fmt"

func main()  {
	l := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
				},
			},
		},
	}
	reorderList(l)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode)  {

	l := &ListNode{}
	nextNode := &ListNode{}

	for head != nil {
		l.Val = head.Val
		fmt.Println(l.Val)
		nextNode = head.Next
		for nextNode.Next != nil {
			nextNode = nextNode.Next
		}
		head = head.Next
		l.Next = nextNode
		nextNode = nil
	}

	for l.Next != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

}
