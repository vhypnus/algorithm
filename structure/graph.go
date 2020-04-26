
//为什么临接表是连，删除时方便，新增都是往后加
//什么是临接矩阵

package structure

import (
	// "fmt"
)

type Vertex struct {

	val int

	//degree
	d int

}

func NewVertex(val int) *Vertex{
	var v = new (Vertex)
	v.val = val
	return v
}


type Edge struct {

	w int

	s *Vertex

	e *Vertex

}

// type Node struct {

// 	//head
// 	h *Edge

// 	//tail
// 	t *Edge

// 	next *Node

// }


type Graph struct {

	//vertex num
	vnum int 

	//edge num
	eno int

	Adj map[*Vertex][]*Edge
}

const (

	WHITE = iota

	GRAY 

	BLACK

)

func NewGraph() *Graph {
	var g = new(Graph)
	g.Adj = make(map[*Vertex][]*Edge)
	return g 
}

func (g *Graph) AddVertex(s *Vertex) {
	if g.Adj[s] != nil {
		//panic 
	}

	g.Adj[s] = make([]*Edge,0,4)
}


func (g *Graph) AddEdge(w int, s *Vertex,e *Vertex) {

	var edge = &Edge{w,s,e}

	g.Adj[s] = append(g.Adj[s],edge)
}


func (g *Graph) Bfs(s Vertex){

	// var m = make(map[Vertex]int)

	// for k,_ := range g.adj {
	// 	m[k] = WHITE
	// }

	// var q = NewQueue(32)
	// q.Push(&s)

	// for q.Size() > 0 {
	// 	v := q.Pop()
	// 	var adj = g.adj[*v]
	// 	for _,v := range adj {
	// 		if m[v] == WHITE {
	// 			m[v] = GRAY
	// 			q.Push(v.e)
	// 		}	
	// 	} 
	// } 
}

func (g *Graph) Dfs(s Vertex){

}

func (g *Graph) TopSort(){

}