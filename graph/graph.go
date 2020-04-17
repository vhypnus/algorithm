package graph 


type Vertex struct {
	Val int 

}


type Edge struct {
	Weight int

	Start *Vertex

	end *Vertex

}


type Graph struct {

	//哈希方案遍历
	EdgeMap map[int][]Vertex
}


func(g *Graph) addVertex(val int) {

}

func (g *Graph) addEdge(s int, e int) {

}

func (g *Graph) Bfs(val int ){

	//if not exists

	if g.EdgeMap[val] == nil {
		//panic
	}

	var s = g.EdgeMap[val]


}