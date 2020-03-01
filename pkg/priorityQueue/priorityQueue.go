package priorityQueue

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../path/destination-$GOFILE -pkg=path gen "something=destination"

type something generic.Type

type somethingQueue struct {
	start   *somethingNode
	compare func(a, b *something) int
}

type somethingNode struct {
	internal *something
	next *somethingNode
}

func NewsomethingQueue(compare func(a, b *something) int) *somethingQueue {
	return &somethingQueue{compare: compare}
}

func (q *somethingQueue) push(item *something) {
	if q.start == nil {
		q.start = &somethingNode{item, nil}
		return
	}

	if q.compare(item, q.start.internal) > 0 {
		start := &somethingNode{item, q.start}
		q.start = start
		return
	}

	previous := q.start
	for {
		if previous.next == nil || q.compare(item, previous.next.internal) > 0 {
			previous.next = &somethingNode{item, previous.next}
			return
		}

		previous = previous.next
	}
}

func (q *somethingQueue) pop() *something {
	cur := q.start
	q.start = cur.next
	return cur.internal
}

func (q *somethingQueue) hasNext() bool {
	return q.start != nil
}
