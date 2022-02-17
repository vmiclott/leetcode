package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNodes(node1 *ListNode, node2 *ListNode, add int) *ListNode {
	if node1 == nil && node2 == nil && add == 0 {
		return nil
	}

	var (
		next1, next2 *ListNode
		val1, val2   int
	)

	if node1 != nil {
		val1 = node1.Val
		next1 = node1.Next
	}

	if node2 != nil {
		val2 = node2.Val
		next2 = node2.Next
	}

	sum, addToNext := val1+val2+add, 0
	if sum > 9 {
		sum -= 10
		addToNext = 1
	}

	return &ListNode{
		Val:  sum,
		Next: addTwoNodes(next1, next2, addToNext),
	}
}

func addTwoNumbers(node1 *ListNode, node2 *ListNode) *ListNode {
	return addTwoNodes(node1, node2, 0)
}
