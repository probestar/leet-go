package l82

import "probestar.com/leet-go/list_node"

func deleteDuplicates(head *list_node.ListNode) *list_node.ListNode {
	if head == nil {
		return nil
	}
	ret := &list_node.ListNode{0, nil}
	r := ret
	curr := &list_node.ListNode{head.Val, nil}
	count := 0
	for h := head.Next; h != nil; h = h.Next {
		if curr.Val != h.Val {
			if count == 0 {
				r.Next = curr
				r = r.Next
			}
			curr = &list_node.ListNode{h.Val, nil}
			count = 0
		} else {
			count++
		}
	}
	if count == 0 {
		r.Next = &list_node.ListNode{curr.Val, nil}
	}
	return ret.Next
}
