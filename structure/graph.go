
//为什么临接表是连，删除时方便，新增都是往后加
//什么是临接矩阵

package structure

import (
	"fmt"
	// "container/list"
)


type Adj struct {

	// weight
	w int

	//endpoint
	e interface{}

	next *Adj



}


type Graph struct {

	//vertex num
	vno int 

	//edge num
	eno int

	// 临接矩阵
	Adj map[interface{}]*Adj
}

const (

	WHITE = iota

	GRAY 

	BLACK

)

func NewGraph() *Graph {
	var g = new(Graph)
	g.Adj = make(map[interface {}]*Adj)
	return g 
}

// val values ,not pointers
// case number or string
func (g *Graph) AddVertex(v interface{}) {

	 switch t := v.(type) {
	 	case int,string :
	 		//

	 	default:
	 		//panic
	 		fmt.Println(t)
	 		panic("type error")

	 }

	if g.Adj[v] != nil {
		panic("already exists.")
	}

	g.Adj[v] = &Adj{}
	g.vno++

}

func (g *Graph) AddEdge(w int, s interface{},e interface{}) {

	if g.Adj[e] == nil {
		g.Adj[e] = &Adj{}
	}

	if g.Adj[s] == nil {
		g.Adj[s] = &Adj{w:w,e:e}
	} else {
		g.Adj[s] = &Adj{w:w,e:e,next:g.Adj[s]}
	}

	g.eno++
}


// FIFO
func (g *Graph) Bfs(s interface{}){

	var m = make(map[interface{}]byte)
	for k,_ := range g.Adj {
		m[k] = WHITE
	}
	// list or slice 
	// list 空间比slice 大，但slice会有扩容成本
	var q = make([]interface{},0,16)
	m[s] = GRAY
	q = append(q,s)

	for len(q) > 0 {
		// fmt.Println(q)
		v := q[0]
		fmt.Printf("%v --> ",v)
		q = q[1:] 

		adj := g.Adj[v]
		for adj != nil && adj.e != nil {
			if m[adj.e] == WHITE {
				m[adj.e] = GRAY
				q = append(q,adj.e)
			}

			//
			adj = adj.next
		}
		m[v] = BLACK

	}

	fmt.Println()

	
}


//1、FILO 栈 
//2、递归
func (g *Graph) Dfs(s interface{}){
	//  

	var m = make(map[interface{}]byte)
	for k,_ := range g.Adj {
		m[k] = WHITE
	}
	
	for v,_ := range g.Adj {
		
		if m[v] == WHITE {
			fmt.Printf("%v --> ",v)	
			g.dfsvisit(m,v)	
		}
		
	}



}

func (g *Graph) dfsvisit(m map[interface{}]byte,s interface{})	{

	h := g.Adj[s]
	m[s] = GRAY
	for h != nil && h.e != nil  {
		
		if m[h.e] == WHITE {
			fmt.Printf("%v --> ",h.e)
			g.dfsvisit(m, h.e)	
		}
		h = h.next

		

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


