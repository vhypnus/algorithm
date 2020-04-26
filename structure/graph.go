
//为什么临接表是连，删除时方便，新增都是往后加
//什么是临接矩阵

package structure

type Vertex struct {

	v int

	d int

}

func NewVertex(v int) *Vertex{
	var v = new (Vertex)
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

	adj map[Vertex][]*Edge
}

const (

	WHITE = iota

	GRAY 

	BLACK

)

func (g *Graph) AddVertex(s *Vertex) {
	if g.adj[m] != nil {
		//panic 
	}

	g.adj[s] = make([]*Edge,0,4)
}

func (g *Graph) AddEdge(w int, s *Vertex,e *Vertex) {
	var adj = g.adj[s]

	var e = &Edge{w,s,e}

	g.adj[s] = append(g.adj[s],e)

}


func (g *Graph) Bfs(s Vertex){

	var m = make(map[Vertex]int)

	for k,_ := range g.adj {
		m[k] = WHITE
	}

	var q = NewQueue(32)
	q.Push(s)

	for q.Size() > 0 {
		v := q.Pop()
		var adjlist = g.adj[v]
		for _,v := range adjlist {
			if m[v] == WHITE {
				m[v] == GRAY
				q.Push(v.e)
			}	
		} 
	} 
}

func (g *Graph) Dfs(s Vertex){

}

func (g *Graph) TopSort(){

}