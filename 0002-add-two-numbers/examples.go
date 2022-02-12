package addtwonumbers

import (
	"fmt"
)

func toLinkedList(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}
	return &ListNode{
		Val:  slice[0],
		Next: toLinkedList(slice[1:]),
	}
}

func printLinkedList(node *ListNode) []int {
	res := []int{node.Val}
	for next := node.Next; next != nil; {
		res = append(res, next.Val)
		next = next.Next
	}
	return res
}

func example(slice1 []int, slice2 []int) {
	fmt.Printf("Input: l1 = %v, l2 = %v\n", slice1, slice2)
	fmt.Printf(
		"Output: %v\n",
		printLinkedList(addTwoNumbers(toLinkedList(slice1), toLinkedList(slice2))),
	)
}

func Examples() {
	fmt.Println("2. Add Two Numbers")
	example([]int{2, 4, 3}, []int{5, 6, 4})
	example([]int{0}, []int{0})
	example([]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9})
}
