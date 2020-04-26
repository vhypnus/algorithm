package structure


import (
	"testing"
	"fmt"
)

func TestFind(t *testing.T) {
	var trie = NewTrie()

	trie.add("hello")

	fmt.Println(trie.find("hello"))
}