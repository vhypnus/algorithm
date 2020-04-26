package structure


type Trie struct {
	
	children map[rune]*Trie

}


func NewTrie() *Trie {

	var node = new(Trie)
	node.children = make(map[rune]*Trie)
	return node ;
}


func (root *Trie) add(word string) {
	var p = root
	for _,v := range word {
		if p.children[v] == nil {
			node := NewTrie()
			p.children[v] = node
		}

		p = p.children[v]

	}
}

func (root *Trie) find(word string) bool {
	var r bool = true 
	var p *Trie = root

	for _,v := range word {
		if p.children[v] == nil {
			r = false
			break
		}
		p = p.children[v]
	}

	return r 
}