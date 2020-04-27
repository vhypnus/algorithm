
//为什么临接表是连，删除时方便，新增都是往后加
//什么是临接矩阵

package structure

import (
	// "fmt"
	// "container/list"
)

type Vertex struct {

	val interface{}

	//degree
	d int

}



type Edge struct {

	w int

	s *Vertex

	e *Vertex

}

type Adj struct {

	//head
	val *Edge

	next *Adj



}


type Graph struct {

	//vertex num
	vnum int 

	//edge num
	eno int

	// 临接矩阵
	Adj map[*Vertex]*Adj
}

const (

	WHITE = iota

	GRAY 

	BLACK

)


func NewVertex(val interface{}) *Vertex{
	var v = &Vertex{val:val}
	return v
}

func NewEdge(w int, s *Vertex,e *Vertex) *Edge {
	var edge = &Edge{w,s,e}
	return edge 
}

func NewGraph() *Graph {
	var g = new(Graph)
	g.Adj = make(map[*Vertex]*Adj)
	return g 
}

func (g *Graph) AddVertex(s *Vertex) {
	if g.Adj[s] != nil {
		//panic 
	}

	g.Adj[s] = nil

}


func (g *Graph) AddEdge(edge *Edge) {

	var s = edge.s

	// head
	var h = g.Adj[s]

	if h == nil {
		h = &Adj{val:edge}
		g.Adj[s] = h
	} else {
		//parent
		p := &Adj{val:edge}
		p.next = h
		g.Adj[s] = p 
	}
		
}


// FIFO
func (g *Graph) Bfs(s *Vertex){

	var m = make(map[*Vertex]byte)
	for k,_ := range g.Adj {
		m[k] = WHITE
	}
	// list or slice 
	// list 空间比slice 大，但slice会有扩容成本
	var q = make([]*Vertex,0,16)
	m[s] = GRAY
	q = append(q,s)

	for len(q) > 0 {
		v := q[0]
		q = q[1:] 

		adj := g.Adj[s]
		for adj.val != nil {
			if m[adj.val.s] == WHITE {
				m[adj.val.e] = GRAY
				q = append(q,adj.val.e)
			}

			//
			adj = adj.next
		}
		m[v] = BLACK
	}

	
}


//1、FILO 栈 
//2、递归
func (g *Graph) Dfs(s *Vertex){
	//  

	var m = make(map[*Vertex]byte)
	for k,_ := range g.Adj {
		m[k] = WHITE
	}

	m[s] = GRAY
	
	for k,_ := range g.Adj {
		g.dfsvisit(m,k)
	}



}

func (g *Graph) dfsvisit(m map[*Vertex]byte,v *Vertex)	{

	head := g.Adj[v]
	for head != nil  {
		if m[head.val.e] == WHITE {
			m[head.val.e] = GRAY
			g.dfsvisit(m, head.val.e)
		}

		head = head.next
	}
}	



func (g *Graph) TopolgicalSort() {


}


func (g *Graph) MST(){

}


//Strongly connected components
func (g *Graph) SCC() bool{

	return false
}

func (g *Graph) Tarjan() int {

	return 0
}


