package l82

import . "probestar.com/leet-go/list_node"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ret := &ListNode{0, nil}
	r := ret
	curr := &ListNode{head.Val, nil}
	count := 0
	for h := head.Next; h != nil; h = h.Next {
		if curr.Val != h.Val {
			if count == 0 {
				r.Next = curr
				r = r.Next
			}
			curr = &ListNode{h.Val, nil}
			count = 0
		} else {
			count++
		}
	}
	if count == 0 {
		r.Next = &ListNode{curr.Val, nil}
	}
	return ret.Next
}
