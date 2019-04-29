package main

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
	pre  *node
}

type lrucache struct {
	head     *node
	tail     *node
	kv       map[string]*node
	capicity int
	count    int
}

func (l *lrucache) Set(k string, v int) {
	if n, ok := l.kv[k]; ok {
		n.data = v
		return
	}
	tmp := &node{data: v, next: nil, pre: nil}
	if l.tail == nil {
		l.head = tmp
		l.tail = tmp
	} else {
		l.tail.next = tmp
		tmp.pre = l.tail
		l.tail = tmp
	}
	l.count++

	if l.count > l.capicity {
		l.tail = l.tail.pre
		l.count--
	}

}

func (l *lrucache) Get(k string) (int, error) {
	if n, ok := l.kv[k]; ok {
		if n.pre != nil {
			n.pre.next = n.next
		}
		if n.next != nil {
			n.next.pre = n.pre
		}

		n.next = l.head
		if l.head != nil {
			l.head.pre = n
		}
		l.head = n

		return n.data, nil
	}
	return -1, errors.New("no data")
}

func (l *lrucache) Delete(k string) {
	if n, ok := l.kv[k]; ok && n != nil {
		if n.pre != nil {
			n.pre.next = n.next
		}

		if n.next != nil {
			n.next.pre = n.pre
		}
		l.kv[k] = nil
		delete(l.kv, k)
		return
	}
}

func main() {
	s := make(map[string]*node)
	s["123"] = &node{data: 10}
	s["123"] = nil
	if _, ok := s["123"]; ok {
		fmt.Println("ok")
	}
	fmt.Println("bu ok")
}
