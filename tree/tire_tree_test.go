package tree

import "testing"

func TestTrie_Insert(t *testing.T) {
	trie := Constructor()
	trie.Insert("abp")
	t.Log(trie.Search("apple"))
	t.Log(trie.Search("app"))
	t.Log(trie.StartsWith("app"))
	trie.Insert("abe")
	trie.Insert("aef")
	t.Log(trie.SearchNode("ae"))

}
