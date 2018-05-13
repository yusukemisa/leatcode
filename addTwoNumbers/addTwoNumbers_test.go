package addTwoNumbers

import (
	"reflect"
	"testing"
)

//twoSum(nums []int, target int)
func TestAddTwoNumbers(t *testing.T) {

	l1 := &ListNode{
		2, &ListNode{
			4, &ListNode{
				3, nil}},
	}
	l2 := &ListNode{
		5, &ListNode{
			6, &ListNode{
				4, nil}},
	}
	expected := &ListNode{
		7, &ListNode{
			0, &ListNode{
				8, nil}},
	}
	actual := addTwoNumbers(l1, l2)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v ,actual %v", expected, actual)
	}

}
