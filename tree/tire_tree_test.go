package tree

import "testing"

func TestTrie_Insert(t *testing.T) {
	trie := Constructor()
	trie.Insert("abp")
	/*t.Log(trie.Search("apple"))
	t.Log(trie.Search("app"))
	t.Log(trie.StartsWith("app"))*/
	trie.Insert("aeef")
	trie.Insert("aeff")
	t.Log(trie.StartsPrx("ae"))
	t.Log(trie.SearchNode("ae"))
}
