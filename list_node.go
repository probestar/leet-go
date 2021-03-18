package leet

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateLitNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func generateLitNodeWithNext(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

func generateList(list []int) *ListNode {
	var pre, node *ListNode
	for i := len(list) - 1; i >= 0; i-- {
		node = generateLitNodeWithNext(list[i], pre)
		pre = node
	}
	return node
}
