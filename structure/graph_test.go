package structure

import (
	"testing"
	"fmt"
)

func TestAddVertex(t *testing.T) {
	var v1 = NewVertex(1)

	var v2 = NewVertex(2)
	fmt.Println(v1)
	fmt.Println(v2)

	var g = NewGraph()

	g.AddVertex(v1)
	g.AddVertex(v2)

	g.AddEdge(3,v1,v2)
	fmt.Println(g)

	for k,v := range g.Adj {
		if len(v) > 0 {
			fmt.Println(*k)
			fmt.Println(*v[0])	
		}
		
	}
}