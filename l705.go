package leet

const myHashSetCap = 100000

type MyHashSet struct {
	s [myHashSetCap]bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	this.s[key%myHashSetCap] = true
}

func (this *MyHashSet) Remove(key int) {
	this.s[key%myHashSetCap] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.s[key%myHashSetCap]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
