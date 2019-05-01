package main

import "errors"

type node struct {
	data int
	prev *node
	next *node
}

type lrucache struct {
	kv       map[int]*node
	capicity int
	count    int
	head     *node
	tail     *node
}

func (l *lrucache) Set(k, v int) {
	p, ok := l.kv[k]
	if ok {
		(*p).data = v
		return
	}

	n := node{
		data: v,
	}
	if l.head == nil {
		l.head = &n
		l.tail = &n
		l.count = 1
		return
	}
	l.tail.next = &n
	n.prev = l.tail
	l.tail = &n
	l.count++

	if l.count >= l.capicity {
		tmp := l.tail
		l.tail = l.tail.prev
		tmp.prev = nil
		l.tail.next = nil
		l.count--
	}
}

func (l *lrucache) Get(k int) (int, error) {
	p, ok := l.kv[k]
	if !ok {
		return -1, errors.New("wrong param")
	}
	if p.prev != nil {
		p.prev.next = p.next
	}
	if p.next != nil {
		p.next.prev = p.prev
	}

	tmp := l.head
	tmp.prev = p
	p.next = tmp
	l.head = p

	return p.data, nil
}

func (l *lrucache) Delete(k int) {
	p, ok := l.kv[k]
	if !ok {
		return
	}
	delete(l.kv, k)
	l.count--
	if p.prev != nil {
		p.prev.next = p.next
	} else {
		l.head = p.next
	}
	if p.next != nil {
		p.next.prev = p.prev
	} else {
		l.tail = p.prev
	}
}

func main() {

}
