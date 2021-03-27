package l61

import "probestar.com/leet-go/list_node"

func rotateRight(head *list_node.ListNode, k int) *list_node.ListNode {
	if head == nil || k == 0 {
		return head
	}

	var tail, m *list_node.ListNode
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
		tail = node
	}
	k = k % n
	k = n - k
	n = 0
	for node := head; node != nil; node = node.Next {
		if n == k-1 {
			m = node
			break
		}
		n++
	}
	tail.Next = head
	head = m.Next
	m.Next = nil
	return head
}
