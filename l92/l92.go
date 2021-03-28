package l92

import (
	. "probestar.com/leet-go/list_node"
)

func reverse(head, tail *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	end := tail.Next
	for curr != end {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var l1, l2, l3, l4 *ListNode
	i := 1
	curr := head
	for curr != nil {
		if i == left-1 {
			l1 = curr
		}
		if i == left {
			l2 = curr
		}
		if i == right {
			l3 = curr
		}
		if i == right+1 {
			l4 = curr
			break
		}
		curr = curr.Next
		i++
	}

	l5 := reverse(l2, l3)
	if l1 != nil {
		l1.Next = l5
	} else {
		head = l5
	}
	l2.Next = l4
	return head
}
