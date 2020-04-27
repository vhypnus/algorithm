package structure

import (
	"testing"
	"fmt"
)

/*

1,2,3,4,5





*/

func TestAddVertex(t *testing.T) {
	var v1 = NewVertex(1)

	var v2 = NewVertex(2)

	var edge = NewEdge(3,v1,v2)

	var g = NewGraph()

	g.AddVertex(v1)
	g.AddVertex(v2)

	g.AddEdge(edge)

	var adj = g.Adj[v1]

	fmt.Println(adj.val.s)

	// for k,v := range g.Adj {
	// 	if len(v) > 0 {
	// 		fmt.Println(*k)
	// 		fmt.Println(*v[0])	
	// 	}
		
	// }
}