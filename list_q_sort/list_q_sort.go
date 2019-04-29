package main

import "fmt"

type node struct {
	data int
	pre  *node
	next *node
}

var A = []int{3, 1, 4, 1, 5, 9, 2, 6}

func insert(r **node, d int) {
	if (*r) == nil {
		(*r) = new(node)
		(*r).data = d
		return
	}

	for (*r).next != nil {
		r = &(*r).next
	}
	tmp := &node{data: d}
	tmp.next = nil
	tmp.pre = (*r)
	(*r).next = tmp
}

func partition(l *node, r *node) (*node, *node, *node) {
	if l == nil || r == nil || l == r {
		return l, nil, nil
	}

	var head, tail *node
	provt := l
	l = l.next

	for l != nil {
		next := l.next
		if l.data <= provt.data {

			if l.pre != nil {
				l.pre.next = l.next
			}
			if l.next != nil {
				l.next.pre = l.pre
			}

			pre := provt.pre
			if pre != nil {
				pre.next = l
			}
			l.pre = pre
			l.next = provt
			provt.pre = l
			if head == nil {
				head = l
			}

		} else {
			tail = l
		}
		if l == r {
			break
		}
		l = next

	}

	return head, tail, provt
}

func q_sort(l *node, r *node) *node {
	h, t, provt := partition(l, r)
	if provt == nil {
		return l
	}
	head := q_sort(h, provt.pre)
	q_sort(provt.next, t)
	return head
}

func main() {
	var root *node
	for _, k := range A {
		insert(&root, k)
	}
	var tail *node
	r := root
	for r != nil {
		//	fmt.Print(r.data, " ")
		if r.next == nil {
			tail = r
		}
		r = r.next
	}

	t := tail
	for t != nil {
		//	fmt.Print(t.data, " ")
		t = t.pre
	}


	head := q_sort(root, tail)

	r = head
	for r != nil {
		fmt.Print(r.data, " ")
		if r.next == nil {
			tail = r
		}
		r = r.next
	}
}
