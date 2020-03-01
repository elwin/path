package path

type node struct {
	identifier
	forward  map[identifier]*edge
	backward map[identifier]*edge
}

func newNode(i identifier) *node {
	return &node{
		identifier: i,
		forward:    map[identifier]*edge{},
		backward:   map[identifier]*edge{},
	}
}

type edge struct {
	cost                int
	source, destination *node
}

func newEdge(cost int, source, destination *node) *edge {
	return &edge{
		cost:        cost,
		source:      source,
		destination: destination,
	}
}
