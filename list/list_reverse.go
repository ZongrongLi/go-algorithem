package main

import "fmt"

type node struct {
	data int
	next *node
}

var A = []int{1, 2, 3, 4, 5, 6, 7}

func insert(r **node, d int) {
	if (*r) == nil {
		(*r) = new(node)
		(*r).data = d
		return
	}
	insert(&(*r).next, d)
}

func headinsert(head *node, d *node) *node {
	if head == nil {
		d.next = nil
		return d
	}
	d.next = head
	return d
}

func reverse(r *node) *node {
	var h *node
	for r != nil {
		next := r.next
		h = headinsert(h, r)
		r = next
	}
	return h
}

func main() {
	var root *node
	for _, k := range A {
		insert(&root, k)
	}

	root = reverse(root)

	for root != nil {
		fmt.Print(root.data, " ")
		root = root.next
	}

}
