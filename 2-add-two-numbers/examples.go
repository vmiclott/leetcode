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

func Examples() {
	fmt.Println("2. Add Two Numbers")
	fmt.Println("Input: l1 = [2,4,3], l2 = [5,6,4]")
	fmt.Printf(
		"Output: %v\n",
		printLinkedList(addTwoNumbers(toLinkedList([]int{2, 4, 3}), toLinkedList([]int{5, 6, 4}))),
	)
	fmt.Println("Input: l1 = [0], l2 = [0]")
	fmt.Printf(
		"Output: %v\n",
		printLinkedList(addTwoNumbers(toLinkedList([]int{0}), toLinkedList([]int{0}))),
	)
	fmt.Println("Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]")
	fmt.Printf(
		"Output: %v\n",
		printLinkedList(addTwoNumbers(toLinkedList([]int{9, 9, 9, 9, 9, 9, 9}), toLinkedList([]int{9, 9, 9, 9}))),
	)
}
