package main

import "fmt"

type node struct {
	val  int
	next *node
}

func insert(head *node, v int) *node {
	if head == nil {
		head = &node{v, nil}
	} else {
		t := &node{v, head}
		head = t
	}
	return head
}

func (head *node) String() string {
	for head != nil {
		fmt.Printf("%v ", head.val)
		head = head.next
	}
	return ""
}

func FindMiddle(head *node) *node {
	slowPtr := head
	fastPtr := head

	for fastPtr != nil && fastPtr.next != nil {
		fastPtr = fastPtr.next.next
		slowPtr = slowPtr.next
	}

	return slowPtr
}

func Rev(head *node) *node {
	curr := head
	var prev *node
	var next *node

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	head = prev

	return head
}

func main() {
	var head *node
	for i := 9; i >= 1; i-- {
		head = insert(head, i)
	}

	fmt.Println(head)
	fmt.Println(FindMiddle(head).val)
	fmt.Println(Rev(head))
}
