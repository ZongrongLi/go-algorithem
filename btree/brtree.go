package main

import "fmt"

type node struct {
	data int
	l    *node
	r    *node
}

var A = []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7}

func insert(r **node, d int) {
	if (*r) == nil {
		(*r) = new(node)
		(*r).data = d
		return
	}
	if d < (*r).data {
		insert(&(*r).l, d)
	} else {
		insert(&(*r).r, d)
	}
}

func midorder(r *node) {
	if r == nil {
		return
	}

	midorder(r.l)
	fmt.Print(r.data, " ")
	midorder(r.r)
}

func main() {

	var root *node
	for _, k := range A {
		insert(&root, k)
	}

	midorder(root)
}
