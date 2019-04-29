package main

import "fmt"

var A = []int{5, 3, 4, 2, 1, 7, 6, 8, 9}

type node struct {
	data int
	l    *node
	r    *node
}

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

type qnode struct {
	data *node
	next *qnode
}

type queue struct {
	head *qnode
	tail *qnode
}

func (q *queue) push_back(d *node) {
	n := &qnode{data: d}
	if q.head == nil {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		q.tail = n
	}
}
func (q *queue) get_back() *node {
	return q.tail.data
}

func (q *queue) pop_front() *node {
	if q.head == nil {
		return nil
	}
	d := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return d
}

func (q *queue) empty() bool {
	return q.head == nil
}

var q queue

var level = 0

func levelorder(r *node) {
	q.push_back(r)
	last := r
	fmt.Println("=====", r.data)

	for !q.empty() {
		r := q.pop_front()

		fmt.Print(r.data, " ")
		if r.l != nil {
			q.push_back(r.l)
		}
		if r.r != nil {
			q.push_back(r.r)
		}

		if r == last && !q.empty() {
			last = q.get_back()
			fmt.Println("=====", last.data)
		}
	}

}

func main() {
	var root *node
	for _, k := range A {
		insert(&root, k)
	}
	//midorder(root)
	levelorder(root)

}
