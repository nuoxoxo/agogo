package main

import "fmt"

const AlphaSize = 26

type Node struct {
	children [AlphaSize]*Node
	isEnded  bool
}

type Trie struct {
	root *Node
}

func makeTrie() *Trie {
	res := &Trie{root: &Node{}}
	return res
}

func (trie *Trie) Insert(word string) {
	// traverse the word char by char
	i := 0
	N := len(word)
	curr := trie.root
	for i < N {
		charcode := word[i] - 'a'
		if curr.children[charcode] == nil {
			curr.children[charcode] = &Node{}
		}
		curr = curr.children[charcode]
		i++
	}
	curr.isEnded = true
}

func (trie *Trie) Search(word string) bool {
	i := 0
	N := len(word)
	curr := trie.root
	for i < N {
		charcode := word[i] - 'a'
		if curr.children[charcode] == nil {
			return false
		}
		curr = curr.children[charcode]
		i++
	}
	return curr.isEnded
}

func main() {
	testTrie := makeTrie()
	{
		fmt.Println("/trie", testTrie)
		fmt.Println("/root", testTrie.root)
		fmt.Println("/root.children", testTrie.root.children, "\n")
	}
	{
		testTrie.Insert("maxheap")
		fmt.Println("/trie", testTrie)
		fmt.Println("/root", testTrie.root)
		fmt.Println("/root.children", testTrie.root.children)
		fmt.Println("/m", testTrie.root.children['m'-'a'])
		fmt.Println("/a", testTrie.root.children['m'-'a'].children[0])
		fmt.Println("/x", testTrie.root.children['m'-'a'].children[0].children['x'-'a'])
		fmt.Println("/lookup maxheap/max", testTrie.Search("max"))
		fmt.Println("/lookup maxheap/maxheap", testTrie.Search("maxheap"))
	}
}
