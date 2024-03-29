// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package path

type destinationQueue struct {
	start   *destinationNode
	compare func(a, b *destination) int
}

type destinationNode struct {
	internal *destination
	next     *destinationNode
}

func NewDestinationQueue(compare func(a, b *destination) int) *destinationQueue {
	return &destinationQueue{compare: compare}
}

func (q *destinationQueue) push(item *destination) {
	if q.start == nil {
		q.start = &destinationNode{item, nil}
		return
	}

	if q.compare(item, q.start.internal) > 0 {
		start := &destinationNode{item, q.start}
		q.start = start
		return
	}

	previous := q.start
	for {
		if previous.next == nil || q.compare(item, previous.next.internal) > 0 {
			previous.next = &destinationNode{item, previous.next}
			return
		}

		previous = previous.next
	}
}

func (q *destinationQueue) pop() *destination {
	cur := q.start
	q.start = cur.next
	return cur.internal
}

func (q *destinationQueue) hasNext() bool {
	return q.start != nil
}
