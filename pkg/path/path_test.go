package path

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := New()

	g.addNodes(1, 2, 3, 4, 5, 6)
	g.undirected(2, 6)
	g.undirected(6, 1)
	g.undirected(1, 4)
	g.undirected(4, 3)
	g.undirected(3, 5)

	fmt.Println(g.shortestPath(2, 5, 1))
}