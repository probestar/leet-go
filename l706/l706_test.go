package l706

import (
	"container/list"
	"testing"
)

func Test_MyHashMap(t *testing.T) {
	m := Constructor()
	m.Put(1, 1)
	m.Put(2, 2)
	if m.Get(1) != 1 {
		t.Fail()
	}
	if m.Get(3) != -1 {
		t.Fail()
	}
	m.Put(2, 1)
	if m.Get(2) != 1 {
		t.Fail()
	}
	m.Remove(2)
	//if m.Get(2) != -1 {
	//	t.Fail()
	//}
}

func Test_List(t *testing.T) {
	l := list.New()
	l.PushBack(1)
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == 1 {
			l.Remove(e)
		}
	}
	if l.Len() != 0 {
		t.Fail()
	}
}
