package linkedlist

import "sync"

type Node[T interface{}] struct {
	Data *T
	Prev *Node[T]
	next *Node[T]
}

type LinkedList[T interface{}] struct {
	tail *Node[T]
	head *Node[T]

	size int

	mu sync.Mutex
}

func NewLinkedList[T interface{}]() LinkedList[T] {
	return LinkedList[T]{
		tail: nil,
		head: nil,
		size: 0,
	}
}

func (l *LinkedList[T]) Append(data *T) *Node[T] {
	n := Node[T]{
		Data: data,
		Prev: nil,
		next: nil,
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	l.size += 1

	if l.tail == nil {
		l.tail = &n
		l.head = &n

		return &n
	}

	n.Prev = l.tail
	l.tail.next = &n
	l.tail = &n
	return &n
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

type LinkedListIterator[T interface{}] struct {
	node       *Node[T]
	linkedList *LinkedList[T]
}

func (l *LinkedList[T]) NewIterator() LinkedListIterator[T] {
	return LinkedListIterator[T]{
		node:       nil,
		linkedList: l,
	}
}

func (i *LinkedListIterator[T]) Next() bool {
	if i.node == nil {
		i.node = i.linkedList.head
	} else {
		i.node = i.node.next
	}

	return i.node != nil
}

func (i *LinkedListIterator[T]) Prev() bool {
	if i.node == nil {
		i.node = i.linkedList.tail
	} else {
		i.node = i.node.Prev
	}

	return i.node != nil
}

func (i *LinkedListIterator[T]) Get() *T {
	if i.node == nil {
		return nil
	}
	return i.node.Data
}
