// 380. Insert Delete GetRandom O(1)
package randomized_set

import "math/rand"

type RandomizedSet struct {
	m map[int]int
	l []int
}

func Constructor() RandomizedSet {
	m := make(map[int]int) // map value to index
	l := make([]int, 0)
	return RandomizedSet{m, l}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = len(this.l)
	this.l = append(this.l, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.m[val]
	if !ok {
		return false
	}
	n := len(this.l) - 1
	this.l[i], this.l[n] = this.l[n], this.l[i]
	this.m[this.l[i]] = i
	this.l = this.l[:n]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.l))
	return this.l[i]
}
