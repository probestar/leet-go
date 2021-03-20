package l705

import "testing"

func Test_MyHashSet(t *testing.T) {
	s := Constructor()
	s.Add(1)
	s.Add(2)
	if !s.Contains(1) {
		t.Fail()
	}
	if s.Contains(3) {
		t.Fail()
	}
	s.Add(2)
	if !s.Contains(2) {
		t.Fail()
	}
	s.Remove(2)
	if s.Contains(2) {
		t.Fail()
	}
}
