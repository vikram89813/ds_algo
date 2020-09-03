package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func inorder(root *node) {
	if root != nil {
		inorder(root.left)
		fmt.Printf("%v ", root.val)
		inorder(root.right)
	}
}

func predSucc(root *node, v int, prev **node, next **node) {
	if root != nil {
		predSucc(root.left, v, prev, next)
		if root != nil && root.val > v {
			if (*next == nil) || root.val < (*next).val {
				*next = root
			}
		} else if root != nil && root.val < v {
			*prev = root
		}
		predSucc(root.right, v, prev, next)
	}
}

func main() {
	root := &node{50, nil, nil}
	root.left = &node{20, nil, nil}
	root.right = &node{60, nil, nil}
	root.left.left = &node{10, nil, nil}
	root.left.right = &node{30, nil, nil}
	root.right.left = &node{55, nil, nil}
	root.right.right = &node{70, nil, nil}

	inorder(root)
	fmt.Println("")
	var prev *node
	var next *node
	prev = nil
	next = nil
	predSucc(root, 55, &prev, &next)
	fmt.Printf("%v,%v\n", prev.val, next.val)
}
