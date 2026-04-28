package code

import "fmt"

/*
2. Add Two Numbers
Source: https://leetcode.com/problems/add-two-numbers/

You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

// ListNode represents a node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers adds two numbers represented as linked lists
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{sum % 10, nil}
		current = current.Next
	}

	return dummy.Next
}

// Helper function to create a linked list from a slice of integers
func CreateLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{nums[0], nil}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{nums[i], nil}
		current = current.Next
	}

	return head
}

// Helper function to convert linked list to slice for easy testing/display
func LinkedListToSlice(head *ListNode) []int {
	var result []int
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

// Helper function to format linked list as string "val -> val -> val"
func LinkedListToString(head *ListNode) string {
	if head == nil {
		return "nil"
	}

	result := ""
	current := head

	for current != nil {
		if result != "" {
			result += " -> "
		}
		result += fmt.Sprintf("%d", current.Val)
		current = current.Next
	}

	return result
}
