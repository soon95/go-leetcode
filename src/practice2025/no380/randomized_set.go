package no380

import "math/rand"

type RandomizedSet struct {
	Set map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Set: make(map[int]bool),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exist := this.Set[val]

	this.Set[val] = true

	return !exist
}

func (this *RandomizedSet) Remove(val int) bool {

	_, exist := this.Set[val]

	delete(this.Set, val)

	return exist
}

func (this *RandomizedSet) GetRandom() int {

	cnt := rand.Intn(len(this.Set))

	for key := range this.Set {

		if cnt == 0 {
			return key
		} else {
			cnt--
		}
	}
	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
