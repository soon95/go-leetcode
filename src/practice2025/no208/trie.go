package no208

type Trie struct {
	Data  map[byte]Trie
	Close bool
}

func Constructor() Trie {
	return Trie{
		Data: make(map[byte]Trie),
	}
}

func (this *Trie) Insert(word string) {

	if len(word) == 0 {
		this.Close = true
		return
	}

	startChar := word[0]

	next, exist := this.Data[startChar]
	if !exist {
		next = Constructor()
	}
	next.Insert(word[1:])

	this.Data[startChar] = next
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		if this.Close {
			return true
		} else {
			return false
		}
	}

	startChar := word[0]
	next, exist := this.Data[startChar]
	if exist {
		return next.Search(word[1:])
	} else {
		return false
	}
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}

	startChar := prefix[0]
	next, exist := this.Data[startChar]
	if exist {
		return next.StartsWith(prefix[1:])
	} else {
		return false
	}
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
