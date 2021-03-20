package list_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateLitNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func GenerateLitNodeWithNext(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

func GenerateList(list []int) *ListNode {
	var pre, node *ListNode
	for i := len(list) - 1; i >= 0; i-- {
		node = GenerateLitNodeWithNext(list[i], pre)
		pre = node
	}
	return node
}
