package path

type identifier int // identifier = priority

func New() *graph {
	return &graph{
		nodes: map[identifier]*node{},
	}
}

type graph struct {
	nodes map[identifier]*node
}

func (g *graph) addNodes(ids ...identifier) {
	for _, id := range ids {
		g.nodes[id] = newNode(id)
	}
}

func (g *graph) directed(a, b identifier) {
	g.nodes[a].forward[b] = newEdge(1, g.nodes[a], g.nodes[b])
	g.nodes[b].backward[a] = newEdge(1, g.nodes[b], g.nodes[a])
}

func (g *graph) undirected(a, b identifier) {
	g.directed(a, b)
	g.directed(b, a)
}

// func (g *graph) contract(id identifier) {
// 	u := g.nodes[id]
// 	for v := range u.backward {
// 		for w := range u.forward {
//
// 		}
// 	}
// }

type destination struct {
	identifier
	cost int
}

func (g *graph) shortestPath(a, b, ignore identifier) int {
	// queue with ascending cost
	q := NewDestinationQueue(func(a, b *destination) int {
		if a.cost > b.cost {
			return -1
		} else if a.cost < b.cost {
			return 1
		} else {
			return 0
		}
	})

	// identifier -> cost
	visited := map[identifier]bool{}

	n := g.nodes[a]

	for neighbour, e := range n.forward {
		q.push(&destination{
			identifier: neighbour,
			cost: e.cost,
		})
	}

	for q.hasNext() {
		next := q.pop()
		if visited[next.identifier] || next.identifier == ignore {
			continue
		}

		if next.identifier == b {
			return next.cost
		}

		visited[next.identifier] = true

		for neighbour, edge := range g.nodes[next.identifier].forward {
			q.push(&destination{
				identifier: neighbour,
				cost:      next.cost + edge.cost,
			})
		}
	}

	return -1
}
