// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package path

type edgeQueue struct {
	start   *edgeNode
	compare func(a, b *edge) int
}

type edgeNode struct {
	internal *edge
	next     *edgeNode
}

func NewEdgeQueue(compare func(a, b *edge) int) *edgeQueue {
	return &edgeQueue{compare: compare}
}

func (q *edgeQueue) push(item *edge) {
	if q.start == nil {
		q.start = &edgeNode{item, nil}
		return
	}

	if q.compare(item, q.start.internal) > 0 {
		start := &edgeNode{item, q.start}
		q.start = start
		return
	}

	previous := q.start
	for {
		if previous.next == nil || q.compare(item, previous.next.internal) > 0 {
			previous.next = &edgeNode{item, previous.next}
			return
		}

		previous = previous.next
	}
}

func (q *edgeQueue) pop() *edge {
	cur := q.start
	q.start = cur.next
	return cur.internal
}

func (q *edgeQueue) hasNext() bool {
	return q.start != nil
}
