package twoSum

import (
	"reflect"
	"testing"
)

//twoSum(nums []int, target int)
func TestTwoSum(t *testing.T) {
	expected := []int{0, 1}
	actual := twoSum([]int{2, 7, 11, 15}, 9)
	eval(t, expected, actual)

	expected = []int{1, 2}
	actual = twoSum([]int{2, 7, 11, 15}, 18)
	eval(t, expected, actual)

	expected = []int{2, 3}
	actual = twoSum([]int{2, 7, 11, 15}, 26)
	eval(t, expected, actual)

}

func eval(t *testing.T, expected []int, actual []int) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v ,actual %v", expected, actual)
	}
}
