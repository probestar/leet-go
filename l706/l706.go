package l706

import "container/list"

const myHashMapCap = 1000

type myHashMapElement struct {
	key   int
	value int
}

type MyHashMap struct {
	data []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, myHashMapCap)}
}

func hash(key int) int {
	return key % myHashMapCap
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	k := hash(key)
	l := this.data[k]
	for e := l.Front(); e != nil; e = e.Next() {
		el, _ := e.Value.(myHashMapElement)
		if el.key == key {
			e.Value = myHashMapElement{key, value}
			return
		}
	}
	this.data[k].PushBack(myHashMapElement{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	k := hash(key)
	l := this.data[k]
	for e := l.Front(); e != nil; e = e.Next() {
		el, _ := e.Value.(myHashMapElement)
		if el.key == key {
			return el.value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	k := hash(key)
	l := this.data[k]
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(myHashMapElement).key == key {
			l.Remove(e)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
