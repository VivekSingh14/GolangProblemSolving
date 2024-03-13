package main

import "fmt"

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func main() {
	trie := Constructor()

	trie.Insert("vivek")

	fmt.Println(trie.Search("vive"))

	fmt.Println(trie.StartsWith("viv"))
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (trie *Trie) Insert(word string) {
	curr := trie
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

/** Returns if the word is in the trie. */
func (trie *Trie) Search(word string) bool {
	curr := trie
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return curr.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (trie *Trie) StartsWith(prefix string) bool {
	curr := trie
	for _, ch := range prefix {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return true
}
