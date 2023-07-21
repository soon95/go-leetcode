package main

/*
*208. 实现 Trie (前缀树)
https://leetcode.cn/problems/implement-trie-prefix-tree/
*/
type Trie struct {
	Dic map[byte]*Trie

	Closed bool
}

func Constructor() Trie {

	return Trie{
		Dic: make(map[byte]*Trie),
	}

}

func (this *Trie) Insert(word string) {

	if len(word) == 0 {
		this.Closed = true
		return
	}

	u := word[0]

	nextTrie, exist := this.Dic[u]

	if !exist {

		nextTrie = &Trie{
			Dic: map[byte]*Trie{},
		}

		this.Dic[u] = nextTrie
	}

	nextTrie.Insert(word[1:])

}

func (this *Trie) Search(word string) bool {

	if len(word) == 0 {

		return this.Closed

	}

	u := word[0]

	nextTrie, exist := this.Dic[u]

	if !exist {
		return false
	} else {
		return nextTrie.Search(word[1:])
	}

}

func (this *Trie) StartsWith(prefix string) bool {

	if len(prefix) == 0 {
		return true
	}

	u := prefix[0]

	nextTrie, exist := this.Dic[u]

	if !exist {
		return false
	} else {
		return nextTrie.StartsWith(prefix[1:])
	}
}

func main() {

	trie := Constructor()

	trie.Insert("apple")
	println(trie.Search("apple"))
	println(trie.Search("app"))
	println(trie.StartsWith("app"))
	trie.Insert("app")
	println(trie.Search("app"))
}
