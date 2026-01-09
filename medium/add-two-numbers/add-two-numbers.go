package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	result := addTwoNumbers(l1, l2)

	fmt.Printf("\nRESULT: %v \n", result)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	/*
		- transform the digits into a number
		- add the two numbers
		- reserve the number
	*/

	fixedNode := &ListNode{}
	currentNode := fixedNode
	carry := 0

	for l1 != nil || l2 != nil {

		//? Step 1: Get values or zero (if list ended)
		valList1 := 0
		valList2 := 0

		if l1 != nil {
			valList1 = l1.Val
		}
		if l2 != nil {
			valList2 = l2.Val
		}

		//? Step 2, math: sum the numbers, track the digit value with % and the carry with the division
		sum := valList1 + valList2 + carry
		digit := sum % 10
		carry = sum / 10

		//? Step 3: Create a new node and link it with the current, then slide the current foward
		currentNode.Next = &ListNode{Val: digit}

		fmt.Println("Current: ", digit)
		currentNode = currentNode.Next

		//? Step 4 advance original l1, l2 lists if not nil
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

	}

	if carry > 0 {
		currentNode.Next = &ListNode{Val: carry}
	}

	return fixedNode.Next
}
