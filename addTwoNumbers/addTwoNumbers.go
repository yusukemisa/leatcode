package addTwoNumbers

import (
	"container/list"
	"log"
	"strconv"
)

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
func printList(l *ListNode) {
	j := 1
	for c := l; c != nil; c = c.Next {
		log.Printf("%v:%v", j, c.Val)
		j++
	}
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
	l := list.New()
	for _, v := range strconv.Itoa(sum) {
		val, _ := strconv.Atoi(string(v))
		l.PushFront(val)
	}
	returnList := &ListNode{
		Val:  0,
		Next: nil,
	}
	j := 0
	for e := l.Front(); e != nil; e = e.Next() {
		add(returnList, e.Value.(int), j)
		j++
	}
	return returnList
}

func add(l *ListNode, v int, counter int) *ListNode {
	if counter == 0 {
		l.Val = v
		return l
	}
	for c := l; c != nil; c = c.Next {
		if c.Next == nil {
			c.Next = &ListNode{
				Val:  v,
				Next: nil,
			}
			break
		}
	}
	return l
}
