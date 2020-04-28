package structure

import (
	"testing"
	"fmt"
)

/*

r,s,t,u
v,w,x,y

(r,s),(r,v)
(s,w)
(t,u),(t,x)
(u,t),(u,x),(u,y)
(v,r)
(w,x),(w,t),(w,x)
(x,t),(x,u),(x,w),(x,y)
(y,u),(y,x)
*/

func TestAddVertex(t *testing.T) {

	var g = NewGraph()

	g.AddVertex("r")
	g.AddVertex("s")
	// g.AddVertex(1.0) // error => invalid type
	// g.AddVertex("r") // error => already exists

	fmt.Println(g)

	
}


/*

r,s,t,u
v,w,x,y

(r,s),(r,v)
(s,w)
(t,u),(t,x)
(u,t),(u,x),(u,y)
(v,r)
(w,s),(w,t),(w,x)
(x,t),(x,u),(x,w),(x,y)
(y,u),(y,x)
*/
func initGraph() *Graph{
	var g = NewGraph()
	g.AddEdge(1,"r","s")
	g.AddEdge(1,"r","v")

	g.AddEdge(1,"s","w")	
	g.AddEdge(1,"s","r")

	g.AddEdge(1,"t","u")
	g.AddEdge(1,"t","x")

	g.AddEdge(1,"u","t")
	g.AddEdge(1,"u","x")
	g.AddEdge(1,"u","y")

	g.AddEdge(1,"v","r")

	g.AddEdge(1,"w","w")
	g.AddEdge(1,"w","t")
	g.AddEdge(1,"w","x")

	g.AddEdge(1,"x","t")
	g.AddEdge(1,"x","u")
	g.AddEdge(1,"x","w")
	g.AddEdge(1,"x","y")

	g.AddEdge(1,"y","u")
	g.AddEdge(1,"y","x")
	return g
}

func printGraph(g *Graph) {
	for v,adj := range g.Adj {
		fmt.Printf("vertex %v : ", v)

		for adj != nil && adj.e != nil{
			fmt.Printf(" (%v,%v) ",v,adj.e)
			adj = adj.next
		}
		fmt.Println()
	}
}

func TestBfs(t *testing.T) {
	var g = initGraph()
	g.Bfs("s")
}

func TestDfs(t *testing.T) {
	var g = initGraph()
	g.Dfs("s")
}