package addTwoNumbers

import "strconv"

//ListNode is singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return int2Node(node2int(l1) + node2int(l2))
}
func node2int(l *ListNode) int {
	if l.Next == nil {
		return l.Val
	}
	str := strconv.Itoa(node2int(l.Next)) + strconv.Itoa(l.Val)
	v, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return v
}

func int2Node(sum int) *ListNode {
	for i, v := range "123456" {

	}
	return nil
}
