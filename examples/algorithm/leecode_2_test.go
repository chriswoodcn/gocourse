package algorithm

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func traverse(node *ListNode) {
	nodeCopy := node
	for {
		if nodeCopy == nil {
			break
		}
		fmt.Printf("%d ", nodeCopy.Val)
		nodeCopy = nodeCopy.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func buildListNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	start := &ListNode{0, nil}
	var pre *ListNode
	for i, v := range list {
		if i == 0 {
			start.Val = v
			pre = start
		} else {
			n := &ListNode{v, nil}
			pre.Next = n
			pre = n
		}
	}
	return start
}
func Test(t *testing.T) {
	//a1 := []int{2, 4, 3}
	a1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	node1 := buildListNode(a1)
	//a2 := []int{5, 6, 4}
	a2 := []int{5, 6, 4}
	node2 := buildListNode(a2)
	numbers := addTwoNumbers(node1, node2)
	traverse(numbers)
}
